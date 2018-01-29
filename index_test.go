package main

import (
	"reflect"
	"testing"
)

func Test_newIndex(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name      string
		args      args
		wantIndex *ind
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIndex := newIndex(tt.args.name); !reflect.DeepEqual(gotIndex, tt.wantIndex) {
				t.Errorf("newIndex() = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}

func Test_ind_add(t *testing.T) {
	type fields struct {
		storage map[string][]string
		name    string
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &ind{
				storage: tt.fields.storage,
				name:    tt.fields.name,
			}
			index.add(tt.args.value)
		})
	}
}

func Test_ind_write(t *testing.T) {
	type fields struct {
		storage map[string][]string
		name    string
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &ind{
				storage: tt.fields.storage,
				name:    tt.fields.name,
			}
			index.write()
		})
	}
}

func Test_ind_open(t *testing.T) {
	type fields struct {
		storage map[string][]string
		name    string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &ind{
				storage: tt.fields.storage,
				name:    tt.fields.name,
			}
			if err := index.open(); (err != nil) != tt.wantErr {
				t.Errorf("ind.open() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_ind_contains(t *testing.T) {
	type fields struct {
		storage map[string][]string
		name    string
	}
	type args struct {
		line string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &ind{
				storage: tt.fields.storage,
				name:    tt.fields.name,
			}
			if got := index.contains(tt.args.line); got != tt.want {
				t.Errorf("ind.contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
