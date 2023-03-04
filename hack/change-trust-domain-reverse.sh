#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

cd ./install/k8s/spire/spire || exit
sed -i "s/example.com/aegis.ist/" ./spire-agent.yaml
sed -i "s/example.com/aegis.ist/" ./spire-server.yaml
sed -i "s/example.com/aegis.ist/" ./spire-controller-manager-config.yaml
cd ../..

cd ./demo-workload/using-sdk || exit
sed -i "s/example.com/aegis.ist/" ./Identity.yaml
sed -i "s/example.com/aegis.ist/" ./Deployment.yaml
cd .. || exit

cd ./using-sidecar || exit
sed -i "s/example.com/aegis.ist/" ./Identity.yaml
sed -i "s/example.com/aegis.ist/" ./Deployment.yaml
cd ../..

cd ./safe || exit
sed -i "s/example.com/aegis.ist/" ./Identity.yaml
sed -i "s/example.com/aegis.ist/" ./Deployment.yaml
cd .. || exit

cd ./sentinel || exit
sed -i "s/example.com/aegis.ist/" ./Identity.yaml
sed -i "s/example.com/aegis.ist/" ./Deployment.yaml
cd .. || exit

echo "Everything is awesome!"
