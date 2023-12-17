# Database-2023Autumn

***

## 介绍

线上购物比价系统：

1. GenerateData用于生成数据，使用java编写。
2. client为前端，使用vue编写。
3. src为后端，使用go编写。

## 使用说明

### 环境配置
1. `cd ./client``npm install`下载vue依赖
2. `cd ./src` `go mod download`下载go依赖
3. `go install github.com/swaggo/swag/cmd/swag@latest`
`swag init --parseDependency --parseDepth 1`下载swagger自动生成api文档

### 数据生成和导入
1. GenerateData：运行RunAndGenerate.java即可生成json格式的测试数据
2. 创建数据库表price_comparator
3. 在src/models/init.go 中配置mysql
4. src/utils/json 运行 `go test` 生成数据库表并导入GenerateData中的数据

- 其中数据库表结构可以通过price_comparator.sql导入或由代码自动生成
- 触发器可由trigger.sql导入或使用代码中的hook函数模拟触发器行为
- 若选择使用trigger.sql导入触发器，需要注释掉hook函数即`func (item *CommodityItem) AfterUpdate(tx *gorm.DB) (err error)`

### 构建和运行
1. client：在client目录下，命令行执行`npm run dev`即可启动前端，部署在localhost:8070。
2. src：运行`go run main.go`启动后端，部署在127.0.0.1:8080。
