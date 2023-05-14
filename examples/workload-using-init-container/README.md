# Aegis

![Aegis](../../assets/aegis-icon.png "Aegis")

## Use Case: Using an Init Container

This example demonstrates how to use **Aegis Init Container** to wait for 
secrets to be allocated to a workload before bootstrapping the workload.

## Video Tutorials Anyone?

[Watch **Aegis Showcase** to learn how to use **Aegis** hands-on][videos].

If any of the instructions provided here are unclear, the video tutorials will
help explain them in much greater detail. Each video is designed around a
particular topic to keep it concise and to-the-point.

The container image is also used as the **inspector** workload to debug secret
registration during showcasing various scenarios [in the workshop](../aegis-workshop).

[videos]: https://vimeo.com/showcase/10074951 "Aegis Showcase"

## Deployment Options

To follow this use case, you can either locally build and deploy the container
images; or you can pull and use pre-deployed images from Docker Hub. The
next two sections describe both approaches respectively.

## Local Deployment

```bash
# Switch to the project folder:
cd $WORKSPACE/aegis 
# Build everything locally:
make build-local
# Deploy the use case:
make example-init-container-deploy-local
# Switch to this use case’s folder:
cd $WORKSPACE/aegis/examples/workload-using-init-container
# Check and make sure that the workload pod is still initializing:
kubectl get po -n default
# Register a secret:
./register.sh
# Verify that the workload pod has initialized:
kubectl get po -n default
# Tail the workload’s logs and verify that the secret is there:
./tail.sh
```

## Using Pre-Deployed Images

If you don’t want to build the container images locally, you can use
pre-deployed container images.

```bash 
# Switch to the project folder:
cd $WORKSPACE/aegis 
# Deploy the use case from the pre-built image.
make example-sidecar-deploy
# Switch to this use case’s folder:
cd $WORKSPACE/aegis/examples/workload-using-sidecar
# Check and make sure that the workload pod is still initializing:
kubectl get po -n default
# Register a secret:
./register.sh
# Verify that the workload pod has initialized:
kubectl get po -n default
# Tail the workload’s logs and verify that the secret is there:
./tail.sh
```
