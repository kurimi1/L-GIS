# L-GIS

GUI版

矿产地查询系统

所有表没有主键，不符合标准

## 初始化

> go mod init LGIS

## 使用gorm-gen

<https://github.com/go-gorm/gen>

基于 GORM, 更安全更友好的ORM工具。

### 概览

- 自动生成 CRUD 和 DIY 方法
- 自动根据表结构生成模型（model）代码
- 事务、嵌套事务、保存点、回滚事务点
- 完全兼容 GORM
- 更安全、更友好
- 多种生成代码模式

## 使用gentool生成dao代码

> gentool -dsn "xx:xx@tcp(xx:xx)/xx" -tables "kcdj,kqdz,kttz,mctz,kcmc,kcql,kccl,xksy,kcjs,kckc,kcjj" -outPath "./dao"  
> go mod tidy

### 在model里修改中文字段名的struct

将数据字段改为Shuju

```GO
type Kcdj struct {
	KCAAA  string `gorm:"column:KCAAA" json:"KCAAA"`
	KCC    string `gorm:"column:KCC" json:"KCC"`
	JJDAJ  string `gorm:"column:JJDAJ" json:"JJDAJ"`
	JJGLA  string `gorm:"column:JJGLA" json:"JJGLA"`
	DWAAC  string `gorm:"column:DWAAC" json:"DWAAC"`
	DWAAD  string `gorm:"column:DWAAD" json:"DWAAD"`
	KCBA   string `gorm:"column:KCBA" json:"KCBA"`
	KCCA   string `gorm:"column:KCCA" json:"KCCA"`
	KCCB   string `gorm:"column:KCCB" json:"KCCB"`
	PKGKB  string `gorm:"column:PKGKB" json:"PKGKB"`
	KCAOC  string `gorm:"column:KCAOC" json:"KCAOC"`
	KCAOAE string `gorm:"column:KCAOAE" json:"KCAOAE"`
	KCAOAF string `gorm:"column:KCAOAF" json:"KCAOAF"`
	PKD    string `gorm:"column:PKD" json:"PKD"`
	JJDCBF string `gorm:"column:JJDCBF" json:"JJDCBF"`
	数据     string `gorm:"column:数据" json:"数据"`
}
```

## 链接数据库

```go
ctx := context.Background()
db, _ := gorm.Open(mysql.Open("xx:xxx@tcp(xxx:xxx)/xxx?charset=utf8mb4&parseTime=True&loc=Local"))
```

## GUI使用fyue

<https://github.com/fyne-io/fyne>  
Fyne是一个用 Go 编写的易于使用的 UI 工具包和应用程序 API。它旨在构建在具有单一代码库的桌面和移动设备上运行的应用程序。

## 打包

linux打包win，需要交叉编译

### 方法一：直接交叉编译

- 安装mingw

> apt install mingw-w64 -y

- 设置CC环境变量，设置为mingw的gcc

> export CC=x86_64-w64-mingw32-gcc

- 交叉编译

> fyne package -os windows -icon Icon.png

### 方法二：使用fyne-cross

需要安装docker,通过docker内的环境交叉编译

执行交叉编译
> fyne-cross windows
与方法一一致，只是在docker中安装mingw，但出现go mod 依赖安装问题因为防火墙的原因

### 解决不支持中文问题

设置环境变量   通过go-findfont 寻找simkai.ttf 字体
