package builder

type Builder interface {
	BuildPartA()
	BuildPartB()
}

type Director struct {
	builder Builder
}

func (d *Director) SetBuilder(builder Builder) {
	d.builder = builder
}

func (d *Director) Construct() {
	d.builder.BuildPartA()
	d.builder.BuildPartB()
}

type ConcreteBuilder struct {
	code string
}

func (b *ConcreteBuilder) BuildPartA() {
	b.code = "QAQ"
}

func (b *ConcreteBuilder) BuildPartB() {
	b.code += "qyzzy"
}

func (b *ConcreteBuilder) GetCode() string {
	return b.code
}
