docker login -u "$DOCKER_USER" -p "$DOCKER_PASS"
docker build -f build/Dockerfile -t "$DOCKER_IMAGE:latest" .
docker push "$DOCKER_IMAGE:latest"
