#!/usr/bin/env bash

set -e

sudo launchctl start com.apple.rpcbind
# this is not needed as long as you mount with -o locallocks
sudo launchctl start com.apple.lockd

mkdir /tmp/nfs
