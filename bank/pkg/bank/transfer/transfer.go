package transfer

func Total(amount int) (bonus1 int) {
	bonus := 5
	bonus0 := (amount * bonus) / 1000
	bonus1 = amount + bonus0
	return bonus1
}
