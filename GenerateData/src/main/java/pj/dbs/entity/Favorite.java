package pj.dbs.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.sql.Timestamp;
import java.time.LocalDateTime;
import java.util.List;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class Favorite {
    //主键：user_id + commodity_item_id
    long user_id;         //外键
    long commodity_item_id;    //外键
    double price_limit;
    String update_at;
}
