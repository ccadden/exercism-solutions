class SqueakyClean {
    static String clean(String identifier) {
        StringBuilder b = new StringBuilder();
        boolean capitalizeNext = false;

        for(int i = 0; i < identifier.length(); i++) {
            char c = identifier.charAt(i);

            if (Character.isLetterOrDigit(c)) {
                if (capitalizeNext) {
                    capitalizeNext = false;
                    b.append(convertLeetspeak(Character.toUpperCase(c)));
                } else {
                    b.append(convertLeetspeak(c));
                }
            } else {
                if (Character.isWhitespace(c)) {
                    b.append('_');
                }

                if (c == '-') {
                    capitalizeNext = true;
                }
            }
        }

        return b.toString();
    }

    static private char convertLeetspeak(char c) {
        switch (c) {
            case '0':
                return 'o';
            case '1':
                return 'l';
            case '3':
                return 'e';
            case '4':
                return 'a';
            case '7':
                return 't';
            default:
                return c;
        }
    }
}
