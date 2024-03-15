package pj.dbs.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class Platform {
    //主键：id
    //long id;     //auto
    String name;   //2-20
    String url;             //5-100
    String country;         //2-20
}
