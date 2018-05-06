#!/bin/bash
export GOPATH=/go
export GOROOT=/usr/local/go
export CV_TEAM=team-vulture
BASE_PATH="${GOPATH}/src/bitbucket.org/${CV_TEAM}"
mkdir -p ${BASE_PATH}
export IMPORT_PATH="${BASE_PATH}/${BITBUCKET_REPO_SLUG}"
ln -s ${PWD} ${IMPORT_PATH}
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH