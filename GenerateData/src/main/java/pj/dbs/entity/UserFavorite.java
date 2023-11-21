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
    String user_id;         //引用
    String commodity_id;    //引用
    double price_limit;
    Timestamp create_time;
}
