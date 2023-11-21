package pj.dbs.generate;

import com.fasterxml.jackson.databind.ObjectMapper;
import pj.dbs.entity.PriceChange;
import pj.dbs.entity.Seller;
import pj.dbs.utils.GenerateNum;
import pj.dbs.utils.GenerateTime;

import java.io.File;
import java.io.IOException;
import java.text.ParseException;
import java.util.ArrayList;
import java.util.List;

public class generatePriceChange {
    public static List<PriceChange> priceChangeList = new ArrayList<>();

    public void generate(int num) throws ParseException {
        for(int i=1;i<=num;i++){
            PriceChange priceChange = new PriceChange();
            priceChange.setCommodity_id(Integer.toString(GenerateNum.generate_num(1,generateCommodity.commodityList.size())));
            priceChange.setChange_time(GenerateTime.generate_time());
            priceChange.setNew_price(GenerateNum.generate_num(1,10000));
            priceChangeList.add(priceChange);
        }
        try {
            ObjectMapper mapper = new ObjectMapper();
            mapper.writeValue(new File("Data/price_change.json"), priceChangeList);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
