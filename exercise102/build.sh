#!/bin/bash
docker build -t project:latest .
k3d image import project:latest