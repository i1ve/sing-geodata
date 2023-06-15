# sing-geodata

Generate geodata for sing-box

## Download

- **geoip.db**:
  - [https://github.com/d2184/sing-geodata/raw/release/geoip.db](https://github.com/d2184/sing-geodata/raw/release/geoip.db)

- **geoip-lite.db**:
  - [https://github.com/d2184/sing-geodata/raw/release/geoip-lite.db](https://github.com/d2184/sing-geodata/raw/release/geoip-lite.db)

- **geosite.db**:
  - [https://github.com/d2184/sing-geodata/raw/release/geosite.db](https://github.com/d2184/sing-geodata/raw/release/geosite.db)

## **geoip.db 内容**

同 [d2184/geoip](https://github.com/d2184/geoip)
- 新增类别（方便有特殊需求的用户使用）：
  - `geoip:bilibili`
  - `geoip:cloudflare`
  - `geoip:cloudfront`
  - `geoip:facebook`
  - `geoip:fastly`
  - `geoip:google`
  - `geoip:netflix`
  - `geoip:telegram`
  - `geoip:twitter`

## **geoip-lite.db 内容**

仅包含cn、private地址数据

## **geosite.db 内容**

用法同 [Loyalsoldier/v2ray-rules-dat](https://github.com/Loyalsoldier/v2ray-rules-dat)
  - **加入大量中国大陆域名、Apple 域名和 Google 域名**：
  - [@felixonmars/dnsmasq-china-list/accelerated-domains.china.conf](https://github.com/felixonmars/dnsmasq-china-list/blob/master/accelerated-domains.china.conf) 加入到 `geosite:cn` 类别中
  - [@felixonmars/dnsmasq-china-list/apple.china.conf](https://github.com/felixonmars/dnsmasq-china-list/blob/master/apple.china.conf) 加入到 `geosite:geolocation-!cn` 类别中（如希望本文件中的 Apple 域名直连，请参考 [geosite 的 Routing 配置方式](https://github.com/Loyalsoldier/v2ray-rules-dat#geositedat-1)）
  - [@felixonmars/dnsmasq-china-list/google.china.conf](https://github.com/felixonmars/dnsmasq-china-list/blob/master/google.china.conf) 加入到 `geosite:geolocation-!cn` 类别中（如希望本文件中的 Google 域名直连，请参考 [geosite 的 Routing 配置方式](https://github.com/Loyalsoldier/v2ray-rules-dat#geositedat-1)）
- **加入 GFWList 域名**：
  - 基于 [@gfwlist/gfwlist](https://github.com/gfwlist/gfwlist) 数据，通过仓库 [@cokebar/gfwlist2dnsmasq](https://github.com/cokebar/gfwlist2dnsmasq) 生成
  - 加入到 `geosite:gfw` 类别中，供习惯于 PAC 模式并希望使用 [GFWList](https://github.com/gfwlist/gfwlist) 的用户使用
  - 同时加入到 `geosite:geolocation-!cn` 类别中
- **加入 EasyList 和 EasyListChina 广告域名**：通过 [@AdblockPlus/EasylistChina+Easylist.txt](https://easylist-downloads.adblockplus.org/easylistchina+easylist.txt) 获取并加入到 `geosite:category-ads-all` 类别中
- **加入 AdGuard DNS Filter 广告域名**：通过 [@AdGuard/DNS-filter](https://kb.adguard.com/en/general/adguard-ad-filters#dns-filter) 获取并加入到 `geosite:category-ads-all` 类别中
- **加入 Peter Lowe 广告和隐私跟踪域名**：通过 [@PeterLowe/adservers](https://pgl.yoyo.org/adservers) 获取并加入到 `geosite:category-ads-all` 类别中
- **加入 Dan Pollock 广告域名**：通过 [@DanPollock/hosts](https://someonewhocares.org/hosts) 获取并加入到 `geosite:category-ads-all` 类别中

## Credits
  - [Loyalsoldier/geoip](https://github.com/Loyalsoldier/geoip)
  - [Loyalsoldier/v2ray-rules-dat](https://github.com/Loyalsoldier/v2ray-rules-dat)
  - [SagerNet/sing-geoip](https://github.com/SagerNet/sing-geoip)
  - [SagerNet/sing-geosite](https://github.com/SagerNet/sing-geosite)

## LICENSE

GNU General Public License v3.0
