#!/bin/bash

git rm --cached idea --ignore-unmatch
git rm --cached .idea/\* --ignore-unmatch
git commit -m "delete track files"
