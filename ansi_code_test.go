// Copyright (c) 2020 Michael Treanor
// MIT License

package ansi

import (
	"fmt"
	"reflect"
	"testing"
)

// benchmark results
/*
/ Initial
========================
BenchmarkAnsiBytes/ansiBytes1-8         	15102102	        72.49 ns/op	       0 B/op	       0 allocs/op
/// returning string([]bytes{})
BenchmarkAnsiBytes/ansiBytes2-8         	 5918802	       214.0 ns/op	      88 B/op	       5 allocs/op

/ Compare strings.Builder to bytes.Buffer
BenchmarkAnsiBytes/default_2+(empty)-8         	25913701	        43.22 ns/op	      24 B/op	       1 allocs/op
BenchmarkAnsiBytes/default_2+(empty)#01-8      	29575016	        40.39 ns/op	      16 B/op	       1 allocs/op

/ length 1 and 3 are the most common cases by far ...
BenchmarkAnsiBytes/default_2+(1_byte)-8        	17394726	        70.57 ns/op	      24 B/op	       1 allocs/op
BenchmarkAnsiBytes/default_2+(1_byte)#01-8     	 9911772	       120.8 ns/op	      40 B/op	       3 allocs/op

BenchmarkAnsiBytes/default_2+(2_bytes)-8       	12865417	        92.73 ns/op	      24 B/op	       1 allocs/op
BenchmarkAnsiBytes/default_2+(2_bytes)#01-8    	 9142824	       133.5 ns/op	      40 B/op	       3 allocs/op

BenchmarkAnsiBytes/default_2+(3_bytes)-8       	 6174462	       188.6 ns/op	      56 B/op	       3 allocs/op
BenchmarkAnsiBytes/default_2+(3_bytes)#01-8    	 4515218	       265.3 ns/op	     104 B/op	       6 allocs/op

BenchmarkAnsiBytes/default_2+(4_bytes)-8       	 4898269	       245.8 ns/op	      75 B/op	       3 allocs/op
BenchmarkAnsiBytes/default_2+(4_bytes)#01-8    	 3839222	       312.8 ns/op	     128 B/op	       6 allocs/op

BenchmarkAnsiBytes/default_2+(many_bytes)-8    	 1000000	      1071 ns/op	    1000 B/op	       3 allocs/op
BenchmarkAnsiBytes/default_2+(many_bytes)#01-8 	 1011006	      1194 ns/op	    1240 B/op	       6 allocs/op

/ added 0 len check (returns default with no allocation)
	/ original (with buffer reset and case 0)
	BenchmarkAnsiBytes/default_2+(empty)-8         			24936866	         44.27 ns/op	      24 B/op	       1 allocs/op
	BenchmarkAnsiBytes/default_2+(empty)#01-8      			29483440	         41.17 ns/op	      16 B/op	       1 allocs/op
	/ returning []byte constant
	BenchmarkAnsiBytes/bytes.Buffer_2+(empty)-8         	244844850	         4.779 ns/op	       0 B/op	       0 allocs/op
	BenchmarkAnsiBytes/bytes.Buffer_3+(empty)-8         	254841993	         4.651 ns/op	       0 B/op	       0 allocs/op
	/ returning string constant
	BenchmarkAnsiBytes/strings.Builder_2+(empty)-8         	249329725	         4.685 ns/op	       0 B/op	       0 allocs/op
	BenchmarkAnsiBytes/strings.Builder_3+(empty)-8         	260941171	         4.570 ns/op	       0 B/op	       0 allocs/op

/ length 1 and 3 are the most common cases by far ...
/ since the trend is clear, isolate those cases for further testing
BenchmarkAnsiBytes/default_2+(1_byte)-8        	16995898	        72.64 ns/op	      24 B/op	       1 allocs/op
BenchmarkAnsiBytes/default_2+(1_byte)#01-8     	 9734612	       124.1 ns/op	      40 B/op	       3 allocs/op
BenchmarkAnsiBytes/default_2+(3_bytes)-8       	 6276564	       189.2 ns/op	      56 B/op	       3 allocs/op
BenchmarkAnsiBytes/default_2+(3_bytes)#01-8    	 4495221	       274.5 ns/op	     104 B/op	       6 allocs/op

BenchmarkAnsiBytes/default_2+(empty)-8         	243858049	         4.647 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnsiBytes/default_3+(empty)-8         	258202844	         4.665 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnsiBytes/default_2+(1_byte)-8        	29871585	        37.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnsiBytes/default_3+(1_byte)-8        	33776346	        36.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnsiBytes/default_2+(3_bytes)-8       	 7239094	       164.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkAnsiBytes/default_3+(3_bytes)-8       	18027324	        67.50 ns/op	       0 B/op	       0 allocs/op

BenchmarkAnsiBytes/default_2+(empty)-8         	142359066	         8.656 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnsiBytes/default_3+(empty)-8         	129582454	         9.115 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnsiBytes/default_2+(1_byte)-8        	12934778	        86.40 ns/op	      24 B/op	       2 allocs/op
BenchmarkAnsiBytes/default_3+(1_byte)-8        	13888183	        86.26 ns/op	      24 B/op	       2 allocs/op
BenchmarkAnsiBytes/default_2+(3_bytes)-8       	 4988379	       236.6 ns/op	      88 B/op	       5 allocs/op
BenchmarkAnsiBytes/default_3+(3_bytes)-8       	 8414373	       143.5 ns/op	      56 B/op	       3 allocs/op

/ Compare default is length 2+ to 3+ (special algorithm for length 3)
========================

*/

func BenchmarkAnsiBytes(b *testing.B) {
	inputs := []struct {
		name  string
		input []byte
	}{
		{"empty", []byte{}},
		{"1 byte", []byte{1}},
		// {"2 bytes", []byte{1, 160}},
		// {"3 bytes", []byte{1, 160, 196}},
		// {"4 bytes", []byte{1, 3, 160, 196}},
		// {"many bytes", []byte{1, 3, 160, 196, 1, 3, 160, 196, 1, 3, 160, 196, 1, 3, 160, 196, 1, 3, 160, 196, 1, 3, 160, 196, 1, 3, 160, 196, 1, 3, 160, 196, 1, 3, 160, 196}},
	}
	benchmarks := []struct {
		name string
		fn   func(b ...byte) []byte
		// fn func(b ...byte) string
	}{
		{"bytes.Buffer 2+", ansiBytesBB2d},
		{"bytes.Buffer 3+", ansiBytesBB3d},
		{"Len1special", ansiOneAppend},

		// {"strings.Builder 2+", ansiBytesSB2d},
		// {"strings.Builder 3+", ansiBytesSB3d},
	}
	for _, in := range inputs {
		for _, bb := range benchmarks {
			name := fmt.Sprintf("%s(%s)", bb.name, in.name)
			b.Run(name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					bb.fn(in.input...)
				}
			})
		}
	}
}

func TestNewColor(t *testing.T) {
	type args struct {
		foregroundByte byte
		backgroundByte byte
		effectByte     byte
	}
	tests := []struct {
		name string
		args args
		want Ansi
	}{
		// TODO: Add test cases.
		{"", args{2, 0, 1}, NewColor(2, 0, 1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewColor(tt.args.foregroundByte, tt.args.backgroundByte, tt.args.effectByte); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_ansiBytes(t *testing.T) {
// 	type args struct {
// 		bytes []byte
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want []byte
// 	}{
// 		// TODO: Add test cases.
// 		{"nil", args{[]byte{}}, bReset},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := ansiBytes(tt.args.bytes...); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("ansiBytes() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
