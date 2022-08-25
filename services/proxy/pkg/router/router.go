package router

import (
	"context"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/owncloud/ocis/v2/ocis-pkg/log"
	"github.com/owncloud/ocis/v2/ocis-pkg/registry"
	"github.com/owncloud/ocis/v2/services/proxy/pkg/config"
	"github.com/owncloud/ocis/v2/services/proxy/pkg/proxy/policy"
	"go-micro.dev/v4/selector"
)

const directorCtxKey string = "director"

func Middleware(policySelector *config.PolicySelector, policies []config.Policy, logger log.Logger) func(http.Handler) http.Handler {
	router := New(policySelector, policies, logger)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fn := router.Route(r)
			next.ServeHTTP(w, r.WithContext(SetDirectorFunc(r.Context(), fn)))
		})
	}
}

func (rt Router) Route(r *http.Request) func(*http.Request) {
	pol, err := rt.policySelector(r)
	if err != nil {
		rt.logger.Error().Err(err).Msg("Error while selecting pol")
		return nil
	}

	if _, ok := rt.directors[pol]; !ok {
		rt.logger.
			Error().
			Str("policy", pol).
			Msg("policy is not configured")
		return nil
	}

	method := ""
	// find matching director
	for _, rtype := range config.RouteTypes {
		var handler func(string, url.URL) bool
		switch rtype {
		case config.QueryRoute:
			handler = queryRouteMatcher
		case config.RegexRoute:
			handler = rt.regexRouteMatcher
		case config.PrefixRoute:
			fallthrough
		default:
			handler = prefixRouteMatcher
		}
		if rt.directors[pol][rtype][r.Method] != nil {
			// use specific method
			method = r.Method
		}
		for endpoint := range rt.directors[pol][rtype][method] {
			if handler(endpoint, *r.URL) {

				rt.logger.Debug().
					Str("policy", pol).
					Str("method", r.Method).
					Str("prefix", endpoint).
					Str("path", r.URL.Path).
					Str("routeType", string(rtype)).
					Msg("director found")

				return rt.directors[pol][rtype][method][endpoint]
			}
		}
	}

	// override default director with root. If any
	switch {
	case rt.directors[pol][config.PrefixRoute][method]["/"] != nil:
		// try specific method
		return rt.directors[pol][config.PrefixRoute][method]["/"]
	case rt.directors[pol][config.PrefixRoute][""]["/"] != nil:
		// fallback to unspecific method
		return rt.directors[pol][config.PrefixRoute][""]["/"]
	}

	rt.logger.
		Warn().
		Str("policy", pol).
		Str("path", r.URL.Path).
		Msg("no director found")
	return nil
}

func New(policySelector *config.PolicySelector, policies []config.Policy, logger log.Logger) Router {
	if policySelector == nil {
		firstPolicy := policies[0].Name
		logger.Warn().Str("policy", firstPolicy).Msg("policy-selector not configured. Will always use first policy")
		policySelector = &config.PolicySelector{
			Static: &config.StaticSelectorConf{
				Policy: firstPolicy,
			},
		}
	}

	logger.Debug().
		Interface("selector_config", policySelector).
		Msg("loading policy-selector")

	selector, err := policy.LoadSelector(policySelector)
	if err != nil {
		logger.Fatal().Err(err).Msg("Could not load policy-selector")
	}

	r := Router{
		directors:      make(map[string]map[config.RouteType]map[string]map[string]func(req *http.Request)),
		policySelector: selector,
	}
	for _, pol := range policies {
		for _, route := range pol.Routes {
			logger.Debug().Str("fwd: ", route.Endpoint)

			if route.Backend == "" && route.Service == "" {
				logger.Fatal().Interface("route", route).Msg("neither Backend nor Service is set")
			}
			uri, err2 := url.Parse(route.Backend)
			if err2 != nil {
				logger.
					Fatal(). // fail early on misconfiguration
					Err(err2).
					Str("backend", route.Backend).
					Msg("malformed url")
			}

			// here the backend is used as a uri
			r.addHost(pol.Name, uri, route)
		}
	}
	return r
}

type Router struct {
	logger         log.Logger
	directors      map[string]map[config.RouteType]map[string]map[string]func(req *http.Request)
	policySelector policy.Selector
}

func (rt Router) addHost(policy string, target *url.URL, route config.Route) {
	targetQuery := target.RawQuery
	if rt.directors[policy] == nil {
		rt.directors[policy] = make(map[config.RouteType]map[string]map[string]func(req *http.Request))
	}
	routeType := config.DefaultRouteType
	if route.Type != "" {
		routeType = route.Type
	}
	if rt.directors[policy][routeType] == nil {
		rt.directors[policy][routeType] = make(map[string]map[string]func(req *http.Request))
	}
	if rt.directors[policy][routeType][route.Method] == nil {
		rt.directors[policy][routeType][route.Method] = make(map[string]func(req *http.Request))
	}

	reg := registry.GetRegistry()
	sel := selector.NewSelector(selector.Registry(reg))

	rt.directors[policy][routeType][route.Method][route.Endpoint] = func(req *http.Request) {
		if route.Service != "" {
			// select next node
			next, err := sel.Select(route.Service)
			if err != nil {
				rt.logger.Error().Err(err).
					Str("service", route.Service).
					Msg("could not select service from the registry")
				return // TODO error? fallback to target.Host & Scheme?
			}
			node, err := next()
			if err != nil {
				rt.logger.Error().Err(err).
					Str("service", route.Service).
					Msg("could not select next node")
				return // TODO error? fallback to target.Host & Scheme?
			}
			req.URL.Host = node.Address
			req.URL.Scheme = node.Metadata["protocol"] // TODO check property exists?

		} else {
			req.URL.Host = target.Host
			req.URL.Scheme = target.Scheme
		}

		// Apache deployments host addresses need to match on req.Host and req.URL.Host
		// see https://stackoverflow.com/questions/34745654/golang-reverseproxy-with-apache2-sni-hostname-error
		if route.ApacheVHost {
			req.Host = target.Host
		}

		req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}
	}
}
func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

func queryRouteMatcher(endpoint string, target url.URL) bool {
	u, _ := url.Parse(endpoint)
	if !strings.HasPrefix(target.Path, u.Path) || endpoint == "/" {
		return false
	}
	q := u.Query()
	if len(q) == 0 {
		return false
	}
	tq := target.Query()
	for k := range q {
		if q.Get(k) != tq.Get(k) {
			return false
		}
	}
	return true
}

func (rt Router) regexRouteMatcher(pattern string, target url.URL) bool {
	matched, err := regexp.MatchString(pattern, target.String())
	if err != nil {
		rt.logger.Warn().Err(err).Str("pattern", pattern).Msg("regex with pattern failed")
	}
	return matched
}

func prefixRouteMatcher(prefix string, target url.URL) bool {
	return strings.HasPrefix(target.Path, prefix) && prefix != "/"
}

func SetDirectorFunc(parent context.Context, fn func(*http.Request)) context.Context {
	return context.WithValue(parent, directorCtxKey, fn)
}

// DirectorFunc gets the director function from the context.
func DirectorFunc(ctx context.Context) func(*http.Request) {
	val := ctx.Value(directorCtxKey)
	return val.(func(*http.Request))
}
