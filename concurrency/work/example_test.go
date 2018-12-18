package work

import "testing"

func Test_namePrinter_Task(t *testing.T) {
	tests := []struct {
		name string
		m    *namePrinter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Task()
		})
	}
}

func Test_example(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			example()
		})
	}
}
