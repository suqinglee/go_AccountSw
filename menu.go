package main
import "fmt"

func main() {

	//接受用户输入变量
	var key string

	//账户金额
	balance := 10000.0
	//每次收支金额
	money := 0.0
	//每次收支说明
	note := ""
	//收支详情，使用字符串记录
	detail := "收    支\t账户金额\t收支金额\t说    明"


	//显示这个主菜单
	loop:
	for {
		fmt.Println("\n----------家庭收支记账软件----------")
		fmt.Println("            1 收支明细")
		fmt.Println("            2 登记收入")
		fmt.Println("            3 登记支出")
		fmt.Println("            4 退出软件")
		fmt.Print("请选择(1 - 4):")

		fmt.Scanln(&key)

		switch key {
		case "1":
			fmt.Println("----------当前收支明细记录----------")
			fmt.Println(detail)
		case "2":
			fmt.Print("本次收入金额:")
			fmt.Scanln(&money)
			balance += money	//修改账户余额
			fmt.Print("本次收入说明:")
			fmt.Scanln(&note)
			//将这个收入情况，拼接到detail变量
			detail += fmt.Sprintf("\n收    入\t%v\t\t%v\t\t%v",balance,money,note)
		case "3":fmt.Println("登记支出..")
		case "4":break loop
		default:fmt.Println("请输入正确的选项..")
		}
	}

	fmt.Println("退出成功...")
}