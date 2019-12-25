# Go基础类库(私人订制)

## 主要依赖库列表

- github.com/araddon/dateparse
- github.com/axgle/mahonia
- github.com/clbanning/mxj
- github.com/disintegration/imaging
- github.com/fatih/color
- github.com/gin-gonic/gin
- github.com/go-redis/redis/v7
- github.com/go-resty/resty/v2
- github.com/hokaccha/go-prettyjson
- github.com/jinzhu/gorm
- github.com/rifflock/lfshook
- github.com/rs/xid
- github.com/satori/go.uuid
- github.com/sirupsen/logrus
- github.com/sony/sonyflake
- github.com/spf13/afero
- github.com/spf13/viper
- github.com/tidwall/gjson v1.3.5
- gopkg.in/confluentinc/confluent-kafka-go
- gopkg.in/gomail

## 目录结构

``` bash
├── components          //基础组件(需要注册才能使用)
│   ├── cache
│   ├── db
│   ├── logger
│   ├── mail
│   └── queue
├── middlewares         //中间件
│   └── gin
├── structs             //通用结构体定义
├── tests               //类库单元测试文件
├── thirdpartys         //第三方服务封装
│   ├── aliyun
│   └── dingding
└── tools               //工具函数
    ├── coding
    ├── datetime
    ├── debuger
    ├── file
    ├── http
    ├── image
    ├── random
    └── uuid
```
