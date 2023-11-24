package pj.dbs.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.sql.Timestamp;
@Data
@NoArgsConstructor
@AllArgsConstructor
public class PriceChange {
    //主键：commodity_id + change_time
    String commodity_id;    //外键
    Timestamp change_time;
    double new_price;
}
