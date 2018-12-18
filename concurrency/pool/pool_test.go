// Example provided with help from Fatih Arslan and Gabriel Aszalos.
// Package pool manages a user defined set of resources.
package pool

import (
	"io"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		fn   func() (io.Closer, error)
		size uint
	}
	tests := []struct {
		name    string
		args    args
		want    *Pool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.fn, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPool_Acquire(t *testing.T) {
	tests := []struct {
		name    string
		p       *Pool
		want    io.Closer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.Acquire()
			if (err != nil) != tt.wantErr {
				t.Errorf("Pool.Acquire() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pool.Acquire() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPool_Release(t *testing.T) {
	type args struct {
		r io.Closer
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
			tt.p.Release(tt.args.r)
		})
	}
}

func TestPool_Close(t *testing.T) {
	tests := []struct {
		name string
		p    *Pool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Close()
		})
	}
}
