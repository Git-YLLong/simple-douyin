# simple-douyin

## 运行项目

- 启动docker镜像
```shell
docker-compose up
```
- 启动api
```shell
cd douyin/api
sh run.sh 
```
- 启动core基本服务
```shell
cd douyin/core 
sh build.sh 
sh output/bootstrap.sh
```
- 启动extra部分扩展服务
```shell
cd douyin/extra 
sh build.sh 
sh output/bootstrap.sh
```
- 启动http-server用于访问本地资源
```shell
http-server -p 8088
```
- 进行调试

### 功能说明

* 用户注册、登录
* 视频流推送
* 加载个人信息、投稿列表、喜欢列表
* 上传视频，保存至本地public目录中，通过http-server：8088端口访问
* 点赞视频

### 资料地址

- [抖音项目简介](https://bytedance.feishu.cn/docx/doxcnbgkMy2J0Y3E6ihqrvtHXPg)：包含项目总体要求和接口简要说明
- [各功能对应接口说明文档](https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18345145)：包含各接口的详细说明，请求参数与返回值
- [极简抖音APP使用说明](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7)：包含Apk包、接口与场景说明等

### 文件说明

- controller：demo业务代码，主要前期编写业务逻辑参考用
- douyin/api：客户端
    - handlers：路由绑定，获取前端参数
    - rpc：连接Client端，向Server端传递前端参数，获取处理结果
    - main.go：启动函数，绑定url接口
- douyin/core：基础接口服务
    - dal：数据库操作
    - pack：基本返回结果与雪花算法生成唯一id
    - service：服务调用操作数据库
    - handler.go：服务接口实现
    - main.go：服务启动函数，主要配置Sever参数
- douyin/extra：扩展接口服务
    - 目前只实现了点赞操作、点赞列表两个接口
- idl：采用thrift协议，通过kitex自动生成代码
    - core.thrift：基础接口
    - extra.thrift：扩展接口
- kitex_gen：kitex生成的中间代码
    - 第一次生成时用`kitex -module ... -service ... core.thrift`，生成完整的代码结构
    - 以后再修改thrift文件后只需要执行`kitex -module ... core.thrift`修改kitex_gen文件即可
- pkg
    - constants：一些常量定义
    - errno：自定义错误类型
    - jwt：鉴权，生成token与解析token（这个不知道要不要用gin的jwt扩展代替，先凑合用着吧）
    - middleware：kitex中间件，返回client、server运行信息
    - tracer：jaeger链路追踪
- public：存放本地视频
- docker-compose.yml：docker容器，包括mysql、etcd、jaeger，也可以在本地安装


