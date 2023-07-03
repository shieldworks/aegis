#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

# The common version tag assigned to all the things.
VERSION=0.17.4

# Utils
include ./AegisMacOs.mk
include ./AegisDeploy.mk
## Aegis
include ./AegisSafe.mk
include ./AegisSentinel.mk
include ./AegisInitContainer.mk
include ./AegisSidecar.mk
## Examples
include ./AegisExampleSidecar.mk
include ./AegisExampleSdk.mk
include ./AegisExampleMultipleSecrets.mk
include ./AegisExampleInitContainer.mk

## Build
include ./AegisBuild.mk

## Help
include ./AegisHelp.mk
