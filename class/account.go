package class
import "fmt"

type account struct {
	key string		//接收用户输入
	balance float64	//账户余额
	money float64	//收支金额
	note string		//收支信息
	detail string	//收支记录
}

func NewAccount() *account {
	return &account {
		key:"",
		balance:10000.0,
		money:0.0,
		note:"",
		detail:"收    支\t账户金额\t收支金额\t说    明",
	}
}

func (a *account) ShowMenu() {
	loop:
	for {
		fmt.Println("\n----------家庭收支记账软件----------")
		fmt.Println("            1 收支明细")
		fmt.Println("            2 登记收入")
		fmt.Println("            3 登记支出")
		fmt.Println("            4 退出软件")
		fmt.Print("请选择(1 - 4):")

		fmt.Scanln(&a.key)
		switch a.key {
		case "1":a.showDetail()
		case "2":a.income()
		case "3":a.expend()
		case "4":
			if(a.exit()) {
				break loop
			}
		default:fmt.Println("请输入正确的选项..")
		}
	}
}

func (a *account) showDetail() {
	fmt.Println("----------当前收支明细记录----------")
	fmt.Println(a.detail)
}

func (a *account) income() {
	fmt.Print("本次收入金额:")
	fmt.Scanln(&a.money)
	a.balance += a.money
	fmt.Print("本次收入说明:")
	fmt.Scanln(&a.note)
	a.detail += fmt.Sprintf("\n收    入\t%v\t\t%v\t\t%v",a.balance,a.money,a.note)
}

func (a *account) expend() {
	fmt.Print("本次支出金额:")
	fmt.Scanln(&a.money)
	if a.money > a.balance {
		fmt.Println("余额不足...")
	} else {
		a.balance -= a.money
		fmt.Print("本次支出说明:")
		fmt.Scanln(&a.note)
		a.detail += fmt.Sprintf("\n支    出\t%v\t\t%v\t\t%v",a.balance,a.money,a.note)
	}
}

func (a *account) exit() bool {
	fmt.Print("你确定要退出么？y/n:")
	for {
		fmt.Scanln(&a.key)
		if a.key == "y" || a.key == "n" {
			break
		}
		fmt.Print("你的输入有误，请重新输入 y/n:")
	}
	if  a.key == "y" {
		return true
	}
	return false
}