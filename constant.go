package main

const logo = `
████████╗██████╗  ██████╗██╗      █████╗ ███████╗██╗  ██╗
╚══██╔══╝██╔══██╗██╔════╝██║     ██╔══██╗██╔════╝██║  ██║
   ██║   ██████╔╝██║     ██║     ███████║███████╗███████║
   ██║   ██╔═══╝ ██║     ██║     ██╔══██║╚════██║██╔══██║
   ██║   ██║     ╚██████╗███████╗██║  ██║███████║██║  ██║
   ╚═╝   ╚═╝      ╚═════╝╚══════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝

   ● 原作者：mritd
   (GitHub: https://github.com/mritd, V2EX: https://www.v2ex.com/member/mritd)

   ● 现继任主要维护组织：TPClash Devs
   (GitHub: https://github.com/TPClash)

`

// https://github.com/torvalds/linux/blob/master/include/uapi/linux/capability.h
const (
	CAP_NET_BIND_SERVICE = 10
	CAP_NET_ADMIN        = 12
	CAP_NET_RAW          = 13
)

const (
	ChainDockerUser = "DOCKER-USER" // https://docs.docker.com/network/packet-filtering-firewalls/#docker-on-a-router
)

const (
	InternalClashBinName = "xclash"
	InternalConfigName   = "xclash.yaml"
)

const (
	bindAddressPatch = `# TPClash Common Config AutoFix
bind-address: '*'
`
	externalControllerPatch = `# TPClash Common Config AutoFix
external-controller: 0.0.0.0:9090
`
	secretPatch = `# TPClash Common Config AutoFix
secret: li0319
`
	tunStandardPatch = `# TPClash TUN AutoFix
tun:
  enable: true
  stack: system
  dns-hijack:
    - any:53
  auto-route: true
  auto-redir: true
`
	tunEBPFPatch = `# TPClash TUN eBPF AutoFix
tun:
  enable: true
  stack: system
  dns-hijack:
    - any:53
  auto-route: false
  auto-redir: false
`
	dnsPatch = `# TPClash DNS AutoFix
dns:
  enable: true
  listen: 0.0.0.0:1053
  enhanced-mode: fake-ip
  fake-ip-range: 198.18.0.1/16
  fake-ip-filter:
    - '*.lan'
    - '*.local'
    - '*.localdomain'
    - '*.example'
    - '*.invalid'
    - '*.localhost'
    - '*.test'
    - '*.local'
    - '*.home.arpa'
    - time.*.com
    - time.*.gov
    - time.*.edu.cn
    - time.*.apple.com
    - time1.*.com
    - time2.*.com
    - time3.*.com
    - time4.*.com
    - time5.*.com
    - time6.*.com
    - time7.*.com
    - ntp.*.com
    - ntp1.*.com
    - ntp2.*.com
    - ntp3.*.com
    - ntp4.*.com
    - ntp5.*.com
    - ntp6.*.com
    - ntp7.*.com
    - '*.time.edu.cn'
    - '*.ntp.org.cn'
    - +.pool.ntp.org
    - time1.cloud.tencent.com
    - music.163.com
    - '*.music.163.com'
    - '*.126.net'
    - musicapi.taihe.com
    - music.taihe.com
    - songsearch.kugou.com
    - trackercdn.kugou.com
    - '*.kuwo.cn'
    - api-jooxtt.sanook.com
    - api.joox.com
    - joox.com
    - y.qq.com
    - '*.y.qq.com'
    - streamoc.music.tc.qq.com
    - mobileoc.music.tc.qq.com
    - isure.stream.qqmusic.qq.com
    - dl.stream.qqmusic.qq.com
    - aqqmusic.tc.qq.com
    - amobile.music.tc.qq.com
    - '*.xiami.com'
    - '*.music.migu.cn'
    - music.migu.cn
    - +.msftconnecttest.com
    - +.msftncsi.com
    - msftconnecttest.com
    - msftncsi.com
    - localhost.ptlogin2.qq.com
    - localhost.sec.qq.com
    - +.srv.nintendo.net
    - +.stun.playstation.net
    - xbox.*.microsoft.com
    - xnotify.xboxlive.com
    - +.ipv6.microsoft.com
    - +.battlenet.com.cn
    - +.wotgame.cn
    - +.wggames.cn
    - +.wowsgame.cn
    - +.wargaming.net
    - proxy.golang.org
    - stun.*.*
    - stun.*.*.*
    - +.stun.*.*
    - +.stun.*.*.*
    - +.stun.*.*.*.*
    - heartbeat.belkin.com
    - '*.linksys.com'
    - '*.linksyssmartwifi.com'
    - '*.router.asus.com'
    - mesu.apple.com
    - swscan.apple.com
    - swquery.apple.com
    - swdownload.apple.com
    - swcdn.apple.com
    - swdist.apple.com
    - lens.l.google.com
    - stun.l.google.com
    - '*.square-enix.com'
    - '*.finalfantasyxiv.com'
    - '*.ffxiv.com'
    - '*.ff14.sdo.com'
    - ff.dorado.sdo.com
    - '*.mcdn.bilivideo.cn'
    - +.media.dssott.com
    - +.pvp.net
    - ydfkvps.site
    - ydfk.site
    - ydfknas.site
    - +.ydfk.site
    - +.ydfknas.site
    - +.ydfkvps.site
    - dns.msftncsi.com
    - www.msftncsi.com
    - www.msftconnecttest.com
    - +.tailscale.com
    - +.tailscale.io
    - controlplane.tailscale.com
    - log.tailscale.io
    - +.tztech.net
  default-nameserver:
    - 223.5.5.5
    - 119.29.29.29
  nameserver:
    - 223.5.5.5
    - 119.29.29.29
`
	nicPatch = `# TPClash Nic AutoFix
interface-name: {{MainNic}}
`
	ebpfPatch = `# TPClash eBPF AutoFix
ebpf:
  redirect-to-tun:
    - {{MainNic}}
`
	routingMarkPatch = `# TPClash routing-mark AutoFix
routing-mark: 666
`
)

const systemdTpl = `[Unit]
Description=Transparent proxy tool for Clash
After=network.target

[Service]
Type=simple
User=root
Restart=on-failure
ExecStart=/usr/local/bin/tpclash%s

RestartSec=10s
TimeoutStopSec=30s

[Install]
WantedBy=multi-user.target
`

const (
	installDir = "/usr/local/bin"
	systemdDir = "/etc/systemd/system"
)

const (
	lokiImage           = "grafana/loki:2.8.0"
	vectorImage         = "timberio/vector:0.X-alpine"
	trafficScraperImage = "vi0oss/websocat:0.10.0"
	tracingScraperImage = "vi0oss/websocat:0.10.0"
	grafanaImage        = "grafana/grafana-oss:latest"

	lokiContainerName           = "tpclash-loki"
	vectorContainerName         = "tpclash-vector"
	trafficScraperContainerName = "tpclash-traffic-scraper"
	tracingScraperContainerName = "tpclash-tracing-scraper"
	grafanaContainerName        = "tpclash-grafana"
)

const installedMessage = logo + `  👌 TPClash 安装完成, 您可以使用以下命令启动:
     ● 启动服务: systemctl start tpclash
     ● 停止服务: systemctl stop tpclash
     ● 重启服务: systemctl restart tpclash
     ● 开启自启动: systemctl enable tpclash
     ● 关闭自启动: systemctl disable tpclash
     ● 查看日志: journalctl -fu tpclash
     ● 重载服务配置: systemctl daemon-reload

     注：如果您使用的是非 systemd 的 Linux 发行版，请按照以下 systemd 的 service 配置作为参考自行编写。
     https://github.com/TPClash/tpclash/blob/master/constant.go#L91


     如有任何问题请开启 issue 或从 Telegram 讨论组反馈

     ● TPClash仓库: https://github.com/TPClash/tpclash
     ● TPClash Telegram 频道: https://t.me/tpclash
     ● TPClash Telegram 讨论组: https://t.me/+98SPc9rmV8w3Mzll
`

const reinstallMessage = `
  ❗监测到您可能执行了重新安装, 重新启动前请执行重载服务配置.
`

const uninstallMessage = `
  ❗️在卸载前请务必先停止 TPClash
  ❗️如果尚未停止请按 Ctrl+c 终止卸载
  ❗️本卸序将会在 30s 后继续执行卸载命令

`

const uninstalledMessage = logo + `  👌 TPClash 已卸载, 如有任何问题请开启 issue 或从 Telegram 讨论组反馈
     ● TPClash仓库: https://github.com/TPClash/tpclash
     ● TPClash Telegram 频道: https://t.me/tpclash
     ● TPClash Telegram 讨论组: https://t.me/+98SPc9rmV8w3Mzll
`

const (
	githubLatestApi   = "https://api.github.com/repos/TPClash/tpclash/releases/latest"
	githubUpgradeAddr = "https://github.com/TPClash/tpclash/releases/download/v%s/%s"
	ghProxyAddr       = "https://mirror.ghproxy.com/"
)

const upgradedMessage = logo + `  👌 TPClash 已升级完成, 请重新启动以应用更改
     ● 启动服务: systemctl start tpclash
     ● 停止服务: systemctl stop tpclash
     ● 重启服务: systemctl restart tpclash
     ● 开启自启动: systemctl enable tpclash
     ● 关闭自启动: systemctl disable tpclash
     ● 查看日志: journalctl -fu tpclash
     ● 重载服务配置: systemctl daemon-reload

     注：如果您使用的是非 systemd 的 Linux 发行版，请按照以下 systemd 的 service 配置作为参考自行编写。
     https://github.com/TPClash/tpclash/blob/master/constant.go#L91
`
