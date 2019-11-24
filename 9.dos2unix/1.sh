#!/bin/bash
cd /data/webroot/www/ServerXianyu
cid=470f01636aeba03a0577124ed70c36e7bfd9745f
git fetch --tags --progress origin +refs/heads/*:refs/remotes/origin/*
if [ $? -ne 0 ]; then
echo "Error : git fetch"
exit -1
fi
cid_test=`git rev-list --no-walk ${cid}`
if [ $? -ne 0 ]; then
echo "Error : check commit id"
exit -2
fi
if [ "${cid_test}" != "${cid}" ]; then
echo "Error : commit id not same"
exit -3
fi
git checkout -f ${cid}
if [ $? -ne 0 ]; then
echo "Error : checkout commit"
exit -4
fi
echo ""
git log -3 --pretty=format:'%h - %s (%an, %ad)' --date=local --stat
echo ""
echo "当前分支 : " `git branch -r --contains ${cid}`

