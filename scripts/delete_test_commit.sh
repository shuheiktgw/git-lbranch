#!/bin/sh
set -e

echo "deleting test commit"

git checkout master
git branch -D test
