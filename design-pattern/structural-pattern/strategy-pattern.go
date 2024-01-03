package structural_pattern

//定义一个策略类
type IStrategy interface {
	do(int64 ,int64) int64
}

// 策略实现：加
type add struct{}

func (ad add) do(a int64, b int64) int64 {
	return a + b
}

// 策略实现：减
type reduce struct{}
func (*reduce) do(a, b int64) int64 { return a - b}

// 具体策略的执行者
type Operator struct {
	strategy IStrategy
}

func (operator *Operator) setStrategy(strategy IStrategy) {
	operator.strategy = strategy
}

// 调用策略中的方法
func (operator *Operator) calculate(a, b int64) int64 {
	return operator.strategy.do(a, b)
}
