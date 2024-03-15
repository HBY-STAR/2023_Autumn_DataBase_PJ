package pj.dbs.utils;

import java.util.Arrays;
import java.util.List;
import java.util.Random;

public class GenerateEnglish {
    public  static final List<String> commonItems = Arrays.asList(
            "apple", "banana", "orange", "grape", "strawberry", "watermelon", "pineapple", "blueberry", "kiwi", "mango",
            "carrot", "broccoli", "potato", "tomato", "lettuce", "cucumber", "onion", "garlic", "pepper", "spinach",
            "milk", "bread", "cheese", "butter", "egg", "yogurt", "cereal", "juice", "coffee", "tea",
            "chicken", "beef", "fish", "pork", "shrimp", "salmon", "bacon", "sausage", "ham", "turkey",
            "rice", "pasta", "noodle", "quinoa", "couscous", "potato chips", "popcorn", "pretzel", "crackers", "cookies",
            "water", "soda", "juice", "lemonade", "iced tea", "coffee", "hot chocolate", "smoothie", "milkshake", "beer",
            "computer", "phone", "tablet", "laptop", "keyboard", "mouse", "headphones", "speaker", "camera", "printer",
            "desk", "chair", "lamp", "bookshelf", "couch", "table", "mirror", "clock", "rug", "painting",
            "shoes", "socks", "hat", "scarf", "gloves", "sunglasses", "umbrella", "purse", "backpack", "wallet",
            "shirt", "pants", "dress", "jacket", "sweater", "shorts", "skirt", "tie", "belt", "socks",
            "toothbrush", "toothpaste", "floss", "shampoo", "conditioner", "soap", "lotion", "razor", "towel", "hairdryer",
            "bed", "pillow", "blanket", "sheets", "alarm clock", "nightstand", "dresser", "mirror", "lamp", "candles",
            "car", "bicycle", "motorcycle", "bus", "train", "subway", "airplane", "boat", "truck", "scooter",
            "house", "apartment", "condo", "cabin", "tent", "hotel", "motel", "hostel", "villa", "cottage",
            "tree", "flower", "grass", "bush", "rock", "sand", "river", "lake", "ocean", "mountain",
            "sun", "moon", "star", "cloud", "rain", "snow", "wind", "thunder", "lightning", "rainbow",
            "cat", "dog", "bird", "fish", "rabbit", "hamster", "turtle", "snake", "lizard", "horse",
            "school", "classroom", "teacher", "student", "book", "pen", "pencil", "notebook", "backpack", "desk",
            "computer", "library", "laboratory", "cafeteria", "gym", "playground", "bus", "homework", "test", "exam",
            "doctor", "nurse", "hospital", "clinic", "medicine", "appointment", "surgery", "emergency", "health", "wellness",
            "movie", "music", "art", "dance", "theater", "book", "poem", "song", "painting", "sculpture",
            "friend", "family", "neighbor", "colleague", "classmate", "teacher", "boss", "employee", "customer", "stranger",
            "love", "happiness", "success", "failure", "hope", "dream", "fear", "anger", "joy", "sadness"
    );

    public static final List<String> commonNames = Arrays.asList(
            "James", "Olivia", "Liam", "Emma", "Noah", "Ava", "Isabella", "Sophia", "Ethan", "Jackson",
            "Lucas", "Aiden", "Mia", "Oliver", "Caden", "Charlotte", "Harper", "Elijah", "Amelia", "Abigail",
            "Benjamin", "Emily", "Henry", "Ella", "Alexander", "Sebastian", "Avery", "Scarlett", "Daniel", "Madison",
            "Matthew", "Aria", "Joseph", "Grace", "David", "Chloe", "Samuel", "Camila", "Carter", "Aubrey",
            "Wyatt", "Zoe", "Jayden", "Penelope", "John", "Lily", "Owen", "Lillian", "Dylan", "Natalie",
            "Luke", "Hannah", "Gabriel", "Eleanor", "Anthony", "Addison", "Isaac", "Stella", "Andrew", "Bella",
            "Christopher", "Nora", "Joseph", "Maya", "Christian", "Elena", "Jackson", "Paisley", "David", "Brooklyn",
            "Mateo", "Savannah", "Jack", "Victoria", "Leo", "Scarlet", "Ryan", "Claire", "Emma", "Elizabeth",
            "Lucas", "Evelyn", "Logan", "Aurora", "Caleb", "Abby", "Nathan", "Penelope", "Isaiah", "Anna",
            "Oscar", "Katherine", "Levi", "Ruby", "Mason", "Alice", "Eli", "Nora", "Julian", "Sadie",
            "Aiden", "Grace", "Gabriel", "Zoey", "Nicholas", "Adeline", "Jaxon", "Hazel", "Jordan", "Luna",
            "Dominic", "Riley", "Cameron", "Eva", "Tyler", "Emilia", "Angel", "Violet", "Isaiah", "Arya",
            "Isaiah", "Scarlett", "Isaiah", "Zoe", "Isaiah", "Madeline", "Isaiah", "Layla", "Isaiah", "Aaliyah",
            "Isaiah", "Ellie", "Isaiah", "Lily", "Isaiah", "Ariana", "Isaiah", "Skylar", "Isaiah", "Natalie",
            "Isaiah", "Nova", "Isaiah", "Julia", "Isaiah", "Hudson", "Claire", "Bryson", "Leah", "Hunter", "Mackenzie",
            "Miles", "Kylie", "Evan", "Aubree", "Adam", "Georgia"
    );

    public static final List<String> commonAdjectives = Arrays.asList(
            "happy", "sad", "big", "small", "bright", "dark", "fast", "slow", "hot", "cold",
            "loud", "quiet", "high", "low", "strong", "weak", "rich", "poor", "old", "young",
            "beautiful", "ugly", "clean", "dirty", "hard", "soft", "thick", "thin", "deep", "shallow",
            "long", "short", "expensive", "cheap", "safe", "dangerous", "happy", "sad", "empty", "full",
            "wide", "narrow", "tall", "short", "good", "bad", "easy", "difficult", "light", "heavy",
            "serious", "funny", "simple", "complex", "friendly", "unfriendly", "warm", "cool", "brave", "fearful",
            "calm", "nervous", "kind", "cruel", "honest", "dishonest", "generous", "selfish", "patient", "impatient",
            "polite", "rude", "efficient", "inefficient", "bright", "dull", "flexible", "rigid", "efficient", "inefficient",
            "modern", "ancient", "up-to-date", "outdated", "careful", "careless", "healthy", "unhealthy", "lively", "lifeless",
            "productive", "unproductive", "innovative", "conventional", "happy", "miserable", "successful", "unsuccessful", "positive", "negative",
            "creative", "destructive", "helpful", "unhelpful", "courageous", "cowardly", "confident", "doubtful", "efficient", "inefficient",
            "organized", "disorganized", "flexible", "inflexible", "graceful", "awkward", "diligent", "lazy", "energetic", "lethargic",
            "responsible", "irresponsible", "efficient", "inefficient", "innocent", "guilty", "wise", "foolish", "modest", "arrogant",
            "patient", "impatient", "humble", "proud", "sincere", "insincere", "logical", "illogical", "rational", "irrational",
            "charming", "repulsive", "sociable", "unsociable", "innocent", "guilty", "innovative", "conventional", "grateful", "ungrateful",
            "loyal", "disloyal", "tolerant", "intolerant", "sensitive", "insensitive", "generous", "stingy", "flexible", "inflexible",
            "mature", "immature", "adventurous", "cautious", "selfless", "selfish", "caring", "indifferent", "efficient", "inefficient",
            "articulate", "inarticulate", "resourceful", "unresourceful", "amiable", "unamiable", "ethical", "unethical", "confident", "doubtful"
    );

    static public String generate_adjectives(){
        Random random = new Random();
        int randomIndex = random.nextInt(commonAdjectives.size());
        return commonAdjectives.get(randomIndex);
    }

    static public String generate_thing(){
        Random random = new Random();
        int randomIndex = random.nextInt(commonItems.size());
        return commonItems.get(randomIndex);
    }

    static public String generate_half_name(){
        Random random = new Random();
        int randomIndex = random.nextInt(commonNames.size());
        return commonNames.get(randomIndex);
    }

    static public String generate_name(){
        String name_1 = generate_half_name();
        String name_2 = generate_half_name();
        while (name_2.equals(name_1)){
            name_2 = generate_half_name();
        }
        return name_1+" "+name_2;
    }

    static public String generate_adjectives_thing(){
        return generate_adjectives()+" "+generate_thing();
    }
}
