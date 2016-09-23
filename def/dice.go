package def

import "github.com/WindomZ/go-dice/dice"

var _dice10 dice.DiceInt = *dice.NewDiceInt(10, dice.TYPE_POLL)
var _dice100 dice.DiceInt = *dice.NewDiceInt(100, dice.TYPE_POLL)
var _dice10000 dice.DiceInt = *dice.NewDiceInt(10000, dice.TYPE_POLL)
var _dice100000000 dice.DiceInt = *dice.NewDiceInt(100000000, dice.TYPE_POLL)

func tv10() int {
	return _dice10.TV()
}

func tv100() int {
	return _dice100.TV()
}

func tv10000() int {
	return _dice10000.TV()
}

func tv100000000() int {
	return _dice100000000.TV()
}
