package tech.kbtg;

import tech.kbtg.MyThread;

public class ThreadRunExample {

    public static void main(String[] args){
        Thread t = new MyThread("t");
        t.start();
    }
}
