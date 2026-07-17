public enum LogLevel {
    TRACE(1),
    DEBUG(2),
    INFO(4),
    WARNING(5),
    ERROR(6),
    FATAL(42),
    UNKNOWN(0);

    private final int logCode;

    LogLevel(int logCode) {
        this.logCode = logCode;
    }

    public int getLogCode() {
        return logCode;
    }
}
