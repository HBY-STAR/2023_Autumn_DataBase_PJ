package pj.dbs.generate;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;
import org.apache.commons.lang3.RandomStringUtils;
import pj.dbs.entity.Commodity;
import pj.dbs.entity.Seller;
import pj.dbs.utils.GenerateChinese;
import pj.dbs.utils.GenerateNum;
import pj.dbs.utils.GenerateTime;

import java.io.File;
import java.io.IOException;
import java.text.ParseException;
import java.util.ArrayList;
import java.util.List;

public class generateCommodity {
    public static List<Commodity> commodityList = new ArrayList<>();

    public void generate(int num) throws ParseException {
        for(int i=1;i<=num;i++){
            String item = GenerateChinese.generate_thing();
            String adj = GenerateChinese.generate_adjectives();
            Commodity commodity = new Commodity();

            //here
            commodity.setCategory(item);
            commodity.setDefault_name(adj+item);
            commodity.setProduce_address(RandomStringUtils.random(GenerateNum.generate_num(5,20),32,127,true,true));
            commodity.setProduce_at(GenerateTime.generate_time());

            commodityList.add(commodity);
        }
        try {
            ObjectMapper mapper = new ObjectMapper();
            mapper.registerModule(new JavaTimeModule());
            mapper.writeValue(new File("Data/commodity.json"), commodityList);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
