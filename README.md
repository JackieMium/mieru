# mieru / 見える

mieru【見える】是一款安全的、无流量特征、无法被主动探测的，基于 UDP 和 KCP 协议的 socks5 网络代理软件。

mieru 代理软件由称为 mieru【見える】的客户端软件和称为 mita【見た】的代理服务器软件这两部分组成。

## 原理

mieru 的翻墙原理与 shadowsocks / v2ray 等软件类似，在客户端和墙外的代理服务器之间建立一个加密的通道。GFW 不能破解加密传输的信息，无法判定你最终访问的网址，因此只能选择放行。

## 特性

1. 使用高强度的 AES-256-GCM 加密算法，基于用户名、密码和系统时间生成密钥。以现有计算能力，mieru 传输的数据内容无法被破解。
2. mieru 实现了客户端和代理服务器之间所有传输内容的完整加密，不传输任何明文信息。网络观察者（例如 GFW）仅能获知时间、数据包的发送和接收地址，以及数据包的大小。除此之外，观察者无法得到其它任何流量信息。
3. 当 mieru 发送数据包时，会在尾部填充随机信息。即便是传输相同的内容，数据包大小也不相同。这从根本上解决了 KCP 协议的 ACK 包容易被识别的问题。
4. mieru 不需要客户端和服务器进行握手，即可直接发送数据。当服务器无法解密客户端发送的数据时，不会返回任何内容。因此 GFW 不能通过主动探测发现 mieru 服务。
5. mieru 支持多个用户共享代理服务器。
6. 客户端软件支持 Windows，Mac OS 和 Linux 系统。

## 使用教程

1. [服务器安装与配置](https://github.com/enfein/mieru/blob/main/docs/server-install.md)
2. [客户端安装与配置](https://github.com/enfein/mieru/blob/main/docs/client-install.md)
3. [运营维护与故障排查](https://github.com/enfein/mieru/blob/main/docs/operation.md)

## 编译

编译 mieru 的客户端和服务器软件，建议在 Debian/Ubuntu 系列发行版的 Linux 系统中进行。编译过程可能需要翻墙下载依赖的软件包。

编译所需的软件包括：

- curl
- env
- git
- go (version >= 1.15)
- sha256sum
- tar
- zip

编译服务器 debian 安装包还需要：

- dpkg-deb
- fakeroot

编译时，进入项目根目录，调用指令 `./build.sh` 即可。编译结果会存放在项目根目录下的 `release` 文件夹。

## 贡献

出于安全考虑，mieru 实际开发时使用的 git 仓库是一个私有仓库。我们在发布新版本前，会把私有仓库内的改动合并起来移动到这里。由于在两个仓库之间双向同步代码比较困难，我们暂时不接受 pull request。如果有新需求或报告 bug 请提交 GitHub Issue。

## 联系作者

关于本项目，如果你有任何问题，请提交 GitHub Issue 联系我们。

## 许可证

使用本软件需遵从 GPL 第三版协议。
