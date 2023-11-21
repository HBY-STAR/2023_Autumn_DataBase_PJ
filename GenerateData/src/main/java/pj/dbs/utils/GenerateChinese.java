package pj.dbs.utils;

import org.apache.commons.lang3.RandomStringUtils;

import java.util.Arrays;
import java.util.List;
import java.util.Random;

public class GenerateChinese {
    //常用名
    public static final String common="云深婉雅明艳浩晨涵雨婷洋悠然韵瑶思淼婕舒静娅蕾芝伟伦琳翰涛晨璐婧梓萌珂晗瑾菲茜晨悦佳琪鑫嘉寅雅萱浩然志强梦楠欣然雅芳乐怡颖慧晓婷洁灵媛婧佳蓉妍琳梦婷明阳博文春芳盈宸欢佳欣韵涵欣悦丽莎忆柳梦佳欣静艺瑶梦露婷涵心静娴美丽思婉悦萌彤依楠润婷梦蕾露丹翠文婷静萍玲英婕雅紫菲晓芸明慧乐然晨涵海文欢雅馨雅楠馨儿娅佳璇雅雯奕璇芸楠菲瑶琴瑾涵琳晓春欢欢乐涵阳凯文伟毅世杰海涛佳辉文轩洋洋昊天泽宇鑫鑫锦程梓轩博涵晓明哲翔婧琪涵涵弘涛宇航越泽凯安梓晨昊然瑞祥梦洁梦璐梦婷若涵宇霆敬亭梦瑶梦佳依依涵韵洋洋悠然婉婷鑫婷博文悠悠文然琪涵天佳楠鑫婷婉佳欣琪洋涵明博欣然婕若涵玥晨露露琦若涵若涵洋欣悦涵欢";
    //常用姓氏
    public static final String surname="王李张刘陈杨黄赵吴周徐孙马朱胡郭何高林罗郑梁谢宋唐许韩冯邓曹彭曾肖田董袁潘于蒋蔡余杜叶程苏魏吕丁任沈姚卢姜崔钟谭陆汪戴金贾夏邱方邹熊孟秦白江阎薛尹段雷黎史龙陶贺顾毛郝龚邵万钱严赖覃洪武莫孔向胥倪和凤房裘缪解应宗丰邢路柴袁古卜冷甄靳左祁纪屠公宣松沐邴豆党富须修邰都雍卓谷奚丛岳寸景示巫柏瞿阮桂溥万俟赫";
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
            "中国","美国","英国","法国","俄罗斯","日本","印度","巴西","加拿大","德国"
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


    static public String generate_human_name(){
        String name_1 = RandomStringUtils.random(1,surname);

        int name_2_len = GenerateNum.generate_num(1,2);
        String name_2 = RandomStringUtils.random(name_2_len,common);
        return name_1+name_2;
    }

    static public String generate_country(){
        Random random = new Random();
        int randomIndex = random.nextInt(country.size());
        return country.get(randomIndex);
    }

    static public String generate_thing(){
        Random random = new Random();
        int randomIndex = random.nextInt(chineseItemList.size());
        return chineseItemList.get(randomIndex);
    }

    static public String generate_adjectives(){
        Random random = new Random();
        int randomIndex = random.nextInt(chineseAdjectives.size());
        return chineseAdjectives.get(randomIndex);
    }

    static public String generate_adjectives_thing(){
        return generate_adjectives()+generate_thing();
    }

}
