# Maintainer Jonathan Mainguy jon@soh.re
FROM centos:centos7
ADD repeatafterme /usr/sbin/repeatafterme
RUN mkdir -p /opt/repeatafterme
VOLUME /etc/repeatafterme/config.yml
CMD [ "/usr/sbin/repeatafterme" ]
