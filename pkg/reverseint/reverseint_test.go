package reverseint

import "testing"

type args struct {
	nums []int
	num  int
}

type cases []struct {
	name string
	args args
	want int
}

func TestReverse(t *testing.T) {

	tests := cases{
		{"test 1", args{num: 0}, 0},
		{"test 2", args{num: 123}, 321},
		{"test 3", args{num: 4566}, 6654},
		{"test 4", args{num: -321}, -123},
		{"test 5", args{num: 1534236469}, 0},
		{"test 6", args{num: -1534236469}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.num); got != tt.want {
				t.Errorf("GOT: %v, WANT: %v", got, tt.want)
			}
		})
	}
}
