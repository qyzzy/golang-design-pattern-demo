package decorator

type CodeInterface interface {
	Generate() string
}

type BaseCode struct{}

func (c *BaseCode) Generate() string {
	return "partA"
}

type CodeDecorator struct {
	baseCode CodeInterface
	partB    string
}

func NewCodeDecorator(baseCode CodeInterface, partB string) *CodeDecorator {
	return &CodeDecorator{
		baseCode: baseCode,
		partB:    partB,
	}
}

func (c *CodeDecorator) Generate() string {
	return c.baseCode.Generate() + c.partB
}
