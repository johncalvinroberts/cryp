#!/bin/bash

if [ -n "$(git status --porcelain)" ]; then
  echo "Git status is not clean after installing and linting. Please double check lockfiles and your code formatting.";
  git status
  git diff ui/pnpm-lock.yaml
  exit 1;
else
  echo "No changes detected. You r gud.";
fi