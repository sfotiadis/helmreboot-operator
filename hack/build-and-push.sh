#!/bin/bash

# Build and Push Docker Image Script
# Usage: ./build-and-push.sh [tag]

set -e

# Default values
DEFAULT_TAG="latest"
DOCKER_REPO="sfotia2s/helmreboot-operator"

# Get tag from argument or use default
TAG=${1:-$DEFAULT_TAG}

# Extract version from Chart.yaml if no tag provided
if [ "$TAG" = "latest" ]; then
    CHART_VERSION=$(grep '^appVersion:' chart/Chart.yaml | cut -d' ' -f2 | tr -d '"')
    if [ -n "$CHART_VERSION" ]; then
        TAG=$CHART_VERSION
    fi
fi

echo "Building Docker image: ${DOCKER_REPO}:${TAG}"

# Build the image
docker build -t "${DOCKER_REPO}:${TAG}" .

# Also tag as latest
docker tag "${DOCKER_REPO}:${TAG}" "${DOCKER_REPO}:latest"

echo "Image built successfully!"
echo "Image: ${DOCKER_REPO}:${TAG}"
echo "Image: ${DOCKER_REPO}:latest"

# Ask if user wants to push
read -p "Do you want to push the image to Docker Hub? (y/N): " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    echo "Pushing to Docker Hub..."
    docker push "${DOCKER_REPO}:${TAG}"
    docker push "${DOCKER_REPO}:latest"
    echo "Images pushed successfully!"
    echo ""
    echo "To use this image:"
    echo "   docker pull ${DOCKER_REPO}:${TAG}"
    echo ""
    echo "In your HelmRelease:"
    echo "   values:"
    echo "     image:"
    echo "       repository: ${DOCKER_REPO}"
    echo "       tag: ${TAG}"
else
    echo "Skipping push. Images are available locally."
    echo ""
    echo "To push manually:"
    echo "   docker push ${DOCKER_REPO}:${TAG}"
    echo "   docker push ${DOCKER_REPO}:latest"
fi