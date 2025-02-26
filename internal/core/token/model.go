package token

type Constant struct {
	Val  interface{}
	Type string
}

type Branch struct {
	Left  *Branch
	Right *Branch
	Op    string

	ChildBranches *[]Branch
}

type Tree struct {
	ConstantList []Constant

	Branches []Branch
}

var TokTree Tree
