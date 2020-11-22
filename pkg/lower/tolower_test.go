package lower

import "testing"

func TestToLowerCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test 1", args{str: "Hello"}, "hello" },
		{"test 2", args{str: "HeLLo"}, "hello" },
		{"test 3", args{str: "here"}, "here" },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToLowerCase(tt.args.str); got != tt.want {
				t.Errorf("ToLowerCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
