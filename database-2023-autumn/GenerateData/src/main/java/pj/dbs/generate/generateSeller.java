package pj.dbs.generate;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;
import org.apache.commons.lang3.RandomStringUtils;
import pj.dbs.entity.Seller;
import pj.dbs.entity.User;
import pj.dbs.utils.GenerateChinese;
import pj.dbs.utils.GenerateNum;
import pj.dbs.utils.GenerateRandomString;

import java.io.File;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

public class generateSeller {
    public static List<Seller> sellerList = new ArrayList<>();

    public void generate(int num){
        for(int i=1;i<=num;i++){
            Seller seller = new Seller();

            //here
            seller.setUsername(GenerateChinese.generate_human_name());
            seller.setAddress(RandomStringUtils.random(GenerateNum.generate_num(5,20),32,127,true,true));
            seller.setPassword(RandomStringUtils.random(GenerateNum.generate_num(5,20),32,127,true,true));
            seller.setEmail(GenerateRandomString.generateEmail());

            sellerList.add(seller);
        }
        try {
            ObjectMapper mapper = new ObjectMapper();
            mapper.registerModule(new JavaTimeModule());
            mapper.writeValue(new File("Data/seller.json"), sellerList);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
