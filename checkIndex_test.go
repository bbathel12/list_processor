package main

import "testing"

func Test_checkIndex(t *testing.T) {
	type args struct {
		recs           *int
		newRecs        *int
		dupes          *int
		index          *ind
		hashedLineChan *chan string
		newHashChan    *chan string
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkIndex(tt.args.recs, tt.args.newRecs, tt.args.dupes, tt.args.index, tt.args.hashedLineChan, tt.args.newHashChan)
		})
	}
}
