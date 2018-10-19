package tech.kbtg;

public class MyThread extends Thread {
    public MyThread(String name) {
        super(name);
    }

    @Override
    public void run() {
        try {
            doDBProcessing();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }

    private void doDBProcessing() throws InterruptedException {
        Thread.sleep(5000);
    }
}