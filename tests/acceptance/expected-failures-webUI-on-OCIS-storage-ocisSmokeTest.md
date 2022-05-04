## Scenarios from web tests that are expected to fail on OCIS with OCIS storage
The expected failures in this file are from features in the owncloud/web repo.

Lines that contain a format like "[someSuite.someFeature.feature:n](https://github.com/owncloud/web/path/to/feature)"
are lines that document a specific expected failure. Follow that with a URL to the line in the feature file in GitHub.
Please follow this format for the actual expected failures.

Level-3 headings should be used for the references to the relevant issues. Include the issue title with a link to the issue in GitHub.

Other free text and Markdown formatting can be used elsewhere in the document if needed. But if you want to explain something about the issue, then please post that in the issue itself.

Only the web scenarios tagged ocisSmokeTest are run by default in OCIS CI. This file lists the expected-failures of those ocisSmokeTest scenarios.

### [enable re-sharing is not possible](https://github.com/owncloud/ocis/issues/1743)
- [webUISharingFilePermissionMultipleUsers/shareFileWithMultipleUsers.feature:47](https://github.com/owncloud/web/blob/master/tests/acceptance/features/webUISharingFilePermissionMultipleUsers/shareFileWithMultipleUsers.feature#L47)
- [webUISharingFilePermissionMultipleUsers/shareFileWithMultipleUsers.feature:48](https://github.com/owncloud/web/blob/master/tests/acceptance/features/webUISharingFilePermissionMultipleUsers/shareFileWithMultipleUsers.feature#L48)
- [webUISharingFilePermissionMultipleUsers/shareFileWithMultipleUsers.feature:49](https://github.com/owncloud/web/blob/master/tests/acceptance/features/webUISharingFilePermissionMultipleUsers/shareFileWithMultipleUsers.feature#L49)
- [webUISharingFilePermissionMultipleUsers/shareFileWithMultipleUsers.feature:50](https://github.com/owncloud/web/blob/master/tests/acceptance/features/webUISharingFilePermissionMultipleUsers/shareFileWithMultipleUsers.feature#L50)
- [webUISharingFilePermissionMultipleUsers/shareFileWithMultipleUsers.feature:51](https://github.com/owncloud/web/blob/master/tests/acceptance/features/webUISharingFilePermissionMultipleUsers/shareFileWithMultipleUsers.feature#L51)
- [webUISharingFilePermissionMultipleUsers/shareFileWithMultipleUsers.feature:52](https://github.com/owncloud/web/blob/master/tests/acceptance/features/webUISharingFilePermissionMultipleUsers/shareFileWithMultipleUsers.feature#L52)

### [name of public link is empty and not "Public link" when not specified in the create request](https://github.com/owncloud/ocis/issues/1237)
- [webUISharingPublicBasic/publicLinkCreate.feature:11](https://github.com/owncloud/web/blob/master/tests/acceptance/features/webUISharingPublicBasic/publicLinkCreate.feature#L11)
- [webUISharingPublicBasic/publicLinkCreate.feature:28](https://github.com/owncloud/web/blob/master/tests/acceptance/features/webUISharingPublicBasic/publicLinkCreate.feature#L28)

### [Copy private link option not available](https://github.com/owncloud/ocis/issues/1409)
- [webUIPrivateLinks/accessingPrivateLinks.feature:9](https://github.com/owncloud/web/blob/master/tests/acceptance/features/webUIPrivateLinks/accessingPrivateLinks.feature#L9)
- [webUIPrivateLinks/accessingPrivateLinks.feature:17](https://github.com/owncloud/web/blob/master/tests/acceptance/features/webUIPrivateLinks/accessingPrivateLinks.feature#L17)

### [name of public link is empty and not "Public link" when not specified in the create request](https://github.com/owncloud/ocis/issues/1237)
- [webUISharingPublicDifferentRoles/shareByPublicLinkDifferentRoles.feature:33](https://github.com/owncloud/web/blob/master/tests/acceptance/features/webUISharingPublicDifferentRoles/shareByPublicLinkDifferentRoles.feature#L33)
- [webUISharingPublicDifferentRoles/shareByPublicLinkDifferentRoles.feature:34](https://github.com/owncloud/web/blob/master/tests/acceptance/features/webUISharingPublicDifferentRoles/shareByPublicLinkDifferentRoles.feature#L34)
- [webUISharingPublicDifferentRoles/shareByPublicLinkDifferentRoles.feature:35](https://github.com/owncloud/web/blob/master/tests/acceptance/features/webUISharingPublicDifferentRoles/shareByPublicLinkDifferentRoles.feature#L35)

### [impossible to navigate into a folder in the trashbin](https://github.com/owncloud/web/issues/1725)
- [webUITrashbinDelete/trashbinDelete.feature:29](https://github.com/owncloud/web/blob/master/tests/acceptance/features/webUITrashbinDelete/trashbinDelete.feature#L29)

### [Sharing seems to work but does not work](https://github.com/owncloud/ocis/issues/1303)
- [webUISharingInternalUsers/shareWithUsers.feature:52](https://github.com/owncloud/web/blob/master/tests/acceptance/features/webUISharingInternalUsers/shareWithUsers.feature#L52)
- [webUISharingInternalUsers/shareWithUsers.feature:53](https://github.com/owncloud/web/blob/master/tests/acceptance/features/webUISharingInternalUsers/shareWithUsers.feature#L53)
- [webUISharingInternalUsers/shareWithUsers.feature:54](https://github.com/owncloud/web/blob/master/tests/acceptance/features/webUISharingInternalUsers/shareWithUsers.feature#L54)