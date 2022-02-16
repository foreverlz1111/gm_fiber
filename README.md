# 0.目录
快速部署gorm_mysql_fiber项目
`(也支持SQL Server、SQlite)`
- [环境](https://github.com/foreverlz1111/gm_fiber#1)

- [运行](https://github.com/foreverlz1111/gm_fiber#2)

- [补充说明](https://github.com/foreverlz1111/gm_fiber#3)
<h1 name="1">1.环境</h1>

<h4>下载fiber：</h4>

```
go get github.com/gofiber/fiber/v2
```
<h4>下载gorm：</h4>

```
go get -u gorm.io/gorm
```
<h4>克隆项目：</h4>

```
git clone git@github.com:foreverlz1111/gm_fiber.git && cd gm_fiber
```

<h1 name="2">2.运行</h1>

<h4>检查错误：</h4>

```
go test app.go
```

<h4>正常输出：</h4>

`?       command-line-arguments  [no test files]`

<h4>运行项目：</h4>
  
```
go run app.go
```

- 浏览器输入: 

```
localhost:3000/hello
```

<h1 name="3">3.补充说明</h1>

- 默认尝试连接数据库并查询student表
