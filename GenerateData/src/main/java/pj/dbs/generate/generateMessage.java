package pj.dbs.generate;

import com.fasterxml.jackson.databind.ObjectMapper;
import pj.dbs.entity.Message;
import pj.dbs.entity.Seller;
import pj.dbs.utils.GenerateNum;
import pj.dbs.utils.GenerateTime;

import java.io.File;
import java.io.IOException;
import java.sql.Timestamp;
import java.text.ParseException;
import java.time.LocalDateTime;
import java.util.ArrayList;
import java.util.List;

public class generateMessage {
    public static List<Message> messageList = new ArrayList<>();

    public void generate(int num) throws ParseException {
        for(int i=1;i<=num;i++){
            Message message = new Message();
            message.setUser_id(Integer.toString(GenerateNum.generate_num(1,generateUser.userList.size())));
            message.setSeller_id(Integer.toString(GenerateNum.generate_num(1,generateSeller.sellerList.size())));
            message.setCommodity_id(Integer.toString(GenerateNum.generate_num(1,generateCommodity.commodityList.size())));
            message.setPlatform_id(Integer.toString(GenerateNum.generate_num(1,generatePlatform.platformList.size())));
            message.setCurrent_price(GenerateNum.generate_num(1,10000));
            message.setSend_time(GenerateTime.generate_time());
            messageList.add(message);
        }
        try {
            ObjectMapper mapper = new ObjectMapper();
            mapper.writeValue(new File("Data/message.json"), messageList);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
