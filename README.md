# iptv_sc

first, get IPTV channel json file

wget http://iptv-src.byqwe.cn:23334/SCIPTV/channel/1/channel.json

then, run iptv_process

logo dir content created by command:

cat ../resource/shanxiunicom.m3u | awk '{print $4}' | grep 'tvg-logo' | sed 's/"/ /g' | awk '{print $2}' | xargs -n 1 wget -nc
