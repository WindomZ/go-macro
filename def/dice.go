package def

import "github.com/WindomZ/go-dice/dice"

var _dice10 dice.DiceInt = *dice.NewDiceInt(10, dice.TYPE_POLL)
var _dice100 dice.DiceInt = *dice.NewDiceInt(100, dice.TYPE_POLL)

func tv10() int {
	return _dice10.TV()
}

func tv100() int {
	return _dice100.TV()
}
