`docker build --no-cache -f Dockerfile.cache -t goapp-cache .`

31s
`docker build --no-cache --build-arg="COMPONENT=goapp1" -t goapp .`

8s
`docker build --no-cache --build-arg="COMPONENT=goapp1" --build-arg="BUILD_IMAGE=goapp-cache" --build-arg="BUILD_TAG=latest" -t goapp .`