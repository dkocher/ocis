package web

import (
	"embed"
)

//go:generate make generate
//go:embed assets/*
var Assets embed.FS
