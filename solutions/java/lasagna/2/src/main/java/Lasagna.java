public class Lasagna {
    private static final int EXPECTED_MINUTES_IN_OVEN = 40;
    private static final int PREPARATION_TIME_PER_LAYER = 2;

    public int expectedMinutesInOven() {
        return EXPECTED_MINUTES_IN_OVEN;
    }

    public int remainingMinutesInOven(int timeInOven) {
        return expectedMinutesInOven() - timeInOven;
    }

    public int preparationTimeInMinutes(int numLayers) {
        return PREPARATION_TIME_PER_LAYER * numLayers;
    }

    public int totalTimeInMinutes(int numLayers, int timeInOven) {
        return preparationTimeInMinutes(numLayers) + timeInOven;
    }
}
