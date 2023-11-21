package pj.dbs;


import pj.dbs.generate.generateSeller;
import pj.dbs.generate.generateUser;
import pj.dbs.generate.generateCommodity;
import pj.dbs.generate.generatePlatform;
import pj.dbs.generate.generateCommodityRelation;
import pj.dbs.generate.generateUserFavorite;
import pj.dbs.generate.generatePriceChange;
import pj.dbs.generate.generateMessage;

import java.text.ParseException;

public class RunAndGenerate {
    public static void main(String[] args) throws ParseException {
        generate_in_order(1000,200,100,4,2000,2000,500,1000);
    }

    //直接使用这个函数就可以了，由于是按照顺序产生的，故外键都是正确的。
    public static void generate_in_order(int user_num, int seller_num, int commodity_num, int platform_num, int commodity_relation_num,
                                         int user_favorite_num, int price_change_num, int message_num) throws ParseException {
        generateUser user = new generateUser();
        generateSeller seller = new generateSeller();
        generateCommodity commodity = new generateCommodity();
        generatePlatform platform = new generatePlatform();
        generateCommodityRelation commodityRelation = new generateCommodityRelation();
        generateUserFavorite userFavorite = new generateUserFavorite();
        generatePriceChange priceChange = new generatePriceChange();
        generateMessage message = new generateMessage();

        System.out.println("生成中...");

        //无引用表先生成
        user.generate(user_num);
        seller.generate(seller_num);
        commodity.generate(commodity_num);
        platform.generate(platform_num);

        //有引用表后生成
        commodityRelation.generate(commodity_relation_num);
        userFavorite.generate(user_favorite_num);
        priceChange.generate(price_change_num);
        message.generate(message_num);

        System.out.println("生成结束！");
    }
}