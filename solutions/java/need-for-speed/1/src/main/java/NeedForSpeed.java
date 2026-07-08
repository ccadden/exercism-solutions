class NeedForSpeed {
    public int speed;
    public int batteryDrain;
    private int batteryCharge = 100;
    private int distance;
    
    NeedForSpeed(int speed, int batteryDrain) {
        this.speed = speed;
        this.batteryDrain = batteryDrain;
    }

    public boolean batteryDrained() {
        return this.batteryCharge < this.batteryDrain;
    }

    public int distanceDriven() {
        return this.distance;
    }

    public void drive() {
        if (this.batteryDrained()) {
            return;
        }
        
        this.distance += this.speed;
        this.batteryCharge -= this.batteryDrain;
    }

    public static NeedForSpeed nitro() {
        return new NeedForSpeed(50, 4);
    }
}

class RaceTrack {
    private int distance;
    
    RaceTrack(int distance) {
        this.distance = distance;
    }

    public boolean canFinishRace(NeedForSpeed car) {
        return ((this.distance + car.speed - 1) / car.speed) * car.batteryDrain <= 100;
    }
}
