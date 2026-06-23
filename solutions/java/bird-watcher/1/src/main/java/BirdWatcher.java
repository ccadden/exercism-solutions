
class BirdWatcher {
    private final int[] birdsPerDay;

    public BirdWatcher(int[] birdsPerDay) {
        this.birdsPerDay = birdsPerDay.clone();
    }

    public static int[] getLastWeek() {
        return new int[] {0, 2, 5, 3, 7, 8, 4};
    }

    public int getToday() {
        return this.birdsPerDay[this.birdsPerDay.length - 1];
    }

    public void incrementTodaysCount() {
        this.birdsPerDay[this.birdsPerDay.length - 1] += 1;
    }

    public boolean hasDayWithoutBirds() {
        for (int count : this.birdsPerDay) {
            if(count == 0) {
                return true;
            }
        }

        return false;
    }

    public int getCountForFirstDays(int numberOfDays) {
        int birdCount = 0;

        for (int i = 0; i < min(this.birdsPerDay.length, numberOfDays); i++) {
            birdCount += this.birdsPerDay[i];
        }

        return birdCount;
    }

    public int getBusyDays() {
        int busyDays = 0;

        for (int count: this.birdsPerDay) {
            if(count >= 5) {
                busyDays += 1;
            }
        }

        return busyDays;
    }

    private int min(int a, int b) {
        return a < b ? a : b;
    }
}
