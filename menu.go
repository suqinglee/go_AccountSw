package main
import "fmt"

func main() {

	//接受用户输入变量
	var key string

	//显示这个主菜单
	loop:
	for {
		fmt.Println("----------家庭收支记账软件----------")
		fmt.Println("            1 收支明细")
		fmt.Println("            2 登记收入")
		fmt.Println("            3 登记支出")
		fmt.Println("            4 退出软件")
		fmt.Print("请选择(1 - 4):")

		fmt.Scanln(&key)

		switch key {
		case "1":fmt.Println("----------当前收支明细记录----------")
		case "2":
		case "3":fmt.Println("登记支出..")
		case "4":break loop
		default:fmt.Println("请输入正确的选项..")
		}
	}

	fmt.Println("退出成功...")
}