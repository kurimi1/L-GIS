package main

import (
	"LGIS/dao"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/flopp/go-findfont"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Province = map[string]string{
	"北京":  "11",
	"天津":  "12",
	"河北":  "13",
	"山西":  "14",
	"内蒙":  "15",
	"辽宁":  "21",
	"吉林":  "22",
	"黑龙江": "23",
	"上海":  "31",
	"江苏":  "32",
	"浙江":  "33",
	"安徽":  "34",
	"福建":  "35",
	"江西":  "36",
	"山东":  "37",
	"河南":  "41",
	"湖北":  "42",
	"湖南":  "43",
	"广东":  "44",
	"广西":  "45",
	"海南":  "46",
	"重庆":  "50",
	"四川":  "51",
	"贵州":  "52",
	"云南":  "53",
	"西藏":  "54",
	"陕西":  "61",
	"甘肃":  "62",
	"青海":  "63",
	"宁夏":  "64",
	"新疆":  "65",
}

var ProvinceZh = []string{
	"北京",
	"天津",
	"河北",
	"山西",
	"内蒙",
	"辽宁",
	"吉林",
	"黑龙江",
	"上海",
	"江苏",
	"浙江",
	"安徽",
	"福建",
	"江西",
	"山东",
	"河南",
	"湖北",
	"湖南",
	"广东",
	"广西",
	"海南",
	"重庆",
	"四川",
	"贵州",
	"云南",
	"西藏",
	"陕西",
	"甘肃",
	"青海",
	"宁夏",
	"新疆",
}

var Table = map[string]string{
	"矿产地基本情况":  "kcdj",
	"矿区地质情况":   "kqdz",
	"矿体特征":     "kttz",
	"煤矿体特征":    "mctz",
	"主要可采煤层特征": "kcmc",
	"勘查区资源量":   "kcql",
	"矿产资源储量":   "kccl",
	"选矿试验":     "xksy",
	"开采技术条件":   "kcjs",
	"矿产勘查工作概况": "kckc",
	"矿床技术经济评价": "kcjj",
}

var Mine = map[string]string{
	"煤":  "1001",
	"铁矿": "2001",
	"锰矿": "2002",
	"铜矿": "2006",
	"铅矿": "2007",
	"锌矿": "2008",
	"镍矿": "2012",
	"钨矿": "2014",
	"锡矿": "2015",
	"钼矿": "2017",
	"汞矿": "2018",
	"锇矿": "2105",
	"金矿": "2201",
}

// 方式一  设置环境变量   通过go-findfont 寻找simkai.ttf 字体
func init() {
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		//楷体:simkai.ttf
		//黑体:simhei.ttf
		if strings.Contains(path, "simkai.ttf") {
			os.Setenv("FYNE_FONT", path) // 设置环境变量  // 取消环境变量 os.Unsetenv("FYNE_FONT")
			break
		}
	}
}

func main() {
	// 链接数据库

	myApp := app.New()
	myWindow := myApp.NewWindow("LGIS矿产地查询系统")
	// 设置窗口大小
	myWindow.Resize(fyne.NewSize(1000, 800))
	sed := widget.NewLabel("")

	// // 选择数据
	var tableList []string
	for name := range Table {
		tableList = append(tableList, name)
	}
	table := widget.NewSelect(tableList, func(s string) {})
	table.PlaceHolder = "请选择数据"
	// 选择省份中文
	provincezh := widget.NewSelect(ProvinceZh, func(value string) {
		//sed.SetText(value)
	})
	provincezh.PlaceHolder = "请选择省份"

	// 查询按钮
	button := widget.NewButton("查询", func() {
		// 查询数据库
		ctx := context.Background()
		db, err := gorm.Open(mysql.Open("xx:xxx@tcp(127.0.0.1:3306)/databse?charset=utf8mb4&parseTime=True&loc=Local"))
		if err != nil {
			sed.SetText("数据库链接失败")
		}
		// 矿产地基本情况
		kcdj := dao.Use(db).Kcdj
		// 根据选择的省份获取编号
		province := fmt.Sprintf("%s[1-9]", provincezh.Selected)
		mine := Mine["金矿"]
		temp, err := dao.Use(db).Kcdj.WithContext(ctx).Where(kcdj.KCAAA.Regexp(province), kcdj.KCC.Eq(mine)).Find()
		fmt.Println(temp)
		if err != nil {
			sed.SetText("查询失败")
		} else {
			sed.SetText("查询成功")
		}
		result := [][]string{}
		for _, v := range temp {
			// struct 数据转list
			list := []string{}
			js, _ := json.Marshal(v)
			list = append(list, string(js))
			result = append(result, list)
		}
		// 查询结果开个窗口
		w2 := myApp.NewWindow("查询结果")
		list := widget.NewTable(
			func() (int, int) {
				return len(result), len(result[0])
			},
			func() fyne.CanvasObject {
				return widget.NewLabel("wide content")
			},
			func(i widget.TableCellID, o fyne.CanvasObject) {
				o.(*widget.Label).SetText(result[i.Row][i.Col])
			})
		w2.SetContent(list)
		w2.Resize(fyne.NewSize(800, 600))
		w2.Show()
	})
	// 上下两个布局，list单独一个布局需要绘画
	myWindow.SetContent(container.NewVBox(table, provincezh, button, sed))
	myWindow.Show()
	myApp.Run()
}
