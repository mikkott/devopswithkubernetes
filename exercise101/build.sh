#!/bin/bash
docker build -t logoutput:latest .
k3d image import logoutput:latest