package deck

//card represents a playing card
type card struct {
	Face  string //the face value of the card
	Court suite  //the group the card is in
}

//suite represents the group a card is in
type suite struct {
	Name   string //the name of the suite
	Symbol string //symbol of representation
}

//all possible faces a card can have
var cardFaces = []string{"A", "1", "2", "3", "4", "5",
	"6", "7", "8", "9", "10", "J", "Q", "K"}

//CardGroups is a collection of our suites
var CardGroups = []suite{Diamonds, Spades, Hearts, Clubs}

//One of the four prominent suites in playing cards
var (
	Diamonds = suite{
		"Diamonds",
		"♦",
	}
	Spades = suite{
		"Spades",
		"♠",
	}
	Hearts = suite{
		"Hearts",
		"♥",
	}
	Clubs = suite{
		"Clubs",
		"♣",
	}
)

//DeckOfCards represents all possible playing cards
var DeckOfCards = func() []card {
	var result []card
	for _, grp := range CardGroups {
		for _, face := range cardFaces {
			result = append(result, card{face, grp})
		}
	}
	return result
}()
