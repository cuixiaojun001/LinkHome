FROM golang:1.19

# 容器的工作目录
WORKDIR /home/workspace

# 将当前目录下的所有文件拷贝到容器的工作目录下
COPY . /home/workspace

# 运行 build.sh 脚本
RUN /bin/bash build.sh

# 在容器中运行的命令
CMD ["bash", "release/bin/run.sh", "run"]
