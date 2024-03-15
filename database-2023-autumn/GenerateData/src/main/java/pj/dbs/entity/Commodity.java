package pj.dbs.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.sql.Timestamp;
import java.time.LocalDateTime;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class Commodity {
    //主键：id
    //long id;              //auto
    String default_name;     //2-20
    String produce_at;
    String produce_address;  //5-50
    String category;        //2-20
}
