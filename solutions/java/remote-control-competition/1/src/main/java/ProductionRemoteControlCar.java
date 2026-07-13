class ProductionRemoteControlCar implements RemoteControlCar, Comparable<ProductionRemoteControlCar> {

    private int distance = 0;
    private int victories = 0;

    public void drive() {
        this.distance += 10;
    }

    public int getDistanceTravelled() {
        return distance;
    }

    public int getNumberOfVictories() {
        return victories;
    }

    public void setNumberOfVictories(int numberOfVictories) {
        this.victories = numberOfVictories;
    }

    public int compareTo(ProductionRemoteControlCar car) {
        if (car.getNumberOfVictories() < victories) {
            return 1;
        } else if (car.getNumberOfVictories() > victories) {
            return -1;
        } else {
            return 0;
        }
    }
}
