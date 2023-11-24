package pj.dbs.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class Seller {
    //主键：seller_id
    String seller_id;       //auto
    String seller_name;     //2-20
    String password;        //5-20
    String address;         //5-50

}
