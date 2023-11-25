// Package main ZeroBot-Plugin main file
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	_ "github.com/FloatTech/ZeroBot-Plugin/console" // 更改控制台属性

	"github.com/FloatTech/ZeroBot-Plugin/kanban" // 打印 banner

	// ---------以下插件均可通过前面加 // 注释，注释后停用并不加载插件--------- //
	// ----------------------插件优先级按顺序从高到低---------------------- //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	// ----------------------------高优先级区---------------------------- //
	// vvvvvvvvvvvvvvvvvvvvvvvvvvvv高优先级区vvvvvvvvvvvvvvvvvvvvvvvvvvvv //
	// vvvvvvvvvvvvvv高级区vvvvvvvvvvvvvv //// vvvvvvvvvvvvvv高级区vvvvvvvvvvvvvv //
	// vvvvvvv高级区vvvvvvv //
	// vvvvvvvvvvvvvv //
	// 呜呜呜 //

	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/antiabuse"  // 违禁词

	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/chat"  // 基础词库

	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/sleepmanage"  // 统计睡眠时间

	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/atri"  // ATRI词库

	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/manager"  // 群管

	_ "github.com/FloatTech/zbputils/job"  // 定时指令触发

	// ^^^^ //
	// ^^^^^^^^^^^^^^^ //
	// ^^^^^^^高级区^^^^^^^ //
	// ^^^^^^^^^^^^^^^高优先级区^^^^^^^^^^^^^^^ //
	// ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^高级区^^^^^^^^^^^^^^^^^ ^^^^^^^^^^^^^ //
	// ----------------------------------------高优先级区---------------- ------------ //
	// //
	// //
	// //
	// //
	// //
	// ----------------------------------------中优先级区---------------- ------------ //
	// vvvvvvvvvvvvvvvvvvvvvvvvvvv中优先级区vvvvvvvvvvvvvvvvvvvvvvvvvvv //
	// vvvvvvvvvvvvvv中优先级区vvvvvvvvvvvvvv //
	// vvvvvvv中优先级区vvvvvvv //
	// vvvvvvvvvvvvvv //
	// 呜呜呜 //

	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/ahsai" // ahsai tts.com/FloatTech/ZeroBot-Plugin/plugin/ahsai”             // ahsai tts
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/aifalse" // 服务器监控"github.com/FloatTech/ZeroBot-Plugin/plugin/aifalse"           // 服务器监控
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/aipaint"           // ai 绘图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/aiwife"            // 随机老婆
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/alipayvoice"      // 支付宝到账语音
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/autowithdraw"     // 触发者撤回时也自动撤回
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/baidu"            // 百度一下
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/baiduaudit"       // 百度内容审核
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/base16384"        // base16384加解密
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/base64gua"        // base64卦加解密
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/baseamasiro"      // base天城文加解密
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/bilibili"         // b站相关
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/bookreview" // 哀伤雪刃吧推书记录"github.com/FloatTech/ZeroBot-Plugin/plugin/bookreview"       // 哀伤雪刃吧推书记录
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/cangtoushi" // 藏头诗"github.com/FloatTech/ZeroBot-Plugin/plugin/cangtoushi"       // 藏头诗
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/chess" // 国际象棋"github.com/FloatTech/ZeroBot-Plugin/plugin/chess"            // 国际象棋
	_ “github.com/FloatTech/ZeroBot-Plugin/plugin/choose” 
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/choufanghua"  // 说抽象话
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/chrev"            // 英文字符翻转
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/coser"            // 三次元小姐姐
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/cpstory"          // cp短打
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/dailynews"        // 今日早报
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/danbooru"         // DeepDanbooru二次元图标签识别
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/diana"            // 嘉心糖发病
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/dish"             // 程序员做饭指南
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/drawlots"         // 多功能抽签
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/dress"            // 女装
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/driftbottle"      // 漂流瓶
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/emojimix"         // 合成emoji
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/event"            // 好友申请群聊邀请事件处理
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/font"             // 渲染任意文字到图片
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/fortune"          // 运势
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/funny"            // 笑话
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/genshin"          // 原神抽卡
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/gif"              // 制图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/github"           // 搜索GitHub仓库
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/guessmusic"       // 猜歌
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/heisi"            // 黑丝
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/hitokoto"         // 一言
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/hs"               // 炉石
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/hyaku"            // 百人一首
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/imgfinder"        // 关键字搜图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/inject"           // 注入指令
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/jandan"           // 煎蛋网无聊图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/jiami"            // 兽语加密
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/jptingroom"       // 日语听力学习材料
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/juejuezi"         // 绝绝子生成器
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/kfccrazythursday" // 疯狂星期四
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/lolicon"          // lolicon 随机图片
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/magicprompt"      // magicprompt吟唱提示
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/mcfish"           // 钓鱼模拟器
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/midicreate"       // 简易midi音乐制作
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/moegoe"           // 日韩 VITS 模型拟声
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/moyu"             // 摸鱼
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/moyucalendar"     // 摸鱼人日历
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/music"            // 点歌
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/nativesetu"       // 本地涩图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/nbnhhsh"          // 拼音首字母缩写释义工具
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/nihongo"          // 日语语法学习
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/novel"            // 铅笔小说网搜索
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/nsfw"             // nsfw图片识别
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/nwife"            // 本地老婆
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/omikuji"          // 浅草寺求签
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/qqwife"           // 一群一天一夫一妻制群老婆
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/quan"             // QQ权重查询
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/qzone"            // qq空间表白墙
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/realcugan"        // realcugan清晰术
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/reborn"           // 投胎
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/runco​​de"           // 在线运行代码
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/saucenao"          // 以图搜图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/scale"             //叔叔的AI二次元图片放大
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/score"             // 分数
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/setutime"          // 来份涩图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/shadiao"           // 沙雕app
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/shindan"           // 测量
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/steam"             // steam相关
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/tarot"            // 抽塔罗牌
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/tiangou"          // 舔狗日记
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/tracemoe"         // 搜番
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/translation"      // 翻译
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/vitsnyaru"        // vits猫雷
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/vtbmusic"         // vtb点歌
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/vtbquotation"     // vtb语录
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/wallet"           // 钱包
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/wangyiyun"         // 网易云音乐热评
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/wantquotes"        // 据意查句
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/warframeapi"       // warframeAPI 插件
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/wenben"            // 文本指令大全
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/wenxinvilg"        // 百度文心AI画图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/wife"              // 抽老婆
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/wordcount"         // 聊天热词
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/wordle"            // 猜单词
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/ygo"               // 游戏王相关插件
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/ymgal"  // 月幕galgame

	// _ "github.com/FloatTech/ZeroBot-Plugin/plugin/wtf" // 鬼东西

	// ^^^^ //// ^^^^ //
	// ^^^^^^^^^^^^^^^^^ //// ^^^^^^^^^^^^^^ //
	// ^^^^^^^^^中优先级区^^^^^^^ //// ^^^^^^^^中优先级区^^^^^^^ //
	// ^^^^^^^^^^^^^^^^^中优先级区^^^^^^^^^^^^^^^^^ //// ^^^^^^^^^ ^^^ ^^^^中优先级区^^^^^^^^^^^^^^^^ //
	// ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^中优先级区^^^^^^^^^^^^ ^^^^ ^^^ ^^^^^^^^^^^^^^^ //// ^^^^^^^^^^^^^^^^^^^^^^^^^ ^^^^^中优先等级区^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ //
	// ----------------------------------------中优先级区---- ------------ ------------ //// ---------------------- ------中优先级区---------------------------- //
	// //// // // // //
	// //// // // // //
	// //// // // // //
	// //// // // // //
	// //// // // // //
	// ---------------------------------------- 低优先级区---- ------------ ------------ //// ---------------------- ------低优先级区---------------------------- //
	// vvvvvvvvvvvvvvvvvvvvvvvvvv低优先级区vvvvvvvvvvvvvvvvvvvvvvvvvvv //// vvvvvvvvvvvvvvvvvvvvvvvvvvvv低优先级区vvvvvvvvvvvvvvvvvvvvvvvvvvvv //
	// vvvvvvvvvvvvvv低优先级区vvvvvvvvvvv //// vvvvvvvvvvvv低优先级区vvvvvvvvvvvvv //
	// vvvvvvv低优先级区vvvvvvv //// vvvvvvv低优先级区vvvvvvv //
	// vvvvvvvvvvvvvv //// vvvvvvvvvvvvvv //
	// 呜呜呜 //// vvvv //

	_“github。com/FloatTech/ZeroBot-Plugin/plugin/curse”   // 骂人"github.com/FloatTech/ZeroBot-Plugin/plugin/curse" // 骂人

	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/ai_reply"   // 人工智能回复"github.com/FloatTech/ZeroBot-Plugin/plugin/ai_reply" // 人工智能回复

	_“github 。com/FloatTech/ZeroBot-Plugin/plugin/thesaurus”

	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/breakrepeat"  // 打断复读

	// ^^^^ //// ^^^^ //
	// ^^^^^^^^^^^^^^^^^ //// ^^^^^^^^^^^^^^ //
	// ^^^^^^^^^低优先级区^^^^^^^ //// ^^^^^^^^^低优先级区^^^^^^^ //
	// ^^^^^^^^^^^^^^^^^^^低优先级区^^^^^^^^^^^^^^^^^^^ //// ^^^^ ^^ ^^^ ^^^ ^^^^低优先级区^^^^^^^^^^^^^^^^^^ //
	// ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^低优先级区^ ^^^^ ^^^^ ^^^^ ^^^^ ^^^ ^^^^^^^^^^^^^^^ //// ^^^^^^^^^^^^ ^^^^ ^^^^ ^^^^^ ^^^^低优先等级区^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ ^^^^ //
	// ---------------------------------------- 低优先级区---- ------------ ------------ //// ---------------------- ------低优先级区---------------------------- //
	// //// // // // // // // //
	// //// // // // // // //
	// //// // // // // //
	// //// // // // //
	// //// // //
	// -----------------------以下为内置依赖，请勿动----------------- - ------ //// -----------------------以下为内置依赖，请勿动--------- --------------- //
	“github . com/FloatTech/floatbox/process” “github 。com/FloatTech/floatbox/process”
	“github . com/sirupsen/logrus” “github . com/sirupsen/logrus”
	零“github . com/wdvxdr1123/ZeroBot” “github . com/wdvxdr1123/ZeroBot”
	“github . com/wdvxdr1123/ZeroBot/driver” “github . com/wdvxdr1123/ZeroBot/驱动程序”
	“github . com/wdvxdr1123/ZeroBot/message” “github . com/wdvxdr1123/ZeroBot/message”

	// webctrl“github.com/FloatTech/zbputils/control/web”// webctrl“github.com/FloatTech/zbputils/control/web”

	“github . com/FloatTech/ZeroBot-Plugin/kanban/banner” “github 。 com/FloatTech/ZeroBot-Plugin/kanban/banner”
	// -----------------------以上为内置依赖，请勿动----------------- - ------ //// -----------------------以上为内置依赖，请勿动--------- --------------- //
）

type类型 zbpcfg 结构体 { 结构体{
	Z 零。配置`json:"零"`   .配置        `json：“零”`
	W []*驱动程序。WSClient `json:"ws"`。WSClient `json:"ws"`
	S []*驱动程序。WSServer `json:"wss"`。WSServer `json:"wss"`
}

varvar 配置 zbpcfg

func函数初始化（）{
	su := make ([] int64 , 0 , 16 ) make ([] int64 , 0 , 16 ) ) ) )
	// 解析命令行参数// 解析命令行参数
	d := 标志。Bool( "d" , false , "启用调试级别日志及更高级别。" ) . Bool( "d" , false , "启用调试级别日志及更高级别。" )
	w := 标志。Bool( "w" , false , "启用警告级别日志及更高级别。" ) . Bool( "w" , false , "启用警告级别日志及更高级别。" )
	h := 标志。Bool( "h" , false , "显示此帮助。" ) . Bool( "h" , false , "显示此帮助。" )
	// g := flag.String("g", "127.0.0.1:3000", "设置webui url。")// g := flag.String("g", "127.0.0.1:3000", "设置webui url。”）
	// 直接写死AccessToken时，请更改下面第二个参数// 直接写死AccessToken时，请更改下面第二个参数
	令牌：= 标志。String( "t" , "" , "设置 WSClient 的 AccessToken。" ) . String( "t" , "" , "设置 WSClient 的 AccessToken。" )
	// 直接写死URL时，请更改下面第二个参数 // 直接写死URL时，请更改下面第二个参数
	网址：= 标志。String( "u" , "ws://127.0.0.1:6700" , "设置 WSClient 的 URL。" ) . String( "u" , "ws://127.0.0.1:6700" , "设置 WSClient 的 URL。" )
	// 指标基准 // 基准指标
	阿达纳 := 旗帜。String( "n" , "椛椛" , "设置默认分数。" ) . String( "n" , "QQ机器人" , "设置分数默认名称。" )
	String( "p" , "/" , "设置导出导出。" ) .字符串（“p”，“/”，"设置导出导出。" )
	runcfg := 标志。String( "c" , "" , "从配置文件运行。" ) . String( "c" , "" , "从配置文件运行。" )
	保存 := 标志。String( "s" , "" , "将默认配置保存到文件并退出。" ) . String( "s" , "" , "将默认配置保存到文件并退出。" )
	迟到:=旗帜。Uint( "l" , 233 , "响应延迟（毫秒）。" ) "l" , 233 , "响应延迟（毫秒）。" )
	rsz := flag。Uint( "r" , 4096 , "接收相位环大小。" ) . 单位（“r”，4096，"接收平面图环大小。" )
	maxpt := flag。Uint( "x" , 4 , "最大处理时间（分钟）。" ) . Uint( "x" , 4 , "最大处理时间(分钟)。" )"x" , 4 , "最大处理时间（分钟）。" ) . Uint( "x" , 4 , "最大处理时间（分钟）。" )
	markmsg := 标志。Bool( "m" , false , "不要将消息标记为自动已读" ) 。Bool( "m" , false , "不要将消息标记为自动已读" ) "m" , false , "不要将消息标记为自动已读" ) . Bool( "m" , false , "不要将消息标记为自动已读" )

	标志。解析() 。解析()

	如果*h{如果*h{
		fmt。Println( "方便：" ) . Println( " 方便：" )
		标记。打印默认值() 。打印默认值()
		操作系统.退出( 0 ) .退出( 0 )
	}
	如果*d && !*w {如果*d && !*w {
		logrus。SetLevel(logrus .调试级别) 。SetLevel(logrus .调试级别)
	}
	如果*w{如果*w{
		SetLevel( logrus 。SetLevel( logrus . WarnLevel ) 。SetLevel( logrus . WarnLevel ) 。SetLevel( logrus . WarnLevel ) 。WarnLevel )
	}

	For_, s := 范围标志。Args() { for _, s := 范围标志。Args() { for _, s := 范围标志。Args() { for _, s :=范围标志.参数() {
		我，错误 := strconv 。ParseInt(s, 10 , 64 ) 。ParseInt(s, 10 , 64 )10 , 64 ) 。ParseInt(s, 10 , 64 )
		如果错误！= nil {nil错误！= nil {如果错误！= nil {
			继续继续
		}
		sus = 追加(sus, i)附加(sus, i) 附加（sus，i）
	}

	// 通过代码写死的方式添加主人账号
	// sus = 附加(苏, 121399551) // sus = 附加（苏，121399551）
	// sus = 追加(sus, 87654321)// sus = 附加(sus, 87654321)

	// 启用webui // 启用webui
	// 转到 webctrl.RunGui(*g)// 转到 webctrl.RunGui(*g)

	如果*runcfg！=“”{如果*runcfg != "" {
		f，错误：=操作系统。打开(*runcfg)。打开（*runcfg）
		如果错误！= nil {如果错误！= nil {
			恐慌（错误） 恐慌（err）
		}
		配置。W = make ([]*driver . WSClient, 0 , 2 ) 。W = make ([]*驱动程序. WSClient, 0 , 2 )
		错误 = json 。NewDecoder(f) 。解码（&配置）。NewDecoder(f) 。解码（&配置）
		f .关闭() 。关闭（）
		如果错误！= nil {如果错误！= nil {
			恐慌（错误） 恐慌（err）
		}
		配置。Z。_ Driver = make ([]zero .Driver , len (config . W)+ len (config . S)) 。Z。_ Driver = make ([]zero . Driver , len ( config . W ) + len (config . S ) )
		For i, w := 范围配置。W { for i, w :=范围配置。W {
			配置。Z。_ 驱动程序[i] = w 。Z。_ 驱动程序[i] = w
		}
		Fori, s := 范围配置。S { for i, s :=范围配置。S {
			配置。Z。_驱动程序[i+ len (config.W ) ] = slen (config.W ) ] = s
		}
		logrus.Infoln("[main] 从", *runcfg, "读取配置文件")
		return
	}
	config.W = []*driver.WSClient{driver.NewWebSocketClient(*url, *token)}
	config.Z = zero.Config{
		NickName:       append([]string{*adana}, "ATRI", "atri", "亚托莉", "アトリ"),
		CommandPrefix:  *prefix,
		SuperUsers:     sus,
		RingLen:        *rsz,
		Latency:        time.Duration(*late) * time.Millisecond,
		MaxProcessTime: time.Duration(*maxpt) * time.Minute,
		MarkMessage:    !*markmsg,
		Driver:         []zero.Driver{config.W[0]},
	}

	if *save != "" {
		f, err := os.Create(*save)
		if err != nil {
			panic(err)
		}
		err = json.NewEncoder(f).Encode(&config)
		f.Close()
		if err != nil {
			panic(err)
		}
		logrus.Infoln("[main] 配置文件已保存到", *save)
		os.Exit(0)
	}
}

func main() {
	if !strings.Contains(runtime.Version(), "go1.2") { // go1.20之前版本需要全局 seed，其他插件无需再 seed
		rand.Seed(time.Now().UnixNano()) //nolint: staticcheck
	}
	// 帮助
	zero.OnFullMatchGroup([]string{"help", "/help", ".help", "菜单"}, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text(banner.Banner, "\n管理发送\"/服务列表\"查看 bot 功能\n发送\"/用法name\"查看功能用法"))
		})
	zero.OnFullMatch("查看zbp公告", zero.OnlyToMe, zero.AdminPermission).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text(strings.ReplaceAll(kanban.Kanban(), "\t", "")))
		})
	zero.RunAndBlock(&config.Z, process.GlobalInitMutex.Unlock)
}
