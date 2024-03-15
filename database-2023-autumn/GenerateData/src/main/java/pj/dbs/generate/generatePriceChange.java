package pj.dbs.generate;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;
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

            //here
            priceChange.setCommodity_item_id(GenerateNum.generate_num(1,generateCommodity.commodityList.size()));
            priceChange.setUpdate_at(GenerateTime.generate_time());
            priceChange.setNew_price(GenerateNum.generate_num(1,10000));

            priceChangeList.add(priceChange);
        }
        try {
            ObjectMapper mapper = new ObjectMapper();
            mapper.registerModule(new JavaTimeModule());
            mapper.writeValue(new File("Data/price_change.json"), priceChangeList);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
