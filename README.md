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
│   ├── img
|   └── db
├── README.md       # 项目文档
│
├── go.mod          # Go 模块文件
└── go.sum          # 依赖文件
```

编译命令

```
初始编译，使用os选择当前系统
fyne package -os windows -icon ./resources/img/joker.jpg
#交叉编译
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc fyne package -icon ./resources/img/joker.jpg
```
