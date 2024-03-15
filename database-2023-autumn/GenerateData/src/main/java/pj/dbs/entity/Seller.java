package pj.dbs.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class Seller {
    //主键：id
    //long id;       //auto
    String username;    //2-20
    String password;    //5-20
    String email;       //5-50
    String address;     //5-50

}
