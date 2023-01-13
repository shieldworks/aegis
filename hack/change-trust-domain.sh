#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

cd ./install/k8s/spire || exit
sed -i "s/aegis.z2h.dev/example.com/" ./spire-agent.yaml
sed -i "s/aegis.z2h.dev/example.com/" ./spire-server.yaml
sed -i "s/aegis.z2h.dev/example.com/" ./spire-controller-manager-config.yaml
cd ../../..

echo "Everything is awesome!"
