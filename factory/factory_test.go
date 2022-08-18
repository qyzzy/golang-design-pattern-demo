package factory

import (
	"reflect"
	"testing"
)

func TestFactory_GetShape(t *testing.T) {
	type args struct {
		t shapeType
	}
	tests := []struct {
		name string
		args args
		want Shape
	}{
		{
			name: "circle",
			args: args{
				t: "circle",
			},
			want: &Circle{},
		},
		{
			"square",
			args{
				t: "square",
			},
			&Square{},
		},
		{
			"nil",
			args{
				t: "",
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			factory := &Factory{}
			if got := factory.GetShape(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetShape() = %v, want %v", got, tt.want)
			}
		})
	}
}
