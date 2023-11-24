package pj.dbs.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.sql.Timestamp;
@Data
@NoArgsConstructor
@AllArgsConstructor
public class Message {
    //主键：user_id + send_time
    String user_id;         //外键
    String commodity_id;    //外键
    String seller_id;       //外键
    String platform_id;     //外键
    double current_price;
    Timestamp send_time;
}
