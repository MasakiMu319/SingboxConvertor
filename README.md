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
