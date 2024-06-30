#!/bin/bash
docker build -t pingpong:latest .
DIGEST=$(docker images --no-trunc --quiet pingpong:latest | cut -d':' -f2)
echo $DIGEST > digest.txt
export DIGEST=$DIGEST
docker image tag pingpong:latest pingpong:$DIGEST
k3d image import pingpong:$DIGEST
sed "s/latest/$DIGEST/g" manifests/deployment_template.yaml > manifests/deployment.yaml
