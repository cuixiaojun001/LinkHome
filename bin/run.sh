#!/bin/bash
# 线上运行时, 默认按本脚本拉起进程;
#
# sh run.sh run # 运行脚本;
# sh run.sh stop # 停止脚本;
#
WORK_DIR=`dirname $0`
cd $WORK_DIR/..
WORK_DIR=`pwd`

BIN_NAME=`ls ${WORK_DIR}/bin/*.exe | head -n 1`

kill_process()
{
    ps -ef | grep $BIN_NAME | grep -v grep
    ps -ef | grep $BIN_NAME | grep -v grep |  awk '{print $2}' | xargs kill -9
}

run_process()
{
#    if [ -z "${ENVTYPE}" ]; then
#        echo "ENVTYPE 变量不存在或没有值"
#        exit 1
#    fi
    $BIN_NAME -f ${WORK_DIR}/conf/dev.yaml
}

start_process()
{
    if [ -z "${ENVTYPE}" ]; then
        echo "ENVTYPE 变量不存在或没有值"
        exit 1
    fi
    nohup $BIN_NAME -f ${WORK_DIR}/conf/${ENVTYPE}.yaml >> /data0/www/logs/linkhome.runtime.log &
}

Main()
{
    if [ "$1" == "run" ]; then
        run_process
    elif [ "$1" == "start" ];then
    	start_process
    elif [ "$1" == "stop" ];then
        kill_process
    fi
}

Main $@
