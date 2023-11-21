package pj.dbs.generate;

import com.fasterxml.jackson.databind.ObjectMapper;
import pj.dbs.entity.Commodity;
import pj.dbs.entity.Seller;
import pj.dbs.utils.GenerateChinese;

import java.io.File;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

public class generateCommodity {
    public static List<Commodity> commodityList = new ArrayList<>();

    public void generate(int num){
        for(int i=1;i<=num;i++){
            String item = GenerateChinese.generate_thing();
            String adj = GenerateChinese.generate_adjectives();
            Commodity commodity = new Commodity();
            commodity.setCommodity_id(Integer.toString(i));
            commodity.setType(item);
            commodity.setDefault_name(adj+item);
            commodityList.add(commodity);
        }
        try {
            ObjectMapper mapper = new ObjectMapper();
            mapper.writeValue(new File("Data/commodity.json"), commodityList);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
