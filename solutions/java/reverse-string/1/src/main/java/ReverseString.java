class ReverseString {

    String reverse(String inputString) {
        StringBuilder s = new StringBuilder();

        for(int i = inputString.length() - 1; i >= 0; i--) {
            s.append(inputString.charAt(i));
        }

        return s.toString();
    } 
}