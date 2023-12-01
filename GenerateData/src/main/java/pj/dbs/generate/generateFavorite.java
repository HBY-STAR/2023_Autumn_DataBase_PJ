package pj.dbs.generate;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;
import pj.dbs.entity.Favorite;
import pj.dbs.utils.GenerateNum;
import pj.dbs.utils.GenerateTime;

import java.io.File;
import java.io.IOException;
import java.text.ParseException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

public class generateFavorite {
    public static List<Favorite> favoriteList = new ArrayList<>();

//     check user doesn't have the same favorite
    public static HashMap<Long, Long> user_favorite = new HashMap<>();

    public void generate(int num) throws ParseException {
        for (int i = 1; i <= num; i++) {
            Favorite favorite = new Favorite();

            //here
            while (true) {
                favorite.setUser_id(GenerateNum.generate_num(1, generateUser.userList.size()));
                favorite.setCommodity_item_id(GenerateNum.generate_num(1, generateCommodityItem.commodityItemList.size()));
                if (user_favorite.containsKey(favorite.getUser_id())) {
                    if (user_favorite.get(favorite.getUser_id()) == favorite.getCommodity_item_id()) {
                        continue;
                    }
                }
                user_favorite.put(favorite.getUser_id(), favorite.getCommodity_item_id());
                break;
            }
            favorite.setCreate_at(GenerateTime.generate_time());
            favorite.setPrice_limit(GenerateNum.generate_num(1, 10000));

            favoriteList.add(favorite);
        }
        try {
            ObjectMapper mapper = new ObjectMapper();
            mapper.registerModule(new JavaTimeModule());
            mapper.writeValue(new File("Data/favorite.json"), favoriteList);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
