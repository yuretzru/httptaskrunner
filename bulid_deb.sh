#!/bin/bash

mkdir -p deb/httptaskrunner/DEBIAN
echo "Package: httptaskrunner" > deb/httptaskrunner/DEBIAN/control
echo "Version: 1.0" >> deb/httptaskrunner/DEBIAN/control
MACHINE_TYPE=`uname -m`
if [ ${MACHINE_TYPE} == 'x86_64' ]; then
    echo "Architecture: amd64" >> deb/httptaskrunner/DEBIAN/control
else
    echo "Architecture: i386" >> deb/httptaskrunner/DEBIAN/control
fi
echo "Maintainer: root <root@localhost>" >> deb/httptaskrunner/DEBIAN/control
echo "Priority: extra" >> deb/httptaskrunner/DEBIAN/control
echo "Description: Http Task Runner" >> deb/httptaskrunner/DEBIAN/control
echo " ." >> deb/httptaskrunner/DEBIAN/control
mkdir deb/httptaskrunner/sbin
cp ./bin/httptaskrunner deb/httptaskrunner/sbin/
mkdir -p deb/httptaskrunner/usr/lib/systemd/system
cp ./bin/httptaskrunner.service deb/httptaskrunner/usr/lib/systemd/system/
cd deb
fakeroot dpkg-deb --build httptaskrunner
mv httptaskrunner.deb ../bin/httptaskrunner_`getconf LONG_BIT`.deb
cd ..
rm -rf deb
