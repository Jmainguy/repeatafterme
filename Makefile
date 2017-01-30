# This expects a few requirements
# one, that https://github.com/Jmainguy/docker_rpmbuild is cloned into ~/Github/docker_rpmbuild
# two, that docker is installed and running
# three, that ~/Github/docker_rpmbuild/dockerbuild/build.sh centos7 has been run
rpm:
	@go build
	@tar -czvf ~/Github/docker_rpmbuild/rpmbuild/SOURCES/repeatafterme.tar.gz ../repeatafterme
	@cp repeatafterme.spec ~/Github/docker_rpmbuild/rpmbuild/SPECS/repeatafterme.spec
	@cd ~/Github/docker_rpmbuild/; ./run.sh centos7 repeatafterme
	@ls -ltrh ~/Github/docker_rpmbuild/rpmbuild/RPMS/x86_64/repeatafterme*

docker:
	@go build
	@mv repeatafterme packaging/docker/
	@cd packaging/docker
	@docker build -t="repeatafterme" packaging/docker
	@if [ ! -d /opt/repeatafterme ]; then mkdir /opt/repeatafterme; fi
	@cp config.yaml /opt/repeatafterme/
	@cp packaging/docker/run.sh /opt/repeatafterme/
	@echo "cd to /opt/repeatafterme, edit config.yaml, run run.sh"
