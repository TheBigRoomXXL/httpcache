#!/bin/bash

if [ -z "$1" ]; then
  echo "Missing version tag."
  echo "USAGE: release.sh <VERSION>"
  exit 1
fi

$NEW_VERSION = $1

git tag core/$NEW_VERSION
git push origin core/$NEW_VERSION

(
    cd ./diskv 
    go get pkg.lovergne.dev/httpcache/core@$NEW_VERSION
)
(
    cd ./otter
    go get pkg.lovergne.dev/httpcache/core@$NEW_VERSION
)
(
    cd ./memcache
    go get pkg.lovergne.dev/httpcache/core@$NEW_VERSION
)
(
    cd ./redis
    go get pkg.lovergne.dev/httpcache/core@$NEW_VERSION
)

git add .
git commit -m "chore: bump httpcache/core version in cache implementation"

git tag diskv/$NEW_VERSION
git tag memcache/$NEW_VERSION
git tag otter/$NEW_VERSION
git tag redis/$NEW_VERSION
git push origin core/$NEW_VERSION otter/$NEW_VERSION diskv/$NEW_VERSION memcache/$NEW_VERSION redis/$NEW_VERSION
