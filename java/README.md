`docker build --no-cache -f Dockerfile.cache -t bloat-cache .`

36s
`docker build --no-cache -t bloat .`

7s
`docker build --no-cache --build-arg="BUILD_IMAGE=bloat-cache" --build-arg="BUILD_TAG=latest" -t bloat .`