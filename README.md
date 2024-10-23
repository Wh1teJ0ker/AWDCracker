# AWDCracker

```
/AWDCracker
│
├── main.go         # 主程序入口
│
├── /pkg            # 核心功能模块
│   ├── scan.go     # 扫描模块
│   ├── attack.go   # 防御模块
│   ├── utils.go    # 工具函数库
│   └── config.go   # 配置模块
│
├── /internal       # 内部工具
│   └── logger.go   # 日志工具
│
├── /resources        # 资源文件存放
│   └── start.sh
│
├── /ui                  # UI模块
│   ├── scan.go          # 扫描模块的 UI
│   ├── attack.go        # 攻击模块的 UI
│   ├── interference.go  # 干扰模块的 UI
│   ├── shell.go         # Shell模块的 UI
│   └── log.go           # 日志模块的 UI
|
├── README.md       # 项目文档
│
├── go.mod          # Go 模块文件
└── go.sum          # 依赖文件
```

编译命令

```
fyne package -os windows -icon joker.jpg
```
