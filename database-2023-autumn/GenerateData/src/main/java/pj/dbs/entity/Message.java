package pj.dbs.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.sql.Timestamp;
import java.time.LocalDateTime;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class Message {
    //主键：id
    //long id;            //auto
    long user_id;         //外键
    long commodity_item_id;    //外键
    double current_price;
    String create_at;
}
