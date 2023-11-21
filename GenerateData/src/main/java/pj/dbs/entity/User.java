package pj.dbs.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class User {
    //主键：user_id
    String user_id;     //auto
    String user_name;   //2-20
    String password;    //5-20
    int age;
    boolean gender;
    String email;       //5-50
}
