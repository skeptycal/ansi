// Copyright (c) 2020 Michael Treanor
// MIT License

package ansi

import (
	"reflect"
	"testing"
)

var tests = []struct {
	name string
	b    []byte
	want string
}{
	{"no bytes", []byte{}, "\x1b[0;38;5;7;48;5;0m"},
	{"1 byte (Green)", []byte{'2'}, "\x1b[38;5;2m"},
	{"green bold text", []byte{'2', '0', '1'}, "\x1b[1;38;5;2;48;5;0m"},
	{"black on red", []byte{'0', '1', '0'}, "\x1b[0;38;5;0;48;5;1m"},
	{"blue on red italic", []byte{'4', '1', '3'}, "\x1b[3;38;5;4;48;5;1m"},
	// {"zero byte", []byte{}, []byte{27, 91, 32, 109}},
	// {"one byte", []byte{32}, []byte{27, 91, 32, 109}},
	// {"three bytes", []byte{32}, []byte{27, 91, 32, 109}},
}

func Test_ansiBytes(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := string(ansiBytesBB3d(tt.b...))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ansiBytesBB3d(%s) = %v, want %v", tt.name, got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			got := string(ansiBytesBB2d(tt.b...))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ansiBytesBB3d(%s) = %v, want %v", tt.name, got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			got := string(ansiBytesSB3d(tt.b...))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ansiBytesBB3d(%s) = %v, want %v", tt.name, got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			got := string(ansiBytesSB2d(tt.b...))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ansiBytesBB3d(%s) = %v, want %v", tt.name, got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			got := string(ansiAppendString(tt.b...))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ansiBytesBB3d(%s) = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestExampleAnsi(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"example", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ExampleAnsi(); (err != nil) != tt.wantErr {
				t.Errorf("ExampleAnsi() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
