## 代理服务

- socket5协议

> 指定一个端口 和地址（默认0.0.0.0）<br>
ssh -D 0.0.0.0:8021 teng@127.0.0.1 <br>
> VER NMETHODS  METHODS


### golang编写代理服务

- 监听地址
- 接受链接
- 建立到目标服务的链接
- 数据拷贝
- 关闭连接

### 加密

- 对称加密  (速度更快)
- 非对称加密 （公钥和私钥）

### go对称加密

- rc4 （不安全 速度快 小巧）