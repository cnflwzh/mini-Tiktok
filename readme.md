# mini-tiktok

## 项目结构

使用hertz框架，项目结构如下：

```bash
.
├── biz                             // 业务
│   ├── dal                         // 数据库操作
│   ├── handler                     // 业务逻辑
│   │   ├── chat
│   │   ├── comment
│   │   ├── favorite
│   │   ├── feed
│   │   ├── ping.go
│   │   ├── publish
│   │   ├── relation
│   │   └── user
│   ├── middleware                  // 中间件
│   ├── model                       // 数据模型 自动生成，不修改
│   │   ├── api
│   │   ├── common
│   │   ├── feed
│   │   ├── interact
│   │   │   ├── comment
│   │   │   └── favorite
│   │   ├── publish
│   │   ├── social
│   │   │   ├── chat
│   │   │   └── relation
│   │   └── user
│   └── router                      // 路由，自动生成，改中间件使用就可以
│       ├── chat
│       ├── comment
│       ├── favorite
│       ├── feed
│       ├── publish
│       ├── register.go
│       ├── relation
│       └── user
├── build.sh                        
├── go.mod
├── go.sum
├── idl                             // proto文件
│   ├── api.proto
│   ├── chat.proto
│   ├── comment.proto
│   ├── common.proto
│   ├── favorite.proto
│   ├── feed.proto
│   ├── google                      // google 的 proto include
│   ├── publish.proto
│   ├── relation.proto
│   └── user.proto
├── main.go                         // 项目入口
├── readme.md                       // 项目说明
├── router_gen.go                
├── router.go                       // api
└── script                         
    └── bootstrap.sh
```

使用 hertz 自带的 proto 生成工具生成初始项目框架。
