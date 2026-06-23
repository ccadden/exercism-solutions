public class LogLevels {
    
    public static String message(String logLine) {
        int i = 0;

        for(; i < logLine.length(); i++) {
            if(logLine.charAt(i) == ' ') {
                break;
            }
        }

        return logLine.substring(i).trim();
    }

    public static String logLevel(String logLine) {
        switch(logLine.charAt(1)) {
            case 'E':
              return "error";
            case 'I':
              return "info";
            case 'W':
              return "warning";
            default:
              throw new Error("unsupported message type");
        }
    }

    public static String reformat(String logLine) {
        return message(logLine) + " (" + logLevel(logLine) + ")";
    }
}
