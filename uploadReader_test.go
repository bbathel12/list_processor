package main

import "testing"

func Test_readUpload(t *testing.T) {
	type args struct {
		uploadName string
		lineChan   *chan string
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			readUpload(tt.args.uploadName, tt.args.lineChan)
		})
	}
}
