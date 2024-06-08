package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	//我实在太懒了 go run main.go 复制粘贴运行
	fmt.Printf("\n####################################")
	fmt.Printf("\n#  欢迎使用 U1Bot GoLang | 重制版  #")
	fmt.Printf("\n#                                  #")
	fmt.Printf("\n#      使用 “菜单” 获取帮助      #")
	fmt.Printf("\n####################################")
	var input string
	fmt.Printf("\n\n请输入您想要使用的功能: ")
	fmt.Scanln(&input)
	if input == "菜单" {
		menu()
	}
	if input == "运势" {
		yunshi()
	}
	if input == "钓鱼" {
		fish()
	}
	if input == "关于" {
		about()
	}
	if input == "go run main.go" {
		fmt.Print("笨蛋！这都能输成运行命令！")
		time.Sleep(time.Second * 5)
		main()
	}
}

func menu(){
	fmt.Print("\n###############################")
	fmt.Print("\n#  U1Bot GoLang 重制版  菜单  #")
	fmt.Print("\n#                             #")
	fmt.Print("\n#     1. 菜单| 打开本菜单     #")
	fmt.Print("\n#     2. 钓鱼| 钓鱼?启动!     #")
	fmt.Print("\n#    3. 运势| 查看今日运势    #")
	fmt.Print("\n#     4. 关于| 关于本项目     #")
	fmt.Print("\n#                             #")
	fmt.Print("\n###############################")
	var back string
	fmt.Scanln(&back)
	fmt.Print("\n\n输入 “返回” 以返回主菜单")
	if back == "返回" {
		main()
	}
}

func fish(){
	fmt.Println(" ")
	rand.Seed(time.Now().UnixNano())
	var fishid = rand.Intn(7)
	var fishlong = rand.Intn(600)
	var fish string
	var danwei string

		if fishid == 0 {
			fish = "鲤鱼"
			danwei = "条"
		}
		if fishid == 1 {
			fish = "鲫鱼"
			danwei = "条"
		}
		if fishid == 2 {
			fish = "尚方宝剑"
			danwei = "把"
		}
		if fishid == 3 {
			fish = "心海"
			danwei = "只"
		}
		if fishid == 4 {
			fish = "Code鱼"
			danwei = "条"
		}
		if fishid == 5 {
			fish = "抑郁鱼"
			danwei = "条"
		}
		if fishid == 6 {
			fish = "小金鱼"
			danwei = "条"
		}
		if fishid == 7 {
			fish = "小杂鱼~❥"
			danwei = "条"
		}
		//判断鱼的种类，到时候慢慢加
	fmt.Print("\n正在钓鱼...")
	time.Sleep(2 * time.Second)
	fmt.Printf("你钓到了一%s%s，长度为%dcm!\n\n", danwei, fish, fishlong)
	var back string
	fmt.Scanln(&back)
	fmt.Print("\n\n输入 “返回” 以返回主菜单")
	if back == "返回" {
		main()
	}
}

func yunshi(){
	var luck int = rand.Intn(100)
	rand.Seed(time.Now().UnixNano())
	var zhufuyu string
	if luck < 10 {
		zhufuyu = "是百分制的哦..."
	}
	if luck < 20 {
		zhufuyu = "呜......"
	}
	if luck < 40 {
		zhufuyu = "还好啦...还好啦..."
	}
	fmt.Print("\n你今天的幸运值是", luck, ",", zhufuyu)
	var back string
	fmt.Scanln(&back)
	fmt.Print("\n输入 “返回” 以返回主菜单")
	if back == "返回" {
		main()
	}
}

func about(){
	fmt.Print("\n=== U1Bot GoLang 重制版 ===")
	fmt.Print("\n当前版本: V0.0.3 Alpha")
	fmt.Print("\n原作者: Mr.ling")
	fmt.Print("\n重制版作者: haitang000")
	fmt.Print("\n原作GitHub仓库: https://github.com/crashvibe/u1bot/")
	fmt.Print("\n原作官方Wiki: https://u1.tblstudio.cn/")
	fmt.Print("\n原作官方QQ体验群: 713478803")
}