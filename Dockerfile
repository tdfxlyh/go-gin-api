# 使用官方的 Golang 镜像作为基础镜像
FROM golang:1.19

# 镜像作者
MAINTAINER tdfxlyh

# 设置工作目录
# 容器内创建 myproject 文件夹
ENV APP_HOME=/home/lyh/code/go-gin-api
RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME

# 给golang设置代理
RUN go env -w GOPROXY=https://goproxy.io,direct

# 将当前目录加入到工作目录中（. 表示当前目录）
ADD . $APP_HOME

# 构建 Go 项目
RUN go build -o main

# 暴露一个端口
EXPOSE 8080

# 启动应用程序 # 是个坑，项目名不行就用main试试
CMD ["./main"]



# 1.进入项目根目录
# 2.打包成镜像
# docker build -t go-gin-api .
# 3.登录dockerhub
# docker login -u tdfxlyh
# 4.新建一个tag，名字必须跟你注册账号一样
# docker tag go-gin-api tdfxlyh/go-gin-api:1.0.1
# 5.推上去
# docker push tdfxlyh/go-gin-api:1.0.1


# 远程服务器拉取
# docker pull tdfxlyh/go-gin-api:1.0.1
# docker run -d -p 8080:8080 tdfxlyh/go-gin-api:1.0.1
# docker run -p 8080:8080 tdfxlyh/go-gin-api:1.0.1
# 下面是本机试试
# docker run -p 8080:8080 go-gin-api



# 一些简单的命令
# 1.列出本地的镜像
# docker images [OPTIONS]
# 2.删除某个镜像
# docker rmi [-f] 镜像名字
# 3.查看容器
# docker ps -n 15
# 4.启动已经停止运行的容器
# docker start 容器ID(容器名)
# 5.重启容器
# docker restart 容器ID(容器名)
# 6.停止容器
# docker stop 容器ID(容器名)
# 7.强制停止容器
# docker kill 容器ID(容器名)
# 8.删除已停止的容器
# docker rm [-f] 容器ID(容器名)

# 数据库容器
# docker run -d -p 3306:3306 --privileged=true -v /home/lyh/app/db/mysql/log:/var/log/mysql -v /home/lyh/app/db/mysql/data:/var/lib/mysql -v /home/lyh/app/db/mysql/conf:/etc/mysql/conf.d -e MYSQL_ROOT_PASSWORD=991113 -e TZ=Asia/Shanghai --name=mysql mysql:5.7
