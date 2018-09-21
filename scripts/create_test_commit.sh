#!/bin/sh
set -e

echo "creating test commit"

# Delete the test branch if it already exists
git branch -D test | true

git checkout -b test

echo "test" > test.txt
git add test.txt
git commit -m 'Add test.txt'
