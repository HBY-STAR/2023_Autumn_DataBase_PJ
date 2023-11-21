package pj.dbs.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class Platform {
    //主键：platform_id
    String platform_id;     //auto
    String platform_name;   //2-20
    String url;             //5-100
    String country;         //2-20
}
