# frpManager

frpManger 是一个针对frp的多进程的管理工具，适应于多ip站群服务器，启动大量frp去绑定每个ip以产生映射。

### How to

创建配置文件`FrpNodeConfig.json`，然后创建多个frps的配置文件，将配置文件名添加到配置文件中就行。
例如有`node1-1.ini`、`node1-2.ini`、`node1-3.ini`三个配置文件。可以写成一下文件：

``` json
{
    "node1-1":"node1-1.ini",
    "node1-2":"node1-2.ini",
    "node1-3":"node1-3.ini"
}
```

> 请将本程序一定放到frps的当前目录

本程序为业余开发，目前只追求能用，后期更多功能看个人时间安排。