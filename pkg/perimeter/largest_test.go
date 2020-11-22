package perimeter

import "testing"

func TestLargestPerimeter(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test 1", args{nums: []int{1, 2, 1}}, 0},
		{"test 2", args{nums: []int{3, 2, 3, 4}}, 10},
		{"test 3", args{nums: []int{2, 1, 2}}, 5},
		{"test 4", args{nums: []int{3, 6, 2, 3}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestPerimeter(tt.args.nums); got != tt.want {
				t.Errorf("GOT: %v, WANT: %v", got, tt.want)
			}
		})
	}
}
