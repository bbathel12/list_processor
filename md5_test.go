package main

import "testing"

func Test_forceMd5(t *testing.T) {
	type args struct {
		lineChan       *chan string
		hashedLineChan *chan string
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			forceMd5(tt.args.lineChan, tt.args.hashedLineChan)
		})
	}
}
