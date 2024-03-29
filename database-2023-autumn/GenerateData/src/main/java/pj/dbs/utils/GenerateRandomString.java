package pj.dbs.utils;

import org.apache.commons.lang3.RandomStringUtils;

import java.util.Random;

public class GenerateRandomString {
    public static String generateEmail(){
        String random_1 = RandomStringUtils.random(GenerateNum.generate_num(5,10),32,127,true,true);
        String random_2 = RandomStringUtils.random(GenerateNum.generate_num(2,6),32,127,true,true);
        return  random_1+"@"+random_2+".com";
    }
    public static String generateUrl(){
        String random_1 =  RandomStringUtils.random(GenerateNum.generate_num(3,8),32,127,true,true);
        String random_2 =  RandomStringUtils.random(GenerateNum.generate_num(3,8),32,127,true,true);
        String random_3 =  RandomStringUtils.random(GenerateNum.generate_num(3,8),32,127,true,true);
        return "https://"+random_1+"."+random_2+"."+random_3+".com";
    }

    public static String generatePhone(){
        // 手机号的第一位为1，后面随机生成10位数字
        StringBuilder phoneNumber = new StringBuilder("1");

        Random random = new Random();

        for (int i = 0; i < 10; i++) {
            int digit = random.nextInt(10);
            phoneNumber.append(digit);
        }

        return phoneNumber.toString();
    }
}
