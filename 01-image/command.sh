#list aplikasi/image
docker image ls

# Download Image
docker image pull redis # cara 1
docker pull mysql # cara 2
docker pull postgres:latest # dengan tag

# Delete Image
docker image rm redis