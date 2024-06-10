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
	if input == "今天吃什么" {
		whateat()
	}
	if input == "疯狂星期四" {
		crazy_thursday()
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
	var fishid = rand.Intn(23)
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
		if fishid == 20 {
			fish = "烤激光鱼"
			fishlong = rand.Intn(4000) + 200
			// 尊重原作
		}
		if fishid == 21 {
			fish = "虚空珍珠"
			danwei = "颗"
			fishlong = rand.Intn(3000) + 20
			// 尊重原作
		}
		if fishid == 22 {
			fish = "金蛙"
			danwei = "只"
			fishlong = rand.Intn(700) + 30
		}
		if fishid == 23 {
			fish = "丁真鱼"
			fishlong = rand.Intn(1000) + 10
		}
		//判断鱼的种类，到时候慢慢加
	fmt.Print("\n正在钓鱼...")
	time.Sleep(2 * time.Second)
	// 本来是等1秒的，但是为了体验，改成了2秒
	fmt.Printf("你钓到了一%s%s，长度为%dcm!\n\n", danwei, fish, fishlong)
	var back string
	fmt.Print("\n\n输入 “返回” 以返回主菜单")
	fmt.Scanln(&back)
	// 返回的代码我以前一直是搞反的，以至于我一时半会我还没发现这个BUG qwq
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
	// :(
	fmt.Print("\n你今天的幸运值是", luck, ",", zhufuyu)
	var back string
	fmt.Scanln(&back)
	fmt.Print("\n输入 “返回” 以返回主菜单")
	if back == "返回" {
		main()
	}
}

func whateat(){
	fmt.Print("\n姚奕正在给你找好吃的...")
	time.Sleep(2 * time.Second)
	rand.Seed(time.Now().UnixNano())
	// 这玩楞也是根据目前时间来整的随机种子
	var something2eat = rand.Intn(29)
	var eat string = "冰清玉洁丸"
	// 夹带私货（
	if something2eat == 0 {
		eat = "小笼包"
	}
	if something2eat == 1 {
		eat = "馄饨"
	}
	if something2eat == 2 {
		eat = "馒头"
	}
	if something2eat == 3 {
		eat = "冰淇淋"
	}
	if something2eat == 4 {
		eat = "汉堡"
	}
	if something2eat == 5 {
		eat = "炸鸡"
	}
	if something2eat == 6 {
		eat = "肉夹馍"
	}
	if something2eat == 7 {
		eat = "蔬菜沙拉"
	}
	if something2eat == 8 {
		eat = "猪脚饭"
	}
	if something2eat == 9 {
		eat = "烤肉"
	}
	if something2eat == 10 {
		eat = "烤鱼"
	}
	if something2eat == 11 {
		eat = "米饭"
	}
	if something2eat == 12 {
		eat = "烤鸡"
	}
	if something2eat == 13 {
		eat = "火锅"
	}
	if something2eat == 14 {
		eat = "北京烤鸭"
	}
	if something2eat == 15 {
		eat = "烤串"
	}
	if something2eat == 16 {
		eat = "铁板烧"
	}
	if something2eat == 17 {
		eat = "蒸鱼"
	}
	if something2eat == 18 {
		eat = "红烧肉"
	}
	if something2eat == 19 {
		eat = "薯片"
	}
	if something2eat == 20 {
		eat = "关东煮"
	}
	if something2eat == 21 {
		eat = "三明治"
	}
	if something2eat == 22 {
		eat = "羊排"
	}
	if something2eat == 23 {
		eat = "重庆小面"
	}
	if something2eat == 24 {
		eat = "牛肉焗米粉"
	}
	if something2eat == 25 {
		eat = "鱼香肉丝"
	}
	if something2eat == 26 {
		eat = "葱油拌面"
	}
	if something2eat == 27 {
		eat = "麻婆豆腐"
	}
	if something2eat == 28 {
		eat = "火腿鹿茸菇炒荷兰豆"
	}
	if something2eat == 29 {
		eat = "烤肉拌饭"
	}
	fmt.Print("\n姚奕推荐你吃", eat)
	var back string
	fmt.Scanln(&back)
	fmt.Print("\n输入 “返回” 以返回主菜单")
	if back == "返回" {
		main()
	}
}

func crazy_thursday(){
	fmt.Print("\n姚奕正在给你找V50文案...")
	time.Sleep(2 * time.Second)
	rand.Seed(time.Now().UnixNano())
	// 这玩楞也是根据目前时间来整的随机种子
	var vivo50num = rand.Intn(4)
	var vivo50 string = "V50文案"
	if vivo50num == 0 {
		vivo50 = "实不相瞒我结婚七年了，他不爱我也不放过我，孩子从小被H集团幕后掌托人也就是他爷爷奶奶带走，所有人都阻拦我见孩子，一场车祸我失去了所有记忆，只记得我好想他，夜不能寐，茶饭不思，寝食难安，重复做着我带着他去吃肯德基疯狂星期四的梦，所以能请我吃肯德基疯狂星期四帮我恢复丢失的记忆吗……"
	}
	if vivo50num == 1 {
		vivo50 = "跟爸妈聊了很多，差不多年后就结婚，挺突然的吧，说真的，从来就没想过我这么快结婚，但现在我想收心了，遇到一个合适的人不容易，我希望我们举行婚礼的时候，所有好朋友都可以来参加，我知道我们都能好好照顾彼此，听到这个消息不知道你会不会伤心难过，会不会祝福，然而就算你不祝福我，也没事，明天肯德基疯狂星期四v我50请我吃"
	}
	if vivo50num == 2 {
		vivo50 = "今天和家人聊了很多，最终还是同意妥协了，答应跟上个月相亲男孩交往，顺利的话差不多年底就结婚，说真的 从来就没想过我这么快结婚，但现在我想收心了不想玩了，虽然还不能说很了解彼此，但我希望我们举行婚礼的时候，所有好朋友都可以来参加，听到这个消息不知道有些人会不会伤心难过、会不会祝福，然而就算你不祝福我也没事、毕竟我是复制过来玩玩的，我刚读别人的文案的时候我也信了。但是今天是肯德基疯狂星期四！谁请我吃？"
	}
	if vivo50num == 3 {
		vivo50 = "插播一条广告：\nApex双锤上大师\n清空赛季通行证\n免费帮上分\n身法教学\now80胜率上4300\n三位置意识教学\n彩虹六号生涯kd3.0\n包上冠军\nCSGO上全球精英\n这些都不接\n接肯德基疯狂星期四代吃"
	}
	if vivo50num == 4 {
		vivo50 = "扒个群里的渣女，我有个朋友被群友睡了，大概是三个月以前吧，我朋友在群里认识了一个女的，他们在网上聊了差不多两个星期左右就见面了，而且第一次见面就去宾馆了 ，之后的每个星期六星期天都会去，就这样持续了好几个月，我朋友给那个女的说，已经一个多月没奖励自己了，想不到这句话说完，隔天就找不到那个女的了，手机一直打不通，关机，直到我朋友前几天跟她父母坦白这件事情，我朋友就在旁边一直哭，那天晚上我偷偷给他买了个石原里美同款，还是没奖励自己了，然后我第二天请假一天,专门给他聊这个事，他说那个女的对他很好很好什么的，我说对你好还不是想玩你，实在不行抱树吧，他当时听完这句话就趴桌子上大哭，正当我打开手机想抱树的时候，不小心打开了肯德基，想起来今天原来是肯德基疯狂星期四，所以说谁请我肯德基？吃完我继续说"
	}
	fmt.Print("\n", vivo50)
	var back string
	fmt.Scanln(&back)
	fmt.Print("\n输入 “返回” 以返回主菜单")
	if back == "返回" {
		main()
	}
}
func about(){
	fmt.Print("\n=== U1Bot GoLang 重制版 ===")
	fmt.Print("\n当前版本: V0.1 Beta")
	fmt.Print("\n原作者: Mr.ling")
	fmt.Print("\n重制版作者: haitang000")
	fmt.Print("\n原作GitHub仓库: https://github.com/crashvibe/u1bot/")
	fmt.Print("\n原作官方Wiki: https://u1.tblstudio.cn/")
	fmt.Print("\n原作官方QQ体验群: 713478803")
	var back string
	fmt.Scanln(&back)
	fmt.Print("\n\n输入 “返回” 以返回主菜单")
	if back == "返回" {
		main()
	}
}
