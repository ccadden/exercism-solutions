import java.util.List;
import java.util.Set;
import java.util.HashSet;
import java.util.stream.Collectors;

class GottaSnatchEmAll {

    static Set<String> newCollection(List<String> cards) {
        return new HashSet<>(cards);
    }

    static boolean addCard(String card, Set<String> collection) {
        return collection.add(card);
    }

    static boolean canTrade(Set<String> myCollection, Set<String> theirCollection) {
        return myCollection.stream().anyMatch(e -> !theirCollection.contains(e)) && theirCollection.stream().anyMatch(e -> !myCollection.contains(e));
    }

    static Set<String> commonCards(List<Set<String>> collections) {
        Set<String> intersection = new HashSet<>(collections.getFirst());

        collections.stream().skip(1).forEach(intersection::retainAll);

        return intersection;
    }

    static Set<String> allCards(List<Set<String>> collections) {
        return collections.stream().flatMap(Set::stream).collect(Collectors.toSet());
    }
}
