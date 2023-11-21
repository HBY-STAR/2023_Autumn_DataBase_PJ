package pj.dbs.generate;

import com.fasterxml.jackson.databind.ObjectMapper;
import pj.dbs.entity.Seller;
import pj.dbs.entity.UserFavorite;
import pj.dbs.utils.GenerateNum;
import pj.dbs.utils.GenerateTime;

import java.io.File;
import java.io.IOException;
import java.text.ParseException;
import java.util.ArrayList;
import java.util.List;

public class generateUserFavorite {
    public static List<UserFavorite> userFavoriteList = new ArrayList<>();

    public void generate(int num) throws ParseException {
        for(int i=1;i<=num;i++){
            UserFavorite userFavorite = new UserFavorite();
            userFavorite.setUser_id(Integer.toString(GenerateNum.generate_num(1,generateUser.userList.size())));
            userFavorite.setCreate_time(GenerateTime.generate_time());
            userFavorite.setCommodity_id(Integer.toString(GenerateNum.generate_num(1,generateCommodity.commodityList.size())));
            userFavorite.setPrice_limit(GenerateNum.generate_num(1,10000));
            userFavoriteList.add(userFavorite);
        }
        try {
            ObjectMapper mapper = new ObjectMapper();
            mapper.writeValue(new File("Data/user_favorite.json"), userFavoriteList);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
