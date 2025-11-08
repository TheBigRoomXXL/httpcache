#!/bin/sh

if [ -z "$1" ]; then
  echo "Missing version tag."
  echo "USAGE: $0 <VERSION>"
  exit 1
fi

NEW_VERSION=$1

git tag core/"$NEW_VERSION"
git tag diskv/"$NEW_VERSION"
git tag memcache/"$NEW_VERSION"
git tag otter/"$NEW_VERSION"
git tag redis/"$NEW_VERSION"
git tag storagetest/"$NEW_VERSION"

git push origin \
  core/"$NEW_VERSION" \
  otter/"$NEW_VERSION" \
  diskv/"$NEW_VERSION" \
  memcache/"$NEW_VERSION" \
  redis/"$NEW_VERSION" \
  storagetest/"$NEW_VERSION"
