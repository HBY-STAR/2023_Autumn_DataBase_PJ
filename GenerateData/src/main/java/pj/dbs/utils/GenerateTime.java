package pj.dbs.utils;

import javax.xml.crypto.Data;
import java.sql.Timestamp;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Calendar;
import java.util.Date;
import java.util.Random;

public class GenerateTime {
    public static Timestamp generate_time() throws ParseException {
        //设置指定时间
        SimpleDateFormat sdf = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
        Date date_start = sdf.parse("2023-01-01 00:00:00");
        Date date_end = sdf.parse("2023-11-20 00:00:00");

        Date randomDate = generateRandomDate(date_start,date_end);
        return new Timestamp(randomDate.getTime());
    }

    // 生成指定范围内的随机日期
    private static Date generateRandomDate(Date startDate, Date endDate) {
        Random random = new Random();
        long startMillis = startDate.getTime();
        long endMillis = endDate.getTime();
        long randomMillis = startMillis + (long) (random.nextDouble() * (endMillis - startMillis));

        return new Date(randomMillis);
    }
}
