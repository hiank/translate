package tool



// UnicodeF used to check if content unicode byte
func UnicodeF(data rune) bool {

	return data > 127
}

// NumberNF used to check if the data is not a number
func NumberNF(data rune) bool {

	rlt := data < 48
	if !rlt {
		rlt = data > 57
	}

	return rlt
}
