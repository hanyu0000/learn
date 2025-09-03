### HTTP 1.0 和 2.0 什么区别?

问你 “能不能举个例子”，你可以说：

比如加载一个网页需要 100 张小图片，在 HTTP 1.0 下，浏览器需要开很多 TCP 连接，效率很低；而 HTTP 2.0 下，只需要一个连接，图片数据可以并行交错传输，不会互相等待，页面加载速度明显更快。

口述版答案：

HTTP1.0 每次请求都要新建 TCP 连接，请完就断开，效率低，还存在头部冗余、队头阻塞等问题。

HTTP2.0 主要解决性能瓶颈：

第一，它用 `多路复用`，一个 TCP 连接里可以并行多个请求和响应，不再互相阻塞；
第二，它把文本传输改成 `二进制分帧`，效率更高；
第三，它支持 `头部压缩`，大幅减少冗余数据；
第四，还支持 `服务端推送`，可以主动把资源发给客户端。

一句话总结：HTTP1.0 简单但低效，缺乏长连接、并发和压缩机制，而 HTTP2.0 通过多路复用、二进制分帧、头部压缩和服务端推送，大幅提升了性能和用户体验。

### HTTP 2.0 和 3.0 什么区别?

1. 协议层的根本区别

HTTP/2.0：基于 TCP。虽然解决了并发问题（多路复用、头部压缩等），但仍然受制于 TCP 本身的限制。

HTTP/3.0：基于 QUIC 协议（UDP 之上实现的可靠传输层），本质上绕开了 TCP 的局限。

1. 队头阻塞

HTTP/2.0：应用层解决了 HTTP 队头阻塞，但 TCP 层仍然有阻塞问题。如果一个数据包丢失，整个 TCP 流都要等待重传。

HTTP/3.0：QUIC 基于 UDP，每个流独立传输，即使某个流丢包，其他流也不会受影响，从根本上解决了队头阻塞。

3. 连接建立速度

HTTP/2.0：依赖 TCP + TLS，完整握手需要 3 RTT（TCP 三次握手 + TLS 握手）。

HTTP/3.0：QUIC 内置了 TLS（基于 TLS 1.3），只需 1 RTT，在缓存密钥时甚至可以 0 RTT 建连。

4. 连接迁移

HTTP/2.0：TCP 连接与四元组（源 IP、源端口、目标 IP、目标端口）绑定，如果换了网络（比如从 WiFi 切到 5G），需要重新建连。

HTTP/3.0：QUIC 用 Connection ID 识别连接，不依赖 IP/端口，切换网络时连接可以无缝迁移。

5. 其他优化

HTTP/2.0：二进制分帧、多路复用、头部压缩、服务端推送。

HTTP/3.0：继承了 HTTP/2 的特性，并通过 QUIC 提供更快、更稳定的传输，特别适合移动端和弱网环境。

口述版答案：

HTTP/2 是基于 `TCP` 的，它解决了应用层的队头阻塞、支持多路复用、头部压缩和服务端推送。但 TCP 本身有局限，比如丢一个包会阻塞整个连接，而且建连需要 TCP 三次握手再加 TLS 握手，延迟比较高。

HTTP/3 则基于 QUIC 协议，也就是跑在 `UDP` 上。QUIC 把 TLS1.3 内置进去，建连只要 1 RTT，甚至可以 0 RTT；同时它把流做成独立的，丢包只影响一个流，彻底解决了队头阻塞；而且用 Connection ID，可以在切换网络时保持连接不断，比如从 WiFi 换到 5G。

一句话总结：HTTP/2 优化了性能但受限于 TCP，而 HTTP/3 基于 QUIC，在弱网和移动端环境下体验更好，速度更快更稳定。

### 常见的HTTP状态码有那些?

#### 1xx（信息性响应）

100 Continue：客户端继续请求（常用于大文件上传，先发请求头，再确认是否继续发送 body）。

#### 2xx（成功）

200 OK：请求成功，最常见。

201 Created：资源已创建成功（比如注册新用户）。

204 No Content：请求成功，但没有返回内容。

#### 3xx（重定向）

301 Moved Permanently：永久重定向，资源位置已改变。

302 Found：临时重定向。

304 Not Modified：资源未修改，客户端可使用缓存。

#### 4xx（客户端错误）

400 Bad Request：请求报文有语法错误或参数错误。

401 Unauthorized：未授权，需要登录认证。

403 Forbidden：服务器理解请求，但拒绝执行（权限不足）。

404 Not Found：请求的资源不存在。

#### 5xx（服务端错误）

500 Internal Server Error：服务器内部错误，最常见。

502 Bad Gateway：网关或代理收到无效响应。

503 Service Unavailable：服务器暂时不可用（超载或维护）。

504 Gateway Timeout：网关或代理超时。

### HTTP请求包含那些内容，请求头和请求体有那些类型？

一、HTTP 请求由哪几部分组成？

一个 HTTP 请求通常包含 4 部分：

1.请求行（Request Line）

格式：<方法> <URL> <协议版本>

例如：GET /index.html HTTP/1.1

常见方法：GET、POST、PUT、DELETE、HEAD、OPTIONS、PATCH。

2.请求头（Request Headers）

键值对形式，描述客户端信息、资源需求、缓存策略等。

3.空行

用来分隔请求头和请求体。

4.请求体（Request Body）

在 GET 请求中一般为空；

在 POST/PUT 等方法中，通常包含提交的数据。

二、常见请求头类型

1.通用头（General Headers）

Host：目标主机和端口

Connection：是否保持长连接（keep-alive / close）

Cache-Control：缓存控制策略

2.请求头（Request Headers）

User-Agent：客户端软件信息

Accept：客户端能接收的内容类型（如 text/html, application/json）

Accept-Encoding：压缩方式（gzip, deflate, br）

Accept-Language：语言偏好

Referer：表示请求来源页面

Authorization：认证信息（如 Bearer token）

3.实体头（Entity Headers，描述请求体内容）

Content-Type：请求体类型（如 application/json、application/x-www-form-urlencoded）

Content-Length：请求体长度

Content-Encoding：请求体使用的压缩方式

三、请求体常见类型

表单数据

application/x-www-form-urlencoded：键值对形式（默认 HTML 表单提交方式）。

multipart/form-data：支持文件上传。

结构化数据

application/json：最常见，用于 RESTful API。

application/xml：XML 格式（老系统或 SOAP 常见）。

二进制数据

application/octet-stream：通用二进制流，比如文件下载、上传。

面试口述版

HTTP 请求一般包含 请求行、请求头、空行和请求体。
请求行里有 方法、URL 和协议版本，比如 GET /index.html HTTP/1.1。

请求头常见的有：

通用的 Host、Connection、Cache-Control；

请求相关的 User-Agent、Accept、Authorization；

实体相关的 Content-Type、Content-Length。

请求体的常见类型主要有：

表单数据：application/x-www-form-urlencoded 和 multipart/form-data（支持文件上传）；

结构化数据：application/json、application/xml；

以及二进制流：application/octet-stream。

一句话总结：HTTP 请求 = 请求行 + 请求头 + 请求体，不同头和体的类型决定了客户端和服务端如何交互。

### HTTP 和 HTTPS 有什么区别？

1. 基本概念

HTTP：明文传输协议，不做加密。

HTTPS：HTTP + SSL/TLS 加密层，保证通信安全。

2. 传输安全

HTTP：所有数据都是明文传输，容易被窃听、篡改、伪造（中间人攻击）。

HTTPS：数据经过 信息加密 、完整性校验 、身份认证 ，保证安全。

#### HTTPS 是如何解决上面的三个风险的？

1. `混合加密`的方式实现信息的机密性，解决了窃听的风险。

2. `摘要算法`的方式来实现完整性，它能够为数据生成独一无二的「指纹」，指纹用于校验数据的完整性，解决了篡改的风险。

3. 将服务器公钥放入到`数字证书`中，解决了冒充的风险。

3. 端口号

HTTP 默认端口：80

HTTPS 默认端口：443

4. 连接过程

HTTP：只需要 TCP 三次握手即可进行 HTTP 的报文传输。

HTTPS：在 TCP 三次握手后，还要进行 SSL/TLS 握手，才可进入加密报文传输。

5. 性能

HTTP：无加密，性能开销小，但不安全。

HTTPS：多了 TLS 握手和加解密过程，开销稍大，但现在硬件加速普及，性能影响很小。

6. 证书

HTTP：不需要证书。

HTTPS：需要向 CA 申请数字证书（或自签名），用于验证服务器身份。

一句话总结（面试口述版）

HTTP 是明文传输，速度快但不安全，容易被窃听和篡改；

HTTPS 在 HTTP 基础上加了 SSL/TLS 加密，使用 443 端口，能保证加密、完整性和身份认证，不过需要数字证书，建立连接比 HTTP 多一次握手。

### HTTP 中 GET 和 POST 的区别是什么？

1. 语义

GET：用于从服务器获取指定的资源（幂等<多次执行相同的操作，结果都是相同的>，即多次调用结果相同，不会改变服务器状态）。

POST：用于根据请求负荷对指定的资源做出处理（非幂等，多次调用可能产生副作用）。

2. 参数位置

GET：参数放在 `URL` 的 查询字符串 中

POST：参数放在 `请求体` 中，支持更大更复杂的数据。

3. 安全性

GET：参数直接暴露在 URL 中，容易被缓存、记录、泄露（浏览器历史、日志）。

POST：参数放在请求体，相对更安全，但仍需配合 HTTPS 才能真正保证机密性。

4. 长度限制

GET：URL 长度受浏览器/服务器限制（通常 2K～8K 字符）。

POST：请求体大小理论上无限制（实际受服务器配置影响）。

5. 缓存 & 书签

GET：天然支持浏览器缓存、书签收藏、分享。

POST：默认不会被缓存，不能直接作为书签保存。

6. 幂等性 & 可重复性

GET：幂等，重复请求不会产生额外影响。

POST：非幂等，可能多次提交数据（如多次下单）。

一分钟面试口述版

GET 和 POST 的主要区别在于：

语义上：GET 用于获取资源，幂等；POST 用于提交或修改资源，非幂等。

参数位置：GET 把参数放在 URL，POST 放在请求体。

安全性：GET 参数会暴露在 URL，不适合敏感数据；POST 相对安全，但真正安全要靠 HTTPS。

其他：GET 有长度限制、可缓存可收藏；POST 没有长度限制，默认不缓存。

一句话总结：GET 用于获取、参数在 URL、可缓存幂等；POST 用于提交、参数在 Body、非幂等。

### WebSocket 与 HTTP 有什么区别？

HTTP 是典型的请求-响应协议，客户端必须先发起请求，服务器才能返回数据，属于半双工通信；WebSocket 在最初通过 HTTP 握手之后，会升级为一个持久的 TCP 连接，这个连接是全双工的，客户端和服务器都可以随时主动发送消息。

从性能上看，HTTP 每次请求都有比较大的头部开销，而 WebSocket 建立连接后通信只需要很小的帧头，开销更低，实时性更强。

所以，一般 API、网页加载用 HTTP，而像聊天、实时通知、游戏这类需要实时双向通信的场景，就会用 WebSocket。

### 服务端是如何解析 HTTP 请求的数据？（考察 HTTP 请求格式的了解程度）

1. 一个典型的 HTTP 请求由三部分组成：

请求行
```shell
METHOD URL VERSION

METHOD：请求方法（GET、POST、PUT、DELETE…）

URL：请求路径（可能带查询字符串）

VERSION：HTTP 版本（如 HTTP/1.1）
```

请求头
```shell
Host: example.com
User-Agent: curl/7.68.0
Content-Type: application/json
Content-Length: 123

每行一个 key: value

通过空行 \r\n\r\n 与请求体分隔
```

请求体

主要在 POST/PUT 请求中出现

数据格式可能是 JSON、表单、二进制等

2. 服务端解析流程

读取 TCP 数据流

HTTP 基于 TCP，服务器先接收原始字节流。

解析请求行

读取第一行，拆分出 METHOD / URL / VERSION。

URL 可能还包含查询字符串，需要进一步解析。

解析请求头

逐行读取，直到遇到空行 \r\n。

将每个 header 转为 key-value 存储（通常用 map 或字典）。

读取请求体（可选）

根据 Content-Length 或 Transfer-Encoding 确定请求体长度和方式。

读取对应字节，解析为应用层数据（如 JSON、表单数据等）。

处理请求并生成响应

根据 METHOD + URL 做路由分发

返回 HTTP 响应

示例:

HTTP 请求由请求行、请求头和可选的请求体组成。服务端先从 TCP 流读取数据，解析请求行获取方法和路径，然后按行解析请求头，遇到空行后判断是否有请求体，如果有就根据 Content-Length 或 Transfer-Encoding 读取请求体。解析完成后，服务端就可以根据方法和 URL 做路由处理并返回响应。