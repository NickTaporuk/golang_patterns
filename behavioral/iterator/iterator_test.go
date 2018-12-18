package iterator

import (
	"reflect"
	"testing"
)

func Test_intStatefulIterator_Value(t *testing.T) {
	type fields struct {
		current int
		data    []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "intStatefulIterator_Value",
			fields: fields{current: 0, data: []int{1}},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			it := &intStatefulIterator{
				current: tt.fields.current,
				data:    tt.fields.data,
			}
			if got := it.Value(); got != tt.want {
				t.Errorf("intStatefulIterator.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intStatefulIterator_Next(t *testing.T) {
	type fields struct {
		current int
		data    []int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:"intStatefulIterator_Next",
			fields:fields{current:0,data:[]int{1}},
			want:false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			it := &intStatefulIterator{
				current: tt.fields.current,
				data:    tt.fields.data,
			}
			if got := it.Next(); got != tt.want {
				t.Errorf("intStatefulIterator.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewIntStatefulIterator(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want *intStatefulIterator
	}{
		{
			name:"NewIntStatefulIterator",
			args:args{data:[]int{1,2,3,4,5,6,7}},
			want:NewIntStatefulIterator([]int{1,2,3,4,5,6,7}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIntStatefulIterator(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIntStatefulIterator() = %v, want %v", got, tt.want)
			}
		})
	}
}
