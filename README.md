欢迎进入 金山卫士开源计划 !在这里:
- 可以接触到中国最专业的安全类软件源代码;
- 可以自由的使用／研究／修订／再发布 这些代码以及延伸作品;

进一步的详细信息请访问:
  http://code.ijinshan.com/


#   URIsAok ~ 网址安全性辅助检验接口服务

## 概述
基于 http://code.ijinshan.com/api/devmore4.html 文档描述的金山云安全网址检验服务,包装为更加易用的在线RESTful 接口
进一步的:

    +-- 提供 FireFox 插件
    +-- 提供 Chrome 插件
    +-- 提供 本地批量扫描脚本
    +-- ...


## 发布

### http://py.kingsoft.net/=/

- 当前接口: http://py.kingsoft.net/=/chk
    
    - POST 想查询的 URL
    - 从金山云安全服务 返回是否安全


### http://urisaok.no.de

- 当前接口: http://urisaok.no.de/chk
    
    - POST 想查询的 URL
    - 从金山云安全服务 返回是否安全

- 增强接口: http://urisaok.no.de/qchk
    
    - POST 想查询的 URL
    - 先从本地 MongoDB 查询是否检验过,如果有,直接返回原有结论
    - 从金山云安全服务 返回是否安全



### http://urisaok.sinaapp.com

- 当前接口: http://urisaok.sinaapp.com/chk
    
    - POST 想查询的 URL
    - 从金山云安全服务 返回是否安全


### urisaok.appsp0t.com

- 当前接口: http://urisaok.appsp0t.com/chk
    
    - POST 想查询的 URL
    - 从金山云安全服务 返回是否安全

## 仓库


### Git
cnodejs.org 发布的 NAE 云平台支持 git 仓库的同步,
以及 Joyent 的 no.de 虚拟机也是以 git 仓库作为工作介质的,,,

所以:

https://github.com/ZoomQuiet/urisaok

    +-- master  供给 NAE
    +-- no.de   供给 no.de
    +-- openresrty   供给 嵌入Nginx 的本地服务


### Hg
以py 为主要开发工具的成果收集在 bitbucket.org 以 Mercurial 进行控制:

- 利用Hg 的分支功能,针对不同目的的代码进行了组织:
- https://bitbucket.org/ZoomQuiet/ok.urisa

```
    +-- default     默认分支包含主要的文档
    +-- GAE         可部署在GAE上的版本
    +-- SAE         可部署在SAE上的版本
    +-- NAE         可部署在NAE上的版本
    +-- bootle      本地的web srv. 版本
    +-- tools       周边应用工具
```


## 记要

    120306  Zoom.Quiet 创建 openresty 分支
    120214  Zoom.Quiet 和no.de 生产仓库同步
    120212  Zoom.Quiet 迁移到 github
    110323  Zoom.Quiet 创建

