![Aegis](../../assets/aegis-banner.png "Aegis")

## Metrics and Monitoring

These manifests can be used to deploy metrics and monitoring support to 
the cluster.

For minikube, the following commands may help too:

```bash 
minikube addons list 
minikube addons enable metrics-server
```

## Registry Addon 

This is useful for testing changes locally before pushing them to dockerhub.

That will ensure that only tested and verified images are pushed to dockerhub.

Here are quick instructions:

Start Minikube: minikube start

Enable registry add-on: minikube addons enable registry

Build a Docker image and tag it for the Minikube registry:

```bash
docker build -t myimage .
docker tag myimage "$(minikube ip):5000/myimage"
```

Push the image to the Minikube registry: 

```bash
docker push "$(minikube ip):5000/myimage"
```

Deploy the image to Minikube:

```bash
kubectl run myapp --image="$(minikube ip):5000/myimage"
```