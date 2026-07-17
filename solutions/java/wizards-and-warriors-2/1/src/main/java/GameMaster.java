public class GameMaster {
    public String describe(Character character) {
        return "You're a level " + Integer.toString(character.getLevel()) + " " + character.getCharacterClass() + " with " + Integer.toString(character.getHitPoints()) + " hit points.";
    }

    public String describe(Destination destination) {
        return "You've arrived at " + destination.getName() +", which has " + destination.getInhabitants() +" inhabitants.";
    }

    public String describe(TravelMethod travelMethod) {
        if(travelMethod == TravelMethod.HORSEBACK) {
            return "You're traveling to your destination on horseback.";
        } else {
            return "You're traveling to your destination by walking.";
        }
    }

    public String describe(Character character, Destination destination, TravelMethod travelMethod) {
        return describe(character) + " " + describe(travelMethod) + " " + describe(destination);
    }

    public String describe(Character character, Destination destination) {
        return describe(character, destination, TravelMethod.WALKING);
    }
}
