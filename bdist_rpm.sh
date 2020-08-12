#!/bin/bash

here=$PWD

export BDISTRPMBASEDIR=$here/rpmbuild
rm -rf "${BDISTRPMBASEDIR}"
mkdir -p "${BDISTRPMBASEDIR}"/{BUILD,RPMS,SRPMS,SOURCES}

./buildrpmfromspec.sh

cd "$here" || exit
rm -rf dist
mkdir -p dist

find "$BDISTRPMBASEDIR" -regex '.*/RPMS/.*rpm' -not -name '*debuginfo*' -print0 | xargs -0 -I{} cp -a '{}' dist
