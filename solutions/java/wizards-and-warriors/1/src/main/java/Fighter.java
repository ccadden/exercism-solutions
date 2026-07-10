class Fighter {

    boolean isVulnerable() {
        return true;
    }

    int getDamagePoints(Fighter fighter) {
        return 1;
    }
}

class Warrior extends Fighter {
    
    @Override
    public String toString() {
        return "Fighter is a Warrior";
    }

    @Override
    boolean isVulnerable() {
        return false;
    }

    @Override
    int getDamagePoints(Fighter fighter) {
        if (fighter.isVulnerable()) {
            return 10;
        } else {
            return 6;
        }
    }
    
}

// TODO: define the Wizard class
class Wizard extends Fighter {
    private boolean prepared = false;
    
    @Override
    public String toString() {
        return "Fighter is a Wizard";
    }

    @Override
    public boolean isVulnerable() {
        return !prepared;
    }

    public void prepareSpell() {
        this.prepared = true;
    }

    @Override
    public int getDamagePoints(Fighter fighter) {
        if(prepared) {
            return 12;
        } else {
            return 3;
        }
    }
}
