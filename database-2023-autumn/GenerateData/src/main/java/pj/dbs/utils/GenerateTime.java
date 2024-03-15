package pj.dbs.utils;

import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.time.Instant;
import java.time.LocalDateTime;
import java.time.ZoneId;
import java.util.Date;
import java.util.Random;

public class GenerateTime {
    public static String generate_time() throws ParseException {
        //设置指定时间
        SimpleDateFormat sdf = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
        Date date_start = sdf.parse("2023-01-01 00:00:00");
        Date date_end = sdf.parse("2023-11-27 00:00:00");

        Date generate= generateRandomDate(date_start,date_end);

        // 将Date转换为Instant
        // Instant instant = generate.toInstant();
        // 将Instant转换为LocalDateTime
        return sdf.format(generate);
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
