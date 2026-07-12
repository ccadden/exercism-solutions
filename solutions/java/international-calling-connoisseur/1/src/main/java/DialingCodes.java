import java.util.Map;
import java.util.HashMap;

public class DialingCodes {
    private Map<Integer, String> codes = new HashMap<Integer, String>();

    public Map<Integer, String> getCodes() {
        return codes;
    }

    public void setDialingCode(Integer code, String country) {
        this.codes.put(code, country);
        return;
    }

    public String getCountry(Integer code) {
        return codes.get(code);
    }

    public void addNewDialingCode(Integer code, String country) {
        if (findDialingCode(country) != null || codes.containsKey(code)) {
            return;
        }

        setDialingCode(code, country);
        return;
    }

    public Integer findDialingCode(String country) {
        for(Map.Entry<Integer, String> entry: codes.entrySet()) {
            if(entry.getValue() == country) {
                return entry.getKey();
            }
        }

        return null;
    }

    public void updateCountryDialingCode(Integer code, String country) {
        Integer existingCode = findDialingCode(country);

        if(existingCode != null) {
            codes.remove(existingCode);
            setDialingCode(code, country);
        }

        return;
    }
}
