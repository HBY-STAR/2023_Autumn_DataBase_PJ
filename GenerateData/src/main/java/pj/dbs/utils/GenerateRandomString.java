package pj.dbs.utils;

import org.apache.commons.lang3.RandomStringUtils;

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
}
