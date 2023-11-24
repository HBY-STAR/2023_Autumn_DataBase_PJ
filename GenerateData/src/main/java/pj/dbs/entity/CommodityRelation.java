package pj.dbs.entity;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.sql.Timestamp;
@Data
@NoArgsConstructor
@AllArgsConstructor
public class CommodityRelation {
    //主键：commodity_id + seller_id + platform_id
    String commodity_id;        //外键
    String commodity_name;      //2-20
    String seller_id;           //外键
    String platform_id;         //外键
    double price;
    Timestamp update_time;
    Timestamp production_date;
    String production_address;  //5-50
}
