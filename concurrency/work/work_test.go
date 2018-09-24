// Example provided with help from Jason Waldrip.
// Package work manages a pool of goroutines to perform work.
package work

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		maxGoroutines int
	}
	tests := []struct {
		name string
		args args
		want *Pool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.maxGoroutines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPool_Run(t *testing.T) {
	type args struct {
		w Worker
	}
	tests := []struct {
		name string
		p    *Pool
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Run(tt.args.w)
		})
	}
}

func TestPool_Shutdown(t *testing.T) {
	tests := []struct {
		name string
		p    *Pool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Shutdown()
		})
	}
}
