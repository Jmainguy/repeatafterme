%if 0%{?rhel} == 7
  %define dist .el7
%endif
%define _unpackaged_files_terminate_build 0
Name: repeatafterme
Version: 0.1
Release:    1%{?dist}
Summary: A golang daemon for retweeting / liking posts from twitter for a user provided list.

License: GPLv2
URL: https://github.com/Jmainguy/repeatafterme
Source0: repeatafterme.tar.gz
Requires(pre): shadow-utils

%description
A golang daemon for retweeting / liking posts from twitter for a user provided list.

%prep
%setup -q -n repeatafterme
%install
mkdir -p $RPM_BUILD_ROOT/usr/sbin
mkdir -p $RPM_BUILD_ROOT/opt/repeatafterme
mkdir -p $RPM_BUILD_ROOT/usr/lib/systemd/system
mkdir -p $RPM_BUILD_ROOT/etc/repeatafterme
install -m 0755 $RPM_BUILD_DIR/repeatafterme/repeatafterme %{buildroot}/usr/sbin
install -m 0644 $RPM_BUILD_DIR/repeatafterme/service/repeatafterme.service %{buildroot}/usr/lib/systemd/system
install -m 0644 $RPM_BUILD_DIR/repeatafterme/config.yaml %{buildroot}/etc/repeatafterme/

%files
/usr/sbin/repeatafterme
/usr/lib/systemd/system/repeatafterme.service
%dir /opt/repeatafterme
%dir /etc/repeatafterme
%config(noreplace) /etc/repeatafterme/config.yaml

%pre
getent group repeatafterme >/dev/null || groupadd -r repeatafterme
getent passwd repeatafterme >/dev/null || \
    useradd -r -g repeatafterme -d /opt/repeatafterme -s /sbin/nologin \
    -c "User to run repeatafterme service" repeatafterme
exit 0

%post
chown -R repeatafterme:repeatafterme /opt/repeatafterme
if [ -f /usr/bin/systemctl ]; then
  systemctl daemon-reload
fi

%changelog
