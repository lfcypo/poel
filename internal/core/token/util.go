package token

func AddConstantString(s string) (index int) {
	constant := TokTree.ConstantList
	if constant == nil {
		constant = make([]Constant, 1)
	}
	for i, c := range constant {
		if c.Val == s {
			return i
		}
	}
	constant = append(constant, Constant{Val: s, Type: "String"})
	TokTree.ConstantList = constant
	return len(constant) - 1
}
