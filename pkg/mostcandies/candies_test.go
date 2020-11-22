package mostcandies

import (
	"reflect"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	type args struct {
		candies      []int
		extraCandies int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{"test 1", args{
			candies:      []int{4, 2, 1, 1, 2},
			extraCandies: 1,
		}, []bool{true, false, false, false, false},
		},
		{"test 2", args{
			candies:      []int{2, 3, 5, 1, 3},
			extraCandies: 3,
		}, []bool{true, true, true, false, true},
		},
		{"test 3", args{
			candies:      []int{12, 1, 12},
			extraCandies: 10,
		}, []bool{true, false, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KidsWithCandies(tt.args.candies, tt.args.extraCandies); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
