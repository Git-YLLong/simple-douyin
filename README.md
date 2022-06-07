# simple-douyin

## 抖音项目服务端简单示例

项目启动：仿照easy-note项目启动即可

- 首先启动docker-compose
```shell
docker-compose up
```
- 再启动服务端
```shell
cd douyin/core 
sh build.sh 
sh output/bootstrap.sh
```
- 再启动客户端
```shell
cd douyin/api
sh run.sh 
```
- 进行调试

### 功能说明

接口功能不完善，仅作为示例

* 用户登录数据保存在内存中，单次运行过程中有效
* 视频上传后会保存到本地 public 目录中，访问时用 127.0.0.1:8080/static/video_name 即可

### 测试数据

测试数据写在 demo_data.go 中，用于列表接口的 mock 测试

### 资料地址

- [抖音项目简介](https://bytedance.feishu.cn/docx/doxcnbgkMy2J0Y3E6ihqrvtHXPg)：包含项目总体要求和接口简要说明
- [各功能对应接口说明文档](https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18345145)：包含各接口的详细说明，请求参数与返回值
- [极简抖音APP使用说明](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7)：包含Apk包、接口与场景说明等

### 文件说明

- controller：demo业务代码，主要前期编写业务逻辑参考用
- douyin/api：客户端
    - handlers：路由绑定，获取前端参数
    - rpc：连接客户端，向服务端传递前端参数
    - main.go：启动函数，绑定url接口
- douyin/core：基础接口服务端
    - dal：数据库操作
    - output：自动生成的，不用管
    - pack：一些工具包，雪花算法生成唯一id
    - script：运行build.sh自动生成
    - service：服务端业务调用操作数据库
    - build.sh：kitex自动生成
    - handler.go：服务接口实现
    - main.go：服务端启动函数
- idl：代码生成（照着接口文档写就行了）
    - core.thrift：基础接口
    - extra1.thrift：扩展接口1
    - extra2.thrift：扩展接口2
- kitex_gen：kitex生成的中间代码
    - 第一次生成时用`kitex -module -service core.thrift`，生成完整的代码结构
    - 以后再修改thrift文件后只需要执行`kitex core.thrift`修改kitex_gen文件即可
- pkg
    - constants：一些常量定义
    - errno：自定义错误类型
    - jwt：鉴权，生成token与解析token（这个不知道要不要用gin的jwt扩展代替，先凑合用着吧）
    - middleware：kitex中间件，返回client、server运行信息
    - tracer：jaeger链路追踪
- public：存放一些本地视频
- docker-compose.yml：docker容器，包括mysql、etcd、jaeger，也可以在本地安装
- douyin.sql：sql代码
- main.go：demo主函数，没用
- router.go：demo路由，没用
- test.http：我在本地测试的连接，不用管


