package poker

import (
	"fmt"
	"sort"
)

type Card [2]byte

type Suit byte

const (
	Spades Suit = iota
	Hearts
	Diamonds
	Clubs
)

func (s Suit) String() string {
	switch s {
	case Spades:
		return "S"
	case Hearts:
		return "H"
	case Diamonds:
		return "D"
	case Clubs:
		return "C"
	default:
		panic("invalid Suit")
	}
}

func (c Card) Suit() Suit {
	switch c[1] {
	case 'S':
		return Spades
	case 'H':
		return Hearts
	case 'D':
		return Diamonds
	case 'C':
		return Clubs
	default:
		panic("invalid Card suit")
	}
}

type Rank int

const (
	Jack  Rank = 11
	Queen Rank = 12
	King  Rank = 13
	Ace   Rank = 14
)

func (r Rank) String() string {
	switch r {
	case Ace:
		return "A"
	case King:
		return "K"
	case Queen:
		return "Q"
	case Jack:
		return "J"
	case 10:
		return "T"
	default:
		if r < 2 || r > 9 {
			panic("invalid Rank")
		}
		return string('0' + byte(r))
	}
}

func (c Card) Rank() Rank {
	switch c[0] {
	case '2', '3', '4', '5', '6', '7', '8', '9':
		return Rank(c[0] - '0')
	case 'T':
		return 10
	case 'J':
		return Jack
	case 'Q':
		return Queen
	case 'K':
		return King
	case 'A':
		return Ace
	default:
		panic("invalid card rank")
	}
}

func (c Card) String() string {
	return c.Rank().String() + c.Suit().String()
}

type Hand [5]Card

func (h Hand) String() string {
	return fmt.Sprint(h[:])
}

func ParseHand(s string) Hand {
	// no validation
	h := Hand{}
	for i := 0; i < 5; i++ {
		copy(h[i][:], s[3*i:3*i+2])
	}
	return h
}

func seq(ranks []Rank) bool {
	for i, r := range ranks[1:] {
		if r != ranks[0]-Rank(i)-1 {
			return false
		}
	}
	return true
}

type ranks []Rank

func (r ranks) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r ranks) Len() int {
	return len(r)
}

func (r ranks) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (h Hand) Evaluate() Outcome {
	out := Outcome{}
	r := make(ranks, 5)
	for i, c := range h {
		r[i] = c.Rank()
	}
	sort.Sort(sort.Reverse(r))
	same := [4]bool{}
	for i := 0; i < 4; i++ {
		same[i] = r[i] == r[i+1]
	}
	// Four of a Kind: Four cards of the same value.
	if same[0] && same[1] && same[2] {
		out.Category = FourOfAKind
		out.Rank = r[0:5]
		return out
	}
	if same[1] && same[2] && same[3] {
		out.Category = FourOfAKind
		out.Rank = append(r[1:5], r[0])
		return out
	}
	// Full House: Three of a kind and a pair.
	if same[0] && same[1] && same[3] {
		out.Category = FullHouse
		out.Rank = r[0:5]
		return out
	}
	if same[0] && same[2] && same[3] {
		out.Category = FullHouse
		out.Rank = append(r[2:5], r[0:2]...)
		return out
	}
	haveFlush := true
	for _, c := range h[1:] {
		if c.Suit() != h[0].Suit() {
			haveFlush = false
			break
		}
	}
	// Straight: All cards are consecutive values.
	// ace can be the low or high card in a straight
	if r[0] == Ace && r[4] == Rank(2) && seq(r[1:5]) {
		out.Category = Straight
		if haveFlush {
			out.Category = StraightFlush
		}
		out.Rank = append(r[1:5], r[0])
		return out
	}
	if seq(r[:]) {
		out.Category = Straight
		if haveFlush {
			out.Category = StraightFlush
			if Rank(r[0]) == Ace {
				out.Category = RoyalFlush
			}
		}
		out.Rank = r[0:5]
		return out
	}
	// Flush: All cards of the same suit.
	if haveFlush {
		out.Category = Flush
		out.Rank = r[0:5]
		return out
	}
	// Three of a Kind: Three cards of the same value.
	if same[0] && same[1] {
		out.Category = ThreeOfAKind
		out.Rank = r[0:5]
		return out
	}
	if same[1] && same[2] {
		out.Category = ThreeOfAKind
		out.Rank = append(r[1:4], r[0], r[4])
		return out
	}
	if same[2] && same[3] {
		out.Category = ThreeOfAKind
		out.Rank = append(r[2:5], r[0], r[1])
		return out
	}
	if same[0] && same[2] {
		out.Category = TwoPairs
		out.Rank = r[0:5]
		return out
	}
	if same[0] && same[3] {
		out.Category = TwoPairs
		out.Rank = append(r[0:2], r[3], r[4], r[2])
		return out
	}
	if same[1] && same[3] {
		out.Category = TwoPairs
		out.Rank = append(r[1:5], r[0])
		return out
	}
	if same[0] {
		out.Category = OnePair
		out.Rank = r[0:5]
		return out
	}
	if same[1] {
		out.Category = OnePair
		out.Rank = append(r[1:3], r[0], r[3], r[4])
		return out
	}
	if same[2] {
		out.Category = OnePair
		out.Rank = append(r[2:4], r[0], r[1], r[4])
		return out
	}
	if same[3] {
		out.Category = OnePair
		out.Rank = append(r[3:5], r[0], r[1], r[2])
		return out
	}
	out.Category = HighCard
	out.Rank = r[0:5]
	return out
}

type Category int

const (
	HighCard Category = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

func (c Category) String() string {
	switch c {
	case HighCard:
		return "high card"
	case OnePair:
		return "one pair"
	case TwoPairs:
		return "two pairs"
	case ThreeOfAKind:
		return "three of a kind"
	case Straight:
		return "straight"
	case Flush:
		return "flush"
	case FullHouse:
		return "full house"
	case FourOfAKind:
		return "four of a kind"
	case StraightFlush:
		return "straight flush"
	case RoyalFlush:
		return "royal flush"
	default:
		panic("invalid Category")
	}
}

type Outcome struct {
	Category Category
	Rank     []Rank
}

func (o Outcome) String() string {
	return fmt.Sprintf("%s %s", o.Category, o.Rank)
}

func Less(i, j Outcome) bool {
	if i.Category < j.Category {
		return true
	}
	if i.Category > j.Category {
		return false
	}
	for k := 0; k < 5; k++ {
		if i.Rank[k] < j.Rank[k] {
			return true
		}
		if i.Rank[k] > j.Rank[k] {
			return false
		}
	}
	// they are equal
	return false
}
