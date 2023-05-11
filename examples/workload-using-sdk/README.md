# Aegis

![Aegis](../../assets/aegis-icon.png "Aegis")

## Use Case: Leveraging Aegis SDK

This example demonstrates how to use **Aegis SDK** along with your workload.

When you use **Aegis SDK**, you can communicate with **Aegis Safe** directly
and fetch the secrets to your workload whenever you need them.

This approach provides a great deal of flexibility, enabling you to make 
customizations to your code as needed. While adding the **Aegis SDK** as a 
dependency may require some extra effort, it also opens up a range of
features and capabilities that will benefit your project in the long run.

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
make example-sdk-deploy-local

# Switch to this use case’s folder:
cd $WORKSPACE/aegis/examples/workload-using-sdk

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
make example-sdk-deploy

# Switch to this use case’s folder:
cd $WORKSPACE/aegis/examples/workload-using-sdk

# Register a secret:
./register.sh

# Tail the workload’s logs and verify that the secret is there:
./tail.sh
```









