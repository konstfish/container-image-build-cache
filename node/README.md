`docker build --no-cache -f Dockerfile.cache -t nodeapp-cache .`

26s
`docker build --no-cache -t nodeapp .`

10s
`docker build --no-cache --build-arg="BUILD_IMAGE=nodeapp-cache" --build-arg="BUILD_TAG=latest" -t nodeapp .`