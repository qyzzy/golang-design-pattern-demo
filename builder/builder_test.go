package builder

import "testing"

func TestDirector_Construct(t *testing.T) {
	type fields struct {
		builder Builder
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			"test",
			fields{&ConcreteBuilder{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Director{
				builder: tt.fields.builder,
			}
			d.SetBuilder(tt.fields.builder)
		})
	}
}
