package hby.dbs;

import hby.dbs.utils.GenerateChinese;
import hby.dbs.utils.GenerateEnglish;
import hby.dbs.utils.GenerateTime;
import java.text.ParseException;
import java.text.SimpleDateFormat;

public class Main {
    public static void main(String[] args) throws ParseException {
        for (int i=0;i<10;i++) {
            System.out.println(GenerateChinese.generate_human_name() + "有一个" + GenerateChinese.generate_adjectives_thing());
            System.out.println(GenerateEnglish.generate_name()+" has a "+GenerateEnglish.generate_adjectives_thing());
            SimpleDateFormat dateFormat = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
            System.out.println(dateFormat.format(GenerateTime.generate_time()));
        }
    }
}