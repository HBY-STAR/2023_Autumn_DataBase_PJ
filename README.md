# Database-2023Autumn

***

## 介绍

线上购物比价系统：

1. GenerateData用于生成数据，使用java编写。
2. client为前端，使用vue编写。
3. src为后端，使用go编写。

## 使用说明

### 数据生成和导入
1. GenerateData：运行RunAndGenerate.java即可生成json格式的测试数据
2. 创建数据库表price_comparator（或使用price_comparator.sql导入）
3. 在src/models/init.go 中配置mysql
4. src/utils/json 运行 `go test` 生成数据库表并导入GenerateData中的数据

### 构建和运行
1. client：在client目录下，命令行执行npm run dev即可启动前端，部署在localhost:8070。
2. src：运行go build main.go启动后端，部署在127.0.0.1:8080。
