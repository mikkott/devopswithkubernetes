#!/bin/bash
docker build -t project:latest .
DIGEST=$(docker images --no-trunc --quiet project:latest | cut -d':' -f2)
echo $DIGEST > digest.txt
export DIGEST=$DIGEST
docker image tag project:latest project:$DIGEST
k3d image import project:$DIGEST
sed -i "s/latest/$DIGEST/g" manifests/deployment.yaml
