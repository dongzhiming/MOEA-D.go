package moead

// 定义枚举
type (
	NeighborType int
	FunctionType int
)

// iota 初始化后会自动递增
const (
	NEIGHBOR   NeighborType = iota // value --> 0
	POPULATION                     // value --> 1
)

const (
	TCHE FunctionType = iota
	PBI
	EWC
	WCP
	ATCH
	AASF
	ASF
)
