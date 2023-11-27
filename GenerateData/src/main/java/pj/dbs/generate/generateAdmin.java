package pj.dbs.generate;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;
import org.apache.commons.lang3.RandomStringUtils;
import pj.dbs.entity.Admin;
import pj.dbs.entity.Commodity;
import pj.dbs.utils.GenerateChinese;
import pj.dbs.utils.GenerateNum;
import pj.dbs.utils.GenerateTime;

import java.io.File;
import java.io.IOException;
import java.text.ParseException;
import java.util.ArrayList;
import java.util.List;

public class generateAdmin {
    public static List<Admin> adminList = new ArrayList<>();

    public void generate() throws ParseException {

        //here
        Admin admin1 = new Admin("admin","admin");
        Admin admin2 = new Admin("manager","manager");

        adminList.add(admin1);
        adminList.add(admin2);

        try {
            ObjectMapper mapper = new ObjectMapper();
            mapper.registerModule(new JavaTimeModule());
            mapper.writeValue(new File("Data/admin.json"), adminList);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
