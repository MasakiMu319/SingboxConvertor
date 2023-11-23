# SingboxConvertor

![last-commit](https://img.shields.io/github/last-commit/MasakiMu319/SingboxConvertor?style=for-the-badge)
![license](https://img.shields.io/github/license/MasakiMu319/SingboxConvertor?style=for-the-badge)

Used to convert subscription links in Clash.Meta format to sing-box format (tested to work properly on Android, iOS, SFM).

If the generated configuration file is not available, check your configuration template.

## Deploy
The `port` environment variable controls the port on which the program runs. If not set, it opens on port 8080 by default.

## docker
```shell
docker volume create singboxconvertor    
docker run -d -p 8080:8080 -v singboxconvertor:/server/db jiumumu/singboxconvertor
# inspect log
docker logs -f $CONTAINER
```

## Usage
- After launching, use your browser to visit http://ip:port

- The New Profile in the sing-box Profiles fills in the remote link and allows you to switch nodes in Groups by starting a subscription.

- You need to ensure that the template configuration file you use contains the complete content about DNS and offloading, otherwise the converted configuration file will be incorrect and unusable.

- The converted subscription provides "HK", "TW", "JP", "SG", "US", "fallback" as the grouping of tags by default, and you can directly use them as outbound in the shunt rule.

- Groups of type urltest use Apple domains for HTTP latency testing.

## Template Profile
Most changes to the profile template will be preserved, as will adding nodes to the outbounds in the template.

## Support Protocol
- shadowsocks （Only support v2ray-plugin, obfs and shadow-tls plugin）
- shadowsocksR
- vmess
- vless (Include reality)
- trojan
- socks5
- http
- hysteria 1/2
- tuic5

# TO-DO list
- ~~web/frontend.html send some needless requests~~(Fixed in v0.2).
- Refactor the central portion of the codebase.
- ~~When copying the generated subscription address in the dev version, the request to generate the subscription may be sent repeatedly~~. 
- Scenarios where the server level fails to handle concurrent fetch requests correctly may result in lost requests or failure of the server level to issue subscription files correctly.
  - 服务端无法正确处理并发获取请求的场景，可能会导致请求丢失或者服务端无法正确下发订阅文件。
- Consider adding a configuration file upload function or providing a configuration file template.
  - 考虑支持上传配置文件或者服务端提供配置文件模板；

**v1.x 版本的开发计划：**
- 引入用户系统，用于支持用户维护自己的订阅和配置模板；
- 对配置文件部分进行重构，允许用户自行配置 DNS 和分流规则部分；
- DNS 提供的默认配置：海外走 Google/Cloudflare，国内走腾讯/阿里；
- 分流规则可能优先支持：DOMAIN-SUFFIX、DOMAIN-KEYWORD、GEOSITE、GEOIP，GEO*部分会直接提供列表进行勾选；
- 暂定允许用户自定义两个配置文件模板；
- 支持同后端的用户间分享配置模板。

# Change Log
- 0.1
  - Keep the core part of the code and use the gin framework.
  - (0.1.1) Fix container path issue.
- 0.2
  - Adjusted page styles to add more shadows and animations.
  - (0.2.2) Switch Gin to Release mode.
  - (0.2.3) Added more log information.
- 0.3
  - Refactor the project dependency structure and fix the Missing tag error caused by the filter node being empty.
- 0.4(dev)
  - Update index logo.
  - Refactor project front-end page structure.
  - Encrypt incoming subscription connections and external profiles.
  - Fixed an issue with replication subscriptions.
  - Fixed not allowing subscription generation if sub and configurl are empty.