class ArmstrongNumbers {

    boolean isArmstrongNumber(int numberToCheck) {
        String numString = Integer.toString(numberToCheck);
        int lenString = numString.length();
        int sum = 0;

        for(int i = 0; i < lenString; i++) {
            sum += pow(Integer.parseInt(Character.toString(numString.charAt(i)), 10), lenString);
        }

        return sum == numberToCheck;
    }

    int pow(int num, int power) {
        int product = 1;

        for (int i=0; i < power; i++) {
            product *= num;
        }

        return product;
    }

}
