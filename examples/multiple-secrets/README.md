# Aegis

![Aegis](../../assets/aegis-icon.png "Aegis")

## Use Case: Leveraging Aegis SDK

This example demonstrates how to use **Aegis SDK** to register more than one
secret to your workload.

This demo is a slight variation of the 
[Registering Secrets Using Init Container](../using-init-container)
example.

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
images; or you can pull and use pre-deployed images from Docker Hub.

The following sections describe both of these approaches respectively.

## Local Deployment

```bash
# Switch to the project folder:
cd $WORKSPACE/aegis 
# Build everything locally:
make build-local
# Deploy the use case:
make example-multiple-secrets-deploy-local
# Switch to this use case’s folder:
cd $WORKSPACE/aegis/examples/multiple-secrets
# Register a secret:
./register.sh
# List the secrets.
./secrets.sh
```

## Using Pre-Deployed Images

If you don’t want to build the container images locally, you can use
pre-deployed container images.

```bash 
# Switch to the project folder:
cd $WORKSPACE/aegis 
# Deploy the use case from the pre-built image.
make example-multiple-secrets-deploy
# Switch to this use case’s folder:
cd $WORKSPACE/aegis/examples/multiple-secrets
# Register a secret:
./register.sh
# List the secrets.
./secrets.sh
```
