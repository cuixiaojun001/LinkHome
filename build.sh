#!/bin/sh

function print_usage()
{
	echo "
Usage:
	打包部署时,使用本脚本完成编译和打包

	sh build.sh
"
}

WORK_DIR=`dirname $0`
cd $WORK_DIR
WORK_DIR=`pwd`

if [[ $1 == "clean" ]];then
	rm -rf release application.tar.gz bin/*.exe conf/config.yaml
	exit
fi

if [[ ! -d $WORK_DIR/conf/ ]];then
	echo "error: 需要对应的配置文件!"
	print_usage
	exit 1
fi

export GOPROXY="https://goproxy.cn,direct"
go env
make
if [ $? -ne 0 ];then
	echo "make failed.×××××××××××"
	exit 0
fi
if [[ -d release ]]; then
	rm -rf release
fi

mkdir release
cp -R bin release
cp conf/*.yaml conf/
cp -R conf release

ls -R release
if [ $? -ne 0 ];then
	exit 0
fi
echo "========build finish========"
