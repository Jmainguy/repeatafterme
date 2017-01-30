#!/bin/bash
docker kill repeatafterme
docker rm repeatafterme
docker run -d --name repeatafterme -v /opt/repeatafterme/config.yaml:/etc/repeatafterme/config.yaml:z repeatafterme:latest
