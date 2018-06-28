# Gen - 为微服务生成源码的工具

只需要写普通的函数,Gen会为其生成高效的路由源代码和文档  
因为生成的是源码，所以这些都不会影响运行时性能  
工具中的每个变化所引起的差异直接显示在生成的源代码中  
也支持生成客户端  

- [English](https://github.com/wzshiming/gen/blob/master/README.md)
- [简体中文](https://github.com/wzshiming/gen/blob/master/README_cn.md)

## 支持的功能

- [X] 生成路由
- [X] 生成客户端
  - [X] Golang
  - [ ] JavaScript (For Web)
  - [ ] Java (For Android)
  - [ ] Swift (For iOS)
- [X] 生成文档
  - [ ] Swagger 2
  - [X] [OpenAPI 3](https://github.com/OAI/OpenAPI-Style-Guide)
  - [X] [SwaggerUI](https://github.com/swagger-api/swagger-ui)
- [X] 协议
  - [X] HTTP
  - [ ] Protobuf

## 示例

[路由](https://github.com/wzshiming/gen/blob/master/examples/route1/)  
[客户端](https://github.com/wzshiming/gen/blob/master/examples/client1/)  

## 许可证

软包根据MIT License。有关完整的许可证文本，请参阅[LICENSE](https://github.com/wzshiming/gen/blob/master/LICENSE)。