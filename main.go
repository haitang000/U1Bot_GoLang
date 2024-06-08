package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	// 我实在太懒了 go run main.go 复制粘贴运行
	fmt.Printf("\n####################################")
	fmt.Printf("\n#  欢迎使用 U1Bot GoLang | 重制版  #")
	fmt.Printf("\n#                                  #")
	fmt.Printf("\n#      使用 “菜单” 获取帮助      #")
	fmt.Printf("\n####################################")
	// 排这玩意排了十几分钟，之后不想搞边框了
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
		// 加入原因：老是肌肉记忆
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
	// 还是觉得菜单有边框好看
	var back string
	fmt.Print("\n\n输入 “返回” 以返回主菜单")
	fmt.Scanln(&back)
	if back == "返回" {
		main()
	}
}

func fish(){
	fmt.Println(" ")
	rand.Seed(time.Now().UnixNano())
	// 这玩楞就是根据目前时间来整的随机种子
	var fishid = rand.Intn(19)
	// fishid:种类ID(从0开始，在加入/减少鱼的种类时这里也要修改)
	var fishlong = rand.Intn(600)
	// fishlong:鱼的长度(默认600)
	var fish string
	var danwei string = "条"
	// danwei:这里是鱼的单位，可以在“尚方宝剑“之类的物品中改，详情看fishid为0的鱼种类，内含所有变量的用途

		if fishid == 0 {
			fish = "鲤鱼"
			//fish:限定种类
			danwei = "条"
			//danwei:限定单位(默认为条)
			fishlong = rand.Intn(200)
			//fishlong:限定长度(默认600)
			//Tips：在fishlong 后输入“<值> +”可以设定为从<值>开始，但同时也要在括号内减去<值>
		}
		if fishid == 1 {
			fish = "鲫鱼"
			// 这里的单位就可写可不写了，随你便，只要你不嫌麻烦，你随便写
		}
		if fishid == 2 {
			fish = "尚方宝剑"
			danwei = "把"
			fishlong = rand.Intn(80) + 20
		}
		if fishid == 3 {
			fish = "心海"
			danwei = "只"
			fishlong = rand.Intn(100) + 50
			// 从原作那里搬过来的
		}
		if fishid == 4 {
			fish = "Code鱼"
		}
		if fishid == 5 {
			fish = "抑郁鱼"
		}
		if fishid == 6 {
			fish = "小金鱼"
			fishlong = rand.Intn(5) + 1
			// 小金鱼比小鱼小，不过分吧
		}
		if fishid == 7 {
			fish = "小杂鱼~♡"
			// 从原作那里搬过来的
		}
		if fishid == 8 {
			fish = "喝醉的鱼"
		}
		if fishid == 9 {
			fish = "小鱼"
			fishlong = rand.Intn(50)
		}
		if fishid == 10 {
			fish = "404鱼"
		}
		if fishid == 11 {
			fish = "贴图错误鱼"
		}
		if fishid == 12 {
			fish = "派蒙"
			danwei = "只"
			fishlong = rand.Intn(50) + 70
			// 加这个单纯是因为我玩原神玩的
		}
		if fishid == 13 {
			fish = "纯水精灵"
			danwei = "只"
			fishlong = rand.Intn(50) + 100
			// 加这个单纯是因为我玩原神玩的
		}
		if fishid == 14 {
			fish = "社恐鱼"
		}
		if fishid == 15 {
			fish = "金龙鱼"
			// 金龙鱼一比一比一~
		}
		if fishid == 16 {
			fish = "金鱼竿"
			danwei = "根"
			fishlong = rand.Intn(30) + 10
		}
		if fishid == 17 {
			fish = "臭袜子"
			danwei = "双"
			fishlong = rand.Intn(20) + 9
		}
		if fishid == 18 {
			fish = "AirPods"
			danwei = "个"
			fishlong = 3
			// 因为某人丢了一个AirPods，所以加上了
		}
		if fishid == 19 {
			fish = "无人机"
			danwei = "台"
			fishlong = rand.Intn(20) + 20
		}
		//判断鱼的种类，到时候慢慢加
	fmt.Print("\n正在钓鱼...")
	time.Sleep(2 * time.Second)
	// 本来是等1秒的，但是为了体验，改成了2秒
	fmt.Printf("你钓到了一%s%s，长度为%dcm!\n\n", danwei, fish, fishlong)
	var back string
	fmt.Print("\n\n输入 “返回” 以返回主菜单")
	fmt.Scanln(&back)
	// 返回的代码我以前一直是搞反的，以至于我一时半会我还没发现这个BUG
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
	fmt.Print("\n当前版本: V0.0.4 Alpha")
	fmt.Print("\n原作者: Mr.ling")
	fmt.Print("\n重制版作者: haitang000")
	fmt.Print("\n原作GitHub仓库: https://github.com/crashvibe/u1bot/")
	fmt.Print("\n原作官方Wiki: https://u1.tblstudio.cn/")
	fmt.Print("\n原作官方QQ体验群: 713478803")
}
