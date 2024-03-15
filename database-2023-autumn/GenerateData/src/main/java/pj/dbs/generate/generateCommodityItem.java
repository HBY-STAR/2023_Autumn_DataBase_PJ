package pj.dbs.generate;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;
import pj.dbs.entity.CommodityItem;
import pj.dbs.utils.GenerateChinese;
import pj.dbs.utils.GenerateNum;
import pj.dbs.utils.GenerateRandomString;
import pj.dbs.utils.GenerateTime;

import java.io.File;
import java.io.IOException;
import java.sql.Timestamp;
import java.text.ParseException;
import java.time.LocalDateTime;
import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;

public class generateCommodityItem {
    public static List<CommodityItem> commodityItemList = new ArrayList<>();

    public static HashSet<String> commodityItemSet = new HashSet<>();

    public String commodityItemString(Long commodity_id, Long seller_id, Long platform_id) {
        return commodity_id + " " + seller_id + " " + platform_id;
    }

    public void generate(int num) throws ParseException {
        for (int i = 1; i <= num; i++) {
            CommodityItem commodityItem = new CommodityItem();

            //here
            while (true) {
                commodityItem.setCommodity_id(GenerateNum.generate_num(1, generateCommodity.commodityList.size()));
                commodityItem.setSeller_id(GenerateNum.generate_num(1, generateSeller.sellerList.size()));
                commodityItem.setPlatform_id(GenerateNum.generate_num(1, generatePlatform.platformList.size()));

                String s = commodityItemString(commodityItem.getCommodity_id(), commodityItem.getSeller_id(), commodityItem.getPlatform_id());
                if (commodityItemSet.contains(s)) {
                    continue;
                }
                commodityItemSet.add(s);
                break;
            }

            commodityItem.setItem_name(GenerateChinese.generate_adjectives_thing());
            commodityItem.setPrice(GenerateNum.generate_num(1, 10000));
            commodityItem.setUpdate_at(GenerateTime.generate_time());

            commodityItemList.add(commodityItem);
        }
        try {
            ObjectMapper mapper = new ObjectMapper();
            mapper.registerModule(new JavaTimeModule());
            mapper.writeValue(new File("Data/commodity_item.json"), commodityItemList);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
