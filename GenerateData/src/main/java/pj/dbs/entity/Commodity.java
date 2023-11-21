package pj.dbs.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class Commodity {
    //主键：commodity_id
    String commodity_id;    //auto
    String default_name;    //2-20
    String type;            //2-20
}
