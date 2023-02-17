#!/bin/bash
echo "fangsong@fs$ cp -r /home/fangsong/.leetcode/* /home/fangsong/CodeHub"
cp -r /home/fangsong/.leetcode/* /home/fangsong/CodeHub
echo "fangsong@fs$ cd /home/fangsong/CodeHub"
cd /home/fangsong/CodeHub
echo "fangsong@fs$ git add ."
git add .
echo "fangsong@fs$ git status"
git status
echo "fangsong@fs$ git commit -m \"feat: auto update\""
git commit -m "feat: auto update"
echo "fangsong@fs$ git push origin develop:develop"
git push origin main:main
echo "fangsong@fs$ git pull origin"
git pull origin