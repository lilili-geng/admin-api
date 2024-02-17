# vite-react-admin-template-api
prod

## 项目使用技术
gin jwt viper swagger zap go-file-rotatelogs casbin

```
jwt 身份认证中间件
viper 读取配置文件
swagger 接口文档
go-file-rotatelogs 是一个用于在 Go 语言中实现日志文件切割和轮转的库 
casbin 权限 目前权限是 三种 
      /admin/* :只有role为 superAdmin时可访问  
      /user/* : role为superAdmin/user时均可访问 
      /* : login register 接口
```

## 启动
go mod tidy
go run main.go


###  概述
```
目前是未完善版本 先用前端的分支进行dev 后面完善在挪到前端主分支
前端在分支 react-antd-admin-provider
https://github.com/lilili-geng/react-antd-admin/tree/react-antd-admin-provider
```

