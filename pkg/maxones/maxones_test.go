package maxones

import "testing"

type args struct {
	nums []int
}

type cases []struct {
	name string
	args args
	want int
}

func TestFindMaxConsecutiveOnes(t *testing.T) {

	tests := cases{
		{"test 1", args{nums: []int{0}}, 0},
		{"test 2", args{nums: []int{1}}, 1},
		{"test 3", args{nums: []int{0, 1, 1, 1, 0, 0, 1}}, 3},
		{"test 4", args{nums: []int{0, 1, 0, 1, 0, 0, 1}}, 1},
		{"test 5", args{nums: []int{0, 1, 1, 1, 1, 0, 1}}, 4},
		{"test 6", args{nums: []int{0, 1, 1, 0, 0, 0, 1}}, 2},
		{"test 7", args{nums: []int{0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1}}, 9},
		{"test 8", args{nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0}}, 9},
		{"test 9", args{nums: []int{0, 1, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}}, 2},
		{"test 10", args{nums: []int{0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0}}, 9},
		{"test 11", args{nums: []int{}}, 0},
		{"test 12", args{nums: []int{0, 0}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMaxConsecutiveOnes(tt.args.nums); got != tt.want {
				t.Errorf("GOT: %v, WANT %v", got, tt.want)
			}
		})
	}
}
