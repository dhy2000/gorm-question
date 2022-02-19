# 一个 Gorm 的问题

在使用[gorm](https://gorm.io/)编写 ORM 模型时，如何描述及使用在 SQL 中利用外键表示实体间联系的数据模型？

## 运行方法

本例采用的数据库为 postgresql，如果测试环境为其他 gorm 支持的数据库，需将 `main.go` 第 28, 29 行对 `dsn` 和 `conn` 两个变量的赋值调整成相应的[数据库驱动](https://gorm.io/docs/connecting_to_the_database.html)（例如 MySQL）。

数据库的地址, 用户名密码等需填在一个名为 `config.yaml` 的配置文件中, 该文件格式与 `config.sample.yaml` 相同（可直接复制配置文件模板并填入相应字段）

安装依赖:

```shell
go mod download
```

构建和运行:
```shell
go build -o app
./app
```

（也可使用 GoLand 等 IDE 自带的依赖解决与构建功能）

## 代码内容

本项目基于一个简单的含有外键的数据模型, 采用了四种模型描述写法并分别进行了模型迁移, 数据插入, 数据查询测试。

模型内容与测试结果均位于 `main.go` 的注释中。模型描述及测试代码位于 `models` 的每个文件中。

## 主要的疑问

1. 对比 `archer` 和 `assassin` 两种模型, 当外键与被指向的主键名称相同时, 采用 has one 方法描述模型正确, belongs to 方法错误。
  这是否意味着在 gorm 中最好不要使用相同的外键和主键名?
2. 查询时，含有嵌套子结构的结构体的子结构均没有查询出结果，能否一次性查询出含嵌套子结构的查询结果？

