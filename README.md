# Database-2023Autumn

***

## 介绍

线上购物比价系统：

1. GenerateData用于生成测试用的数据，使用java编写。
2. client为前端，使用vue编写。
3. src为后端，使用go编写。

## 使用说明

1. GenerateData：运行RunAndGenerate.java即可生成json格式的测试数据
2. client：在client目录下，命令行执行npm run dev即可启动前端，部署在localhost:8070。
3. src：在src/models/init.go中配置数据库。src/utils下运行go test json将测试数据导入数据库。src下运行go build main.go启动后端，部署在127.0.0.1:8080。
