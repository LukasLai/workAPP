package inventoryparse

import (
	"fmt"
	"strings"
	"unicode"
	"workAPP/Calculate"
)

func ParseInvenroty(str string) ([]string, []string, []string) {
	//初始化各廠商叫貨表輸出文字
	vendor1 := []string{"您好,我們需要"}
	vendor2 := []string{"您好,我們需要"}
	vendor3 := []string{"您好,我們需要"}

	//如有不合規的逗號"，"則更換成和規逗號","
	str = strings.Replace(str, "，", ",", -1)
	fmt.Println("解析文字成功")
	//解析分類資料。數字=數量 文字=物品類別 ','=分隔符號
	input := strings.Split(str, ",")
	for _, items := range input {
		num := ""
		item := ""
		for _, v := range items {
			if unicode.IsNumber(v) || v == '.' {
				num += string(v)
			} else {
				item += string(v)
			}
		}

		UpdateInventory(item, num, &vendor1, &vendor2, &vendor3)
		num = ""
		item = ""

	}

	fmt.Println("檔案執行成功")
	fmt.Println()
	fmt.Println(vendor1)
	fmt.Println()
	fmt.Println(vendor2)
	fmt.Println()
	fmt.Println(vendor3)
	return vendor1, vendor2, vendor3
}

func UpdateInventory(name string, num string, vendor1, vendor2, vendor3 *[]string) { //經解析後的單筆資料分類導入指定的欄位
	switch name {
	case "沐浴乳":
		Calculate.Orderinventory(name, num, vendor1)
	case "洗髮精":
		Calculate.Orderinventory(name, num, vendor1)
	case "酒精":
		Calculate.Orderinventory(name, num, vendor1)
	case "擦手紙":
		Calculate.Orderinventory(name, num, vendor2)
	case "洗手乳":
		Calculate.Orderinventory(name, num, vendor2)
	case "大捲衛生紙":
		Calculate.Orderinventory(name, num, vendor2)
	case "小捲衛生紙":
		Calculate.Orderinventory(name, num, vendor2)
	case "大垃圾袋":
		Calculate.Orderinventory(name, num, vendor2)
	case "小垃圾袋":
		Calculate.Orderinventory(name, num, vendor2)
	case "套房咖啡包":
		Calculate.Orderinventory(name, num, vendor3)
	case "套房餅乾":
		Calculate.Orderinventory(name, num, vendor3)
	case "套房牙刷":
		Calculate.Orderinventory(name, num, vendor3)
	case "套房棉花棒":
		Calculate.Orderinventory(name, num, vendor3)
	case "化妝棉":
		Calculate.Orderinventory(name, num, vendor3)
	case "綠茶":
		Calculate.Orderinventory(name, num, vendor3)
	case "髮圈":
		Calculate.Orderinventory(name, num, vendor3)
	}
}
