# !/bin/bash
if [  $(pwd) != "$HOME/spider/data" ] ;then
mkdir ~/spider || echo "spider目录已存在" || cd ~/spider || echo "进入spider目录失败"
mkdir ~/spider/pic || echo "pic目录已存在"|| cd ~/spider/pic || echo "进入data目录失败"

fi

downloadHome="$HOME/spider/pic"

curl $1 >> $downloadHome/$2

