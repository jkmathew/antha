package wtype

// A block of instructions associated with a particular job, etc.
type BlockID struct {
	Value string
}

func NewBlockID(id string) BlockID {
	return BlockID{Value: id}
}

func (a BlockID) String() string {
	return a.Value
}
