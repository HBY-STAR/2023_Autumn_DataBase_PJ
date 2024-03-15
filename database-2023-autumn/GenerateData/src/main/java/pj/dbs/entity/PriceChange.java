package pj.dbs.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.sql.Timestamp;
import java.time.LocalDateTime;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class PriceChange {
    //主键：id
    //long id;
    long commodity_item_id;    //外键
    double new_price;
    String update_at;
}
