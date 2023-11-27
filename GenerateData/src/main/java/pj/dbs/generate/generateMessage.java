package pj.dbs.generate;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;
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

            //here
            message.setUser_id(GenerateNum.generate_num(1,generateUser.userList.size()));
            message.setCommodity_item_id(GenerateNum.generate_num(1,generateCommodityItem.commodityItemList.size()));
            message.setCurrent_price(GenerateNum.generate_num(1,10000));
            message.setCreate_at(GenerateTime.generate_time());

            messageList.add(message);
        }
        try {
            ObjectMapper mapper = new ObjectMapper();
            mapper.registerModule(new JavaTimeModule());
            mapper.writeValue(new File("Data/message.json"), messageList);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
