#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

# Create a proxy from user’s localhost to Docker for Mac’s docker
# registry’s API port.
mac-tunnel:
	./hack/mac-registry-tunnel.sh