# Aegis

![Aegis](../../assets/aegis-icon.png "Aegis")

## Use Case: Leveraging Aegis Sidecar

This example demonstrates how to use **Aegis Sidecar** along with your workload.

When you use **Aegis Sidecar**, you don’t need to modify your workload. 
**Aegis Sidecar** can fetch and provide the secrets that your workload needs
automatically.

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
make example-sidecar-deploy-local
# Switch to this use case’s folder:
cd $WORKSPACE/aegis/examples/workload-using-sidecar
# Register a secret:
./register.sh
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
# Register a secret:
./register.sh
# Tail the workload’s logs and verify that the secret is there:
./tail.sh
```
