package def

import "github.com/WindomZ/go-dice/dice"

var _dice dice.DiceInt = *dice.NewDiceInt(10, dice.TYPE_POLL)

func tv() int {
	return _dice.TV()
}
