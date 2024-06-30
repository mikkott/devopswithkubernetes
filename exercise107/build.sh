#!/bin/bash
docker build -t logoutput:latest .
DIGEST=$(docker images --no-trunc --quiet logoutput:latest | cut -d':' -f2)
echo $DIGEST > digest.txt
export DIGEST=$DIGEST
docker image tag logoutput:latest logoutput:$DIGEST
k3d image import logoutput:$DIGEST
sed "s/latest/$DIGEST/g" manifests/deployment_template.yaml > manifests/deployment.yaml
