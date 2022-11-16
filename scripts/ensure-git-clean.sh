#!/bin/bash

if [ -n "$(git status --porcelain)" ]; then
  echo "Git status is not clean after installing and linting. Please double check lockfiles and your code formatting.";
  git status
  exit 1;
else
  echo "No changes detected after running i18n string extractor";
fi