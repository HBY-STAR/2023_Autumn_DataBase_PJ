package pj.dbs.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.sql.Timestamp;
@Data
@NoArgsConstructor
@AllArgsConstructor
public class UserFavorite {
    //主键：user_id + commodity_id
    String user_id;         //外键
    String commodity_id;    //外键
    double price_limit;
    Timestamp create_time;
}
