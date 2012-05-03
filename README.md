欢迎进入 金山卫士开源计划 !

在这里你可以接触到中国最专业的安全类软件源代码;
你可以自由的使用／研究／修订／再发布 这些代码以及延伸作品;

进一步的详细信息请访问:
  http://code.ijinshan.com/


#   URIsA ~ 在线网址安全性查询服务

## 概述
基于 http://code.ijinshan.com/api/devmore4.html 服务,包装为更加易用的在线RESTful 服务
进一步的:

    +-- 提供 FireFox 插件
    +-- 提供 Chrome 插件
    +-- 提供 本地批量扫描脚本
    +-- ...


## 发布

### http://urisago1.appsp0t.com

当前接口:

- http://urisago1.appsp0t.com
    
    - POST 想查询的 URL
    - 从金山云安全服务 返回是否安全



### http://urisaok.appsp0t.com

当前接口:

- http://urisaok.appsp0t.com
    
    - POST 想查询的 URL
    - 从金山云安全服务 返回是否安全


### http://py.kingsoft.net:8008/=/

当前接口:

- http://py.kingsoft.net:8008/=/
    
    - POST 想查询的 URL
    - 从金山云安全服务 返回是否安全


### http://urisaok.cnodejs.net

当前接口:

- http://urisaok.cnodejs.net/chk
    
    - POST 想查询的 URL
    - 从金山云安全服务 返回是否安全

- http://urisaok.cnodejs.net/qchk
    
    - POST 想查询的 URL
    - 检查本地 MongoDB 是否有历史结果
    - 否则 从金山云安全服务 返回是否安全



### http://urisaok.sinaapp.com

- 当前接口: http://urisaok.sinaapp.com/chk
    
    - POST 想查询的 URL
    - 从金山云安全服务 返回是否安全


### http://urisaok.no.de

当前接口:

- http://urisaok.no.de/chk
    
    - POST 想查询的 URL
    - 从金山云安全服务 返回是否安全

- http://urisaok.no.de/qchk
    
    - POST 想查询的 URL
    - 检查本地 MongoDB 是否有历史结果
    - 否则 从金山云安全服务 返回是否安全


## 仓库


### Git
cnodejs.org 发布的 NAE 云平台支持 git 仓库的同步,
以及 Joyent 的 no.de 虚拟机也是以 git 仓库作为工作介质的,,,

所以:

https://github.com/ZoomQuiet/urisaok

    +-- master  供给 NAE
    +-- no.de   供给 no.de
    +-- GAEv*   供给 GAE不同版本
    +-- openresty   供给 本地的 openresty lua 脚本

