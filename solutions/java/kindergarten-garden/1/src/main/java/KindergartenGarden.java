import java.util.List;
import java.util.ArrayList;
import java.util.HashMap;

class KindergartenGarden {
    private String[] STUDENT_ROLL = {
        "Alice",
        "Bob",
        "Charlie",
        "David",
        "Eve",
        "Fred",
        "Ginny",
        "Harriet",
        "Ileana",
        "Joseph",
        "Kincaid",
        "Larry"
    };

    private String[] garden;

    KindergartenGarden(String garden) {
        this.garden = garden.split("\n");
    }

    List<Plant> getPlantsOfStudent(String student) {
        List<Plant> plants = new ArrayList(4);

        int studentOffset = getOffset(student);

        for (int i = 0; i < 2; i++) {
            for (int j = studentOffset; j < studentOffset + 2; j++) {
                plants.add(Plant.getPlant(garden[i].charAt(j)));
            }
        }

        return plants;
    }

    private int getOffset(String student) {
        for(int i = 0; i < STUDENT_ROLL.length; i++) {
            if (STUDENT_ROLL[i] == student) {
                return i * 2;
            }
        }

        return -1;
    }

}
