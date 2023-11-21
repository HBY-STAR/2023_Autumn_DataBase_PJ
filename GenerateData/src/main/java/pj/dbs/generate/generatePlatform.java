package pj.dbs.generate;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.apache.commons.lang3.RandomStringUtils;
import pj.dbs.entity.Platform;
import pj.dbs.entity.Seller;
import pj.dbs.utils.GenerateChinese;
import pj.dbs.utils.GenerateNum;
import pj.dbs.utils.GenerateRandomString;

import java.io.File;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

public class generatePlatform {
    public static List<Platform> platformList = new ArrayList<>();

    public void generate(int num){
        for(int i=1;i<=num;i++){
            Platform platform = new Platform();
            platform.setPlatform_id(Integer.toString(i));
            platform.setPlatform_name(RandomStringUtils.random(GenerateNum.generate_num(2,4),"ABCDEFGHIJKLMNOPQRSTUVWXYZ"));
            platform.setCountry(GenerateChinese.generate_country());
            platform.setUrl(GenerateRandomString.generateUrl());
            platformList.add(platform);
        }
        try {
            ObjectMapper mapper = new ObjectMapper();
            mapper.writeValue(new File("Data/platform.json"), platformList);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
