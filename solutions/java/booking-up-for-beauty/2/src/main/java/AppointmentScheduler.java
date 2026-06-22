import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;

class AppointmentScheduler {

    private static final DateTimeFormatter DATE_PRINTER = DateTimeFormatter.ofPattern("'You have an appointment on' EEEE',' MMMM d',' yyyy', at' h:mm a'.'");

    private static final DateTimeFormatter DATE_PARSER = DateTimeFormatter.ofPattern("MM/dd/yyyy HH:mm:ss");

    public LocalDateTime schedule(String appointmentDateDescription) {
        return LocalDateTime.parse(appointmentDateDescription, DATE_PARSER);
    }

    public boolean hasPassed(LocalDateTime appointmentDate) {
        return appointmentDate.isBefore(LocalDateTime.now());
    }

    public boolean isAfternoonAppointment(LocalDateTime appointmentDate) {
        return appointmentDate.getHour() > 11 && appointmentDate.getHour() < 18;
    }

    public String getDescription(LocalDateTime appointmentDate) {
        return DATE_PRINTER.format(appointmentDate);
    }

    public LocalDate getAnniversaryDate() {
        return LocalDate.of(LocalDate.now().getYear(), 9, 15);
    }
}
