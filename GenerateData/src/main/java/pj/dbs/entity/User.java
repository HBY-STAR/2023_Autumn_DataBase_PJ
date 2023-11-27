package pj.dbs.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class User {
    //主键：id
    //String id;     //auto
    String user_name;   //2-20
    String password;    //5-20
    String email;       //5-50
    int age;
    boolean gender;
    String phone;     //11
}
