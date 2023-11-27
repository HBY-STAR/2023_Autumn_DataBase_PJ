package pj.dbs.entity;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.sql.Timestamp;
import java.time.LocalDateTime;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class CommodityItem {
    //主键：commodity_id + seller_id + platform_id
    //long id;  //auto
    long commodity_id;
    long platform_id;
    long seller_id;
    String item_name;
    double price;
    String update_at;
}
