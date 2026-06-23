public class CarsAssemble {
    static int CARS_PER_HOUR = 221;
    static int MINUTES_PER_HOUR = 60;

    public double productionRatePerHour(int speed) {
        return CARS_PER_HOUR * speed * errorRate(speed);
    }

    public int workingItemsPerMinute(int speed) {
        return (int) productionRatePerHour(speed) / MINUTES_PER_HOUR;
    }

    double errorRate(int speed) {
        if (speed >= 1 && speed <= 4) {
            return 1;
        } else if (speed >=5 && speed <= 8) {
            return 0.9;
        } else if (speed == 9) {
            return 0.8;
        }
        else {
            return 0.77;
        }
    }
}
