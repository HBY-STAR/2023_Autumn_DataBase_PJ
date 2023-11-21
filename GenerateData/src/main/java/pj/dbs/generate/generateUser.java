package pj.dbs.generate;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.apache.commons.lang3.RandomStringUtils;
import pj.dbs.entity.User;
import pj.dbs.utils.GenerateChinese;
import pj.dbs.utils.GenerateNum;
import pj.dbs.utils.GenerateRandomString;

import java.io.File;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

public class generateUser {
    public static List<User> userList = new ArrayList<>();

    public void generate(int num){
        for(int i=1;i<=num;i++){
            User user = new User();
            user.setUser_id(Integer.toString(i));
            user.setUser_name(GenerateChinese.generate_human_name());
            user.setAge(GenerateNum.generate_num(10,80));
            user.setPassword(RandomStringUtils.random(GenerateNum.generate_num(5,20),32,127,true,true));
            user.setGender(i % 2 == 0);
            user.setEmail(GenerateRandomString.generateEmail());
            userList.add(user);
        }
        try {
            ObjectMapper mapper = new ObjectMapper();
            mapper.writeValue(new File("Data/user.json"), userList);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
