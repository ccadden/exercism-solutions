public class LogLine {
    private String logLine;

    public LogLine(String logLine) {
        this.logLine = logLine;
    }

    public LogLevel getLogLevel() {
        if(logLine.startsWith("[INF")) {
            return LogLevel.INFO;
        }
        if (logLine.startsWith("[TRC")) {
            return LogLevel.TRACE;
        }
        if(logLine.startsWith("[DBG")) {
            return LogLevel.DEBUG;
        }
        if(logLine.startsWith("[WRN")) {
            return LogLevel.WARNING;
        }
        if(logLine.startsWith("[ERR")) {
            return LogLevel.ERROR;
        }
        if(logLine.startsWith("[FTL")) {
            return LogLevel.FATAL;
        }

        return LogLevel.UNKNOWN;
    }

    public String getOutputForShortLog() {
        LogLevel logLevel = getLogLevel();

        return Integer.toString(logLevel.getLogCode()) + ":" + logLine.substring(7);
    }
}
