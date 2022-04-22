### Docker常用指令

| 指令                                                         | 说明                                       |
| ------------------------------------------------------------ | ------------------------------------------ |
| docker run [-d] [-p \<port1\>:\<port2\> \<image>\] -\<ENV1\>=\<VAL1\> | 运行镜像，可指定是否后台运行，以及端口映射 |
| docker image ls [-a]                                         |                                            |
| docker image rm \<id/name\>                                  |                                            |
| docker image prune                                           | 清除无用的镜像                             |
| docker container ls [-a]                                     |                                            |
| docker container rm \<id/name\>                              |                                            |
| docker container prune                                       | 清除无用的容器                             |
| docker exec -i -t \<container_id\> /bin/sh                   | 使用终端方式进入容器                       |
| docker logs -t \<container_id\>                              | 查看容器日志                               |
|                                                              |                                            |

