class Badge {
    public String print(Integer id, String name, String department) {
        StringBuilder b = new StringBuilder();

        if(id != null) {
            b.append("[" + id + "] - ");
        }

        b.append(name);

        if (department != null) {
                b.append(" - " + department.toUpperCase());
        } else {
                b.append(" - OWNER");

        }

        return b.toString();
    }
}
