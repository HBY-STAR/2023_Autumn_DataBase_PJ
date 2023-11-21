package hby.dbs.utils;

import java.util.Random;

public class GenerateNum {
    public static int generate_num(int min,int max){
        if (min > max) {
            throw new IllegalArgumentException("最小值不能大于最大值");
        }
        Random rand = new Random();
        return rand.nextInt((max - min) + 1) + min;
    }
}
