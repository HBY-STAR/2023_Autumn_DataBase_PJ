package pj.dbs.utils;

import org.apache.commons.lang3.RandomStringUtils;

import java.util.Arrays;
import java.util.HashSet;
import java.util.List;
import java.util.Random;

public class GenerateChinese {
    //常用名
    public static final String common = "云深婉雅明艳浩晨涵雨婷洋悠然韵瑶思淼婕舒静娅蕾芝伟伦琳翰涛晨璐婧梓萌珂晗瑾菲茜晨悦佳琪鑫嘉寅雅萱浩然志强梦楠欣然雅芳乐怡颖慧晓婷洁灵媛婧佳蓉妍琳梦婷明阳博文春芳盈宸欢佳欣韵涵欣悦丽莎忆柳梦佳欣静艺瑶梦露婷涵心静娴美丽思婉悦萌彤依楠润婷梦蕾露丹翠文婷静萍玲英婕雅紫菲晓芸明慧乐然晨涵海文欢雅馨雅楠馨儿娅佳璇雅雯奕璇芸楠菲瑶琴瑾涵琳晓春欢欢乐涵阳凯文伟毅世杰海涛佳辉文轩洋洋昊天泽宇鑫鑫锦程梓轩博涵晓明哲翔婧琪涵涵弘涛宇航越泽凯安梓晨昊然瑞祥梦洁梦璐梦婷若涵宇霆敬亭梦瑶梦佳依依涵韵洋洋悠然婉婷鑫婷博文悠悠文然琪涵天佳楠鑫婷婉佳欣琪洋涵明博欣然婕若涵玥晨露露琦若涵若涵洋欣悦涵欢" +
            "的一是在了不和有大这主中人上为们地个用工时要动国产以我到他会作来分生对于学下级就年阶义发成部民可出能方进同行面说种过命度革而多子后自社加小机也经力线本电高量长党得实家定深法表着水理化争现所二起政三好十战无农使性前等反体合斗路图把结第里正新开论之物从当两些还天资事队批如应形想制心样干都向变关点育重其思与间内去因件日利相由压员气业代全组数果期导平各基或月毛然问比展那它最及外没看治提五解系林者米群头意只明四道马认次文通但条较克又公孔领军流入接席位情运器并飞原油放立题质指建区验活众很教决特此常石强极土少已根共直团统式转别造切九你取西持总料连任志观调七么山程百报更见必真保热委手改管处己将修支识病象几先老光专什六型具示复安带每东增则完风回南广劳轮科北打积车计给节做务被整联步类集号列温装即毫知轴研单色坚据速防史拉世设达尔场织历花受求传口断况采精金界品判参层止边清至万确究书术状厂须离再目海交权且儿青才证低越际八试规斯近注办布门铁需走议县兵固除般引齿千胜细影济白格效置推空配刀叶率述今选养德话查差半敌始片施响收华觉备名红续均药标记难存测士身紧液派准斤角降维板许破述技消底床田势端感往神便贺村构照容非搞亚磨族火段算适讲按值美态黄易彪服早班麦削信排台声该击素张密害侯草何树肥继右属市严径螺检左页抗苏显苦英快称坏移约巴材省黑武培著河帝仅针怎植京助升王眼她抓含苗副杂普谈围食射源例致酸旧却充足短划剂宣环落首尺波承粉践府鱼随考刻靠够满夫失包住促枝局菌杆周护岩师举曲春元超负砂封换太模贫减阳扬江析亩木言球朝医校古呢稻宋听唯输滑站另卫字鼓刚写刘微略范供阿块某功套友限项余倒卷创律雨让骨远帮初皮播优占死毒圈伟季训控激找叫云互跟裂粮粒母练塞钢顶策双留误础吸阻故寸盾晚丝女散焊功株亲院冷彻弹错散商视艺灭版烈零室轻血倍缺厘泵察绝富城冲喷壤简否柱李望盘磁雄似困巩益洲脱投送奴侧润盖挥距触星松送获兴独官混纪依未突架宽冬章湿偏纹吃执阀矿寨责熟稳夺硬价努翻奇甲预职评读背协损棉侵灰虽矛厚罗泥辟告卵箱掌氧恩爱停曾溶营终纲孟钱待尽俄缩沙退陈讨奋械载胞幼哪剥迫旋征槽倒握担仍呀鲜吧卡粗介钻逐弱脚怕盐末阴丰编印蜂急拿扩伤飞露核缘游振操央伍域甚迅辉异序免纸夜乡久隶缸夹念兰映沟乙吗儒杀汽磷艰晶插埃燃欢铁补咱芽永瓦倾阵碳演威附牙芽永瓦斜灌欧献顺猪洋腐请透司危括脉宜笑若尾束壮暴企菜穗楚汉愈绿拖牛份染既秋遍锻玉夏疗尖殖井费州访吹荣铜沿替滚客召旱悟刺脑" +
            "措贯藏敢令隙炉壳硫煤迎铸粘探临薄旬善福纵择礼愿伏残雷延烟句纯渐耕跑泽慢栽鲁赤繁境潮横掉锥希池败船假亮谓托伙哲怀割摆贡呈劲财仪沉炼麻罪祖息车穿货销齐鼠抽画饲龙库守筑房歌寒喜哥洗蚀废纳腹乎录镜妇恶脂庄擦险赞钟摇典柄辩竹谷卖乱虚桥奥伯赶垂途额壁网截野遗静谋弄挂课镇妄盛耐援扎虑键归符庆聚绕摩忙舞遇索顾胶羊湖钉仁音迹碎伸灯避泛亡答勇频皇柳哈揭甘诺概宪浓岛袭谁洪谢炮浇斑讯懂灵蛋闭孩释乳巨徒私银伊景坦累匀霉杜乐勒隔弯绩招绍胡呼痛峰零柴簧午跳居尚丁秦稍追梁折耗碱殊岗挖氏刃剧堆赫荷胸衡勤膜篇登驻案刊秧缓凸役剪川雪链渔啦脸户洛孢勃盟买杨宗焦赛旗滤硅炭股坐蒸凝竟陷枪黎救冒暗洞犯筒您宋弧爆谬涂味津臂障褐陆啊健尊豆拔莫抵桑坡缝警挑污冰柬嘴啥饭塑寄赵喊垫康遵牧遭幅园腔订香肉弟屋敏恢忘衣孙龄岭骗休借丹渡耳刨虎笔稀昆浪萨茶滴浅拥穴覆伦娘吨浸袖珠雌妈紫戏塔锤震岁貌洁剖牢锋疑霸闪埔猛诉刷狠忽灾闹乔唐漏闻沈熔氯荒茎男凡抢像浆旁玻亦忠唱蒙予纷捕锁尤乘乌智淡允叛畜俘摸锈扫毕璃宝芯爷鉴秘净蒋钙肩腾枯抛轨堂拌爸循诱祝励肯酒绳穷塘燥泡袋朗喂铝软渠颗惯贸粪综墙趋彼届墨碍启逆卸航雾冠丙街莱贝辐肠付吉渗瑞惊顿挤秒悬姆烂森糖圣凹陶词迟蚕亿矩" +
            "脊歼羽掩汗碰谱童庭蓬贴岸店怪馆挡肢胆君乏傅凌恰吴鸡盆氮铃荡汇狂偶辽宴珊描监涉伏弃仔坯症睛窝跃串瑚饱巢辑迷诗肃谊胎宾顽钠辛阔牲估禁屑秀催炸搬坑暂埋墓腰隆堡迈慌钾魏踏旺蜜兼扭肺兄撒矮拆叉贮抬痕彩冻丛漆详拨瓜奔腿暖脾棒湾旅潜摄朱纤览融拍愚添抱蓄稿翅蛾锐栓签牌瞧疏舍糊驱泉毁伪锯卢函掘扰淬册棱爬豪螟标授朋俗骂仓脏昌邦欺博伐衰寻杠蜗尿幕絮蘖辨孵垄粹填丘歪鬼挺帅斥摘父狗罢炎疆肝酶恨曼蹲币返颠剩港颜酵梯楼绪淮邻御杰恒弗溉淀苯跨肿抑诸凉胚舒胀氢搭醒逃曰竞疾韩尘寿孤督涡甜拒梅乔锡睡昂烯拧扑郊患购蝗锅蔑赖瓶租怒巧膏涌狭醇惕档燕泰胁盘竭违丽氨框舌膨骤蓝幸诚吓秩扶芬咬牵忍椎愤迁仇滩仿绘辈拚喝驳畦番扑葡款敲邀郭妥隐s轰籽忆旦犹庸崇庙秆闸厉臣窗纺掠涝涨递葬阅堵扁钳棚鳞伴珍敦椭沃欲鼻宇甫锌皆铲砖贼渣济筛斋梦贪哇萄铺桃蟹挝糙颈雅晒韦耻沸雇储畏霍菲徐榜囊腺茨陕抹屈宿硝昨蔬郝铬茧窄哨辆耀仲薯僚浙饰朴恐腊兽蜡惠犁嘛售鳍敬坝烘颂叔卧纠络玩栏剑苹闯丢柏牺奎嚷宫肾笼郑叙奶芒霞朽妹茬码掀阁卑铰铵弦肤拟署淋梨迪俩撑呵申穆杯姑劝崩劣贺棕裁吐嫩凭曹摧疫鸟镍眉梁禾臭冈陵歧幻丧迭脆怨董镀酷罐逻橡浩撤驶享锦俺佛兔姿铅堤址溃胺皱晨胃氟灿漫泄枢" +
            "戴孕扣沼逼肌碗巡吊盗蚜钩汤梢挨翼疯鞭扇冶烦悉蔓泼桌柯罩啮勾舰晋扳谴侯倡诊鸣桂奖贾朵霜萌滞蛹阐偿译稼捞棵戈诬撮洒萧奸饮涅衬镗纱瘤葛饼凶饵沾馏钼鞋姓汞枣溜疼凑醛颌肖篡邓撞搅铡卜歇妨挽审凯轧垒箭炕浑龟账趣俭泪泊乃捉窑驾汁凿饿帽湘郎欣慎芳肪蔽绵畅盲缚焕惜仰衍廷玄泻蒲捣妙帕蛇锰棘溪匪绒潘疲纬鸭坎盒拼荫兆熊悲捧锄奉陪玛微钨籍蚊漂糟嘉狼桶拾唉默皂吕馈酯邪孝睛屠畸峡祥蒂拜蝉艾叹淑烤骄篮伞尝吏吞雹勘萎闲佳耙剿鳃砍冯毅骑酚咳煮披佩杏偷摊肚昔韧唇喘吵荆刑拦镁蹄瓷澳塌饥垮滋钝醋捍诡哩宏瞬缔婆扛捷刹猿葱亏阮帆纂喀邵丑郁茂糠俞泳夸砚抖渴聪拱泌藻靶褶扯藤悄逊岘姜砾舆瘦咸焰榴涛垦媳圃胳肆仑叠攀莲债汪棍飘闷蛄蔗贷俊傲哺蝼颁蠢鲤噪膀氛洼栅凤溢炊浦橄陡胰仙柔咂呆姐哭懈兹赋岳楔蜕嵌僵晰挠熙婚缠鬃佣吾辞抚暑遮嚣赴钒嫁磺膛辣谨鄙桩惨杉秸蝇鞘匆娟晃涕萍钛眇趁邮蛮廉熏侦浴俯圭颇赢掏帜枚酮瓣宙谣踩奏竖鞍曝耶茄谐躺榄臼哎抄铆晓虱矢艇坞鞅履恳弥搜肛逸喉苔茁欠叭扔琴芦俱砌拢礁茫筹辱靳枕惩醉挣婶拣嫂荚膊铂昏滨誓夕扮昼艘遥戒逢苍匈慈愁唤蕾帐掺丈瘟顷裕誉祸坛彭橘匹傍淤烷绞豫庞咒芝荀弓罚捏嗨楞仕嘻沫崖瞅帘榨墒捐恕螨汛诞赏琼贩鸿铭嘱隘驰娃瞎遣跌挪耘悦钴魂裸薛鲢躲鳙悠碘沥嘿灶饶酬艳堪淹怠砷吁涤慰缴窜羔趟脖锭兜魔梗炒纽奈硼鼎惑栗谎袁滥亨浊埂垅匝轲遂乒踪俘怔陨噬惧颖茅摔粳垃圾疮厅鄂讥隧睁痰镶哀劈峻尸拐拳眠蔡腋哑契翁肋砧捆哟菊笨垛谦畴膝铍猜殷咽巾赌骚挫钦乓痹嘲渍杭蕉妻壶仆耸蛙廊蛛翠鹰喻扼蕴寇腥瞪籼咕猫况鹏钮搏溅胫萝臀鲫羞罕殿忌亭盅菇旨吻厌宰氰屏桐颚佐栖蒜卓殉搁煌橙窖眨墩躁沛翘蜘酰矫钵哗梭毂嗓禽壕凳筐耦漠屁恭钡驴姚怖滔煽虾哼匠禄稚蚁窃咐茵坊裤勿熬狱熄荐镰柑屯醚耿髓戊腕愉蕨眶煎盈慨晕盼勉虏釉皿瘀昭蝈嗽讽秃谚畔疽冕宵窍峪槐癌敷岂侮携脓卿丸柜碾咀烃怜蝽傻椰逮猎崎淆寺恼胖颊氩盯赠甩坪淘谭莎雏棺躯熹蚧懒踢爵衷仟陋撕缆晾狡庇蓖酿拓簇蚌阎雀鹿卤荸荠搂琢猾苷祛崭硕苞逞炫厄焚铀舵耽爽稠跗邱盏廖韵豹钓奠溴枫犬猖驯侨灼翟擂嘀汹磅嚼狮爹鹅贤颅煞萤烙蛀裹骡痢巷寡碧猴栋嗯柿篷吱厩鳄蕊甸澄闺荧黔嫌瑟玲撇敞葫硷乞蛭阜矾瞒聊琅傀儡啃澜绥豌删龚衔敛厢堕潭舶翔赔夷稗啉僻堰恋萘扒瞄韶笋蚴媒榆廊衅吼锹睦颤刽啬慧碑彝瘠祭侣赚蝶郡叨岔坟疤螬悔譬乖巍疡禹魁掷棋憎阱坠碲卟哄彬绑腑押揉枷菱蹈汰渎愧珩贬衫宅蛴夯吭烫灸竿酱倦镦寮戳睾拴孜迄秤笛羟蜱樟鲍蠕芍诫慕虹厦弊翰锣沪逝诈劫锂咧凋毡蓟椅毯斧绸矣";
    //常用姓氏
    public static final String surname = "赵钱孙李周吴郑王冯陈褚卫蒋沈韩杨朱秦尤许何吕施张孔曹严华金魏陶姜戚谢邹喻柏水窦章云苏潘葛奚范彭郎鲁韦昌马苗凤花方俞任袁柳酆鲍史唐费廉岑薛雷贺倪汤滕殷罗毕郝邬安常乐于时傅皮卞齐康伍余元卜顾孟平黄和穆萧尹姚邵湛汪祁毛禹狄米贝明臧计伏成戴谈宋茅庞熊纪舒屈项祝董梁杜阮蓝闵席季麻强贾路娄危江童颜郭梅盛林刁钟徐邱骆高夏蔡田樊胡凌霍虞万支柯昝管卢莫经房裘缪干解应宗丁宣贲邓郁单杭洪包诸左石崔吉钮龚程嵇邢滑裴陆荣翁荀羊於惠甄麴家封芮羿储靳汲邴糜松井段富巫乌焦巴弓牧隗山谷车侯宓蓬全郗班仰秋仲伊宫宁仇栾暴甘钭厉戎祖武符刘景詹束龙叶幸司韶郜黎蓟薄印宿白怀蒲邰从鄂索咸籍赖卓蔺屠蒙池乔阴欎胥能苍双闻莘党翟谭贡劳逄姬申扶堵冉宰郦雍舄璩桑桂濮牛寿通边扈燕冀郏浦尚农温别庄晏柴瞿阎充慕连茹习宦艾鱼容向古易慎戈廖庾终暨居衡步都耿满弘匡国文寇广禄阙东殴殳沃利蔚越夔隆师巩厍聂晁勾敖融冷訾辛阚那简饶空曾毋沙乜养鞠须丰巢关蒯相查後荆红游竺权逯盖益桓公万俟司马上官欧阳夏侯诸葛闻人东方赫连皇甫尉迟公羊澹台公冶宗政濮阳淳于单于太叔申屠公孙仲孙轩辕令狐钟离宇文长孙慕容鲜于闾丘司徒司空亓官司寇仉督子车颛孙端木巫马公西漆雕乐正壤驷公良拓跋夹谷宰父谷梁晋楚闫法汝鄢涂钦段干百里东郭南门呼延归海羊舌微生岳帅缑亢况后有琴梁丘左丘东门西门商牟佘佴伯赏南宫墨哈谯笪年爱阳佟第五言福";
    //常见物品
    public static final List<String> chineseItemList = Arrays.asList(
            "手机", "电脑", "书籍", "衣服", "鞋子", "眼镜", "手表", "自行车", "咖啡机", "真空吸尘器",
            "双肩包", "钱包", "电视", "空调", "冰箱", "洗衣机", "热水器", "微波炉", "电磁炉", "沙发",
            "餐桌", "椅子", "锅", "锅碗瓢盆", "刀具", "筷子", "碗", "杯子", "马克杯", "水壶",
            "餐具", "盘子", "高跟鞋", "运动鞋", "拖鞋", "围巾", "帽子", "手套", "手帕", "披肩",
            "床", "床垫", "被子", "枕头", "床单", "毛巾", "牙刷", "牙膏", "洗发水", "沐浴露",
            "毛巾", "电吹风", "电动牙刷", "洗衣液", "护手霜", "香水", "沐浴球", "沐浴巾", "化妆品",
            "口红", "眼影", "化妆刷", "香水", "护肤品", "洗面奶", "面膜", "防晒霜", "香皂",
            "洗手液", "沐浴露", "洗发水", "电动牙刷", "空气净化器", "加湿器", "扫地机器人", "空气净化器",
            "音响", "耳机", "音响", "收音机", "相机", "摄像机", "笔记本", "咖啡杯", "打印机",
            "照相机", "游戏机", "游戏手柄", "跑步机", "瑜伽垫", "运动手环", "哑铃", "游泳镜",
            "网球拍", "自行车", "登山杖", "羽毛球拍", "滑板", "血压计", "体重秤"
    );

    public static final List<String> country = Arrays.asList(
            "中国", "美国", "英国", "法国", "俄罗斯", "日本", "印度", "巴西", "加拿大", "德国"
    );

    //常用形容词
    public static final List<String> chineseAdjectives = Arrays.asList(
            "美丽的", "聪明的", "善良的", "快乐的", "幸福的", "勤劳的", "懒惰的", "可爱的", "漂亮的", "丑陋的",
            "高大的", "矮小的", "强壮的", "瘦弱的", "胖乎乎的", "机智的", "坚强的", "柔软的", "坚定的", "温暖的",
            "冷酷的", "温和的", "沉默的", "吵闹的", "开心的", "悲伤的", "轻松的", "紧张的", "丰富的", "简单的",
            "复杂的", "神秘的", "平静的", "兴奋的", "无聊的", "有趣的", "特别的", "普通的", "奇怪的", "正常的",
            "精彩的", "平凡的", "忧郁的", "开朗的", "乐观的", "悲观的", "友好的", "冷漠的", "活泼的", "稳重的",
            "谨慎的", "浪漫的", "现实的", "虚拟的", "激动的", "紧张的", "安静的", "喧闹的", "高兴的", "害怕的",
            "胆小的", "勇敢的", "谦虚的", "骄傲的", "开放的", "封闭的", "诚实的", "诚信的", "花哨的", "简约的",
            "强大的", "微弱的", "自信的", "怀疑的", "强烈的", "微弱的", "顽固的", "灵活的", "敏感的", "木讷的",
            "热情的", "冷淡的", "积极的", "消极的", "稳定的", "不稳定的", "顺利的", "艰难的", "温馨的", "冷漠的",
            "亲密的", "疏远的", "缓慢的", "快速的", "适应的", "不适应的", "健康的", "不健康的", "强壮的", "虚弱的"
    );

    public static HashSet<String> nameSet = new HashSet<>();

    static public String generate_human_name() {
        String name_1 = RandomStringUtils.random(1, surname);

        int name_2_len = GenerateNum.generate_num(1, 2);
        String name_2 = RandomStringUtils.random(name_2_len, common);
        String name = name_1 + name_2;
        if (nameSet.contains(name)) {
            return generate_human_name();
        }
        nameSet.add(name);
        return name;
    }

    static public String generate_country() {
        Random random = new Random();
        int randomIndex = random.nextInt(country.size());
        return country.get(randomIndex);
    }

    static public String generate_thing() {
        Random random = new Random();
        int randomIndex = random.nextInt(chineseItemList.size());
        return chineseItemList.get(randomIndex);
    }

    static public String generate_adjectives() {
        Random random = new Random();
        int randomIndex = random.nextInt(chineseAdjectives.size());
        return chineseAdjectives.get(randomIndex);
    }

    static public String generate_adjectives_thing() {
        return generate_adjectives() + generate_thing();
    }

}
