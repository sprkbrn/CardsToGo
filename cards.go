package main

import "fmt"

type suite struct {
	name   string
	symbol string
}

var suites = []suite{
	suite{"diamonds", "♦"},
	suite{"spades", "♠"},
	suite{"hearts", "♥"},
	suite{"clubs", "♣"},
}

type card struct {
	value int
	face  string
	group suite
}

var cards = []card{
	card{0, "A", suites[0]},
	card{1, "1", suites[0]},
	card{2, "2", suites[0]},
	card{3, "3", suites[0]},
	card{4, "4", suites[0]},
	card{5, "5", suites[0]},
	card{6, "6", suites[0]},
	card{7, "7", suites[0]},
	card{8, "8", suites[0]},
	card{9, "9", suites[0]},
	card{10, "10", suites[0]},
	card{11, "J", suites[0]},
	card{12, "Q", suites[0]},
	card{13, "K", suites[0]},
	card{0, "A", suites[1]},
	card{1, "1", suites[1]},
	card{2, "2", suites[1]},
	card{3, "3", suites[1]},
	card{4, "4", suites[1]},
	card{5, "5", suites[1]},
	card{6, "6", suites[1]},
	card{7, "7", suites[1]},
	card{8, "8", suites[1]},
	card{9, "9", suites[1]},
	card{10, "10", suites[1]},
	card{11, "J", suites[1]},
	card{12, "Q", suites[1]},
	card{13, "K", suites[1]},
	card{0, "A", suites[2]},
	card{1, "1", suites[2]},
	card{2, "2", suites[2]},
	card{3, "3", suites[2]},
	card{4, "4", suites[2]},
	card{5, "5", suites[2]},
	card{6, "6", suites[2]},
	card{7, "7", suites[2]},
	card{8, "8", suites[2]},
	card{9, "9", suites[2]},
	card{10, "10", suites[2]},
	card{11, "J", suites[2]},
	card{12, "Q", suites[2]},
	card{13, "K", suites[2]},
	card{0, "A", suites[3]},
	card{1, "1", suites[3]},
	card{2, "2", suites[3]},
	card{3, "3", suites[3]},
	card{4, "4", suites[3]},
	card{5, "5", suites[3]},
	card{6, "6", suites[3]},
	card{7, "7", suites[3]},
	card{8, "8", suites[3]},
	card{9, "9", suites[3]},
	card{10, "10", suites[3]},
	card{11, "J", suites[3]},
	card{12, "Q", suites[3]},
	card{13, "K", suites[3]},
}

func main() {
	for _, s := range cards {
		fmt.Println(s.face, s.group.symbol)
	}
}
