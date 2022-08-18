package factory

import (
	"reflect"
	"testing"
)

func TestAbsFactory_GetShape(t *testing.T) {
	type args struct {
		t shapeType
	}
	tests := []struct {
		name string
		args args
		want Shape
	}{
		{
			"rectangle",
			args{
				t: "rectangle",
			},
			&Rectangle{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			factory := &AbsFactory{}
			if got := factory.GetShape(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetShape() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbsFactory_GetColor(t *testing.T) {
	type args struct {
		t colorType
	}
	tests := []struct {
		name string
		args args
		want Color
	}{
		{
			"blue",
			args{
				t: "blue",
			},
			&Blue{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			factory := &AbsFactory{}
			if got := factory.GetColor(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
