package pj.dbs.generate;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.apache.commons.lang3.RandomStringUtils;
import pj.dbs.entity.CommodityRelation;
import pj.dbs.entity.Seller;
import pj.dbs.utils.GenerateChinese;
import pj.dbs.utils.GenerateNum;
import pj.dbs.utils.GenerateTime;

import java.io.File;
import java.io.IOException;
import java.sql.Timestamp;
import java.text.ParseException;
import java.time.LocalDateTime;
import java.util.ArrayList;
import java.util.List;

public class generateCommodityRelation {
    public static List<CommodityRelation> commodityRelationList = new ArrayList<>();

    public void generate(int num) throws ParseException {
        for(int i=1;i<=num;i++){
            CommodityRelation commodityRelation = new CommodityRelation();
            commodityRelation.setCommodity_id(Integer.toString(GenerateNum.generate_num(1,generateCommodity.commodityList.size())));
            commodityRelation.setSeller_id(Integer.toString(GenerateNum.generate_num(1,generateSeller.sellerList.size())));
            commodityRelation.setPlatform_id(Integer.toString(GenerateNum.generate_num(1,generatePlatform.platformList.size())));
            commodityRelation.setCommodity_name(GenerateChinese.generate_adjectives_thing());
            commodityRelation.setPrice(GenerateNum.generate_num(1,10000));
            commodityRelation.setProduction_address(RandomStringUtils.random(GenerateNum.generate_num(5,20),32,127,true,true));
            commodityRelation.setUpdate_time(Timestamp.valueOf(LocalDateTime.now()));
            commodityRelation.setProduction_date(GenerateTime.generate_time());

            commodityRelationList.add(commodityRelation);
        }
        try {
            ObjectMapper mapper = new ObjectMapper();
            mapper.writeValue(new File("Data/commodity_relation.json"), commodityRelationList);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
