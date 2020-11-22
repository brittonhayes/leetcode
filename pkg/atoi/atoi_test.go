package atoi

import "testing"

func TestStringToInt(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test 1", args{"-91283472332"}, -2147483648},
		{"test 2", args{"4193 with words"}, 4193},
		{"test 3", args{"   -42"}, -42},
		{"test 4", args{"words and 987"}, 0},
		{"test 5", args{"42"}, 42},
		{"test 6", args{"+-12"}, 0},
		{"test 7", args{"-+12"}, 0},
		{"test 8", args{"21474836460"}, 2147483647},
		{"test 9", args{"20000000000000000000"}, 2147483647},
		{"test 10", args{"+"}, 0},
		{"test 11", args{"-   234"}, 0},
		{"test 12", args{" ++1"}, 0},
		{"test 13", args{" --1"}, 0},
		{"test 14", args{" + 413"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToInt(tt.args.v); got != tt.want {
				t.Errorf("GOT: %v, WANT: %v", got, tt.want)
			}
		})
	}
}
