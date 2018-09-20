#!/bin/sh
set -e

echo "creating test commit"

git checkout -b test

echo "test" > test.txt
git add test.txt
git commit -m 'Add test.txt'
