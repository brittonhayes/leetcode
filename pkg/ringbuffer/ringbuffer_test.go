package ringbuffer_test

import (
	"fmt"
	"github.com/brittonhayes/leetcode/pkg/ringbuffer"
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want ringbuffer.MyCircularQueue
	}{
		{"Test 1", args{5}, ringbuffer.MyCircularQueue{
			Tasks: []int{0, 1, 2, 3, 4},
			Full:  false,
		}},
		{"Test 2", args{1}, ringbuffer.MyCircularQueue{
			Tasks: []int{0},
			Full:  false,
		}},
		{"Test 3", args{2}, ringbuffer.MyCircularQueue{
			Tasks: []int{0, 1},
			Full:  false,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ringbuffer.Constructor(tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularQueue_DeQueue(t1 *testing.T) {
	type fields struct {
		Tasks []int
		Full  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"Test 1", fields{Tasks: []int{1, 2, 3}}, true},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &ringbuffer.MyCircularQueue{
				Tasks: tt.fields.Tasks,
				Full:  tt.fields.Full,
			}
			fmt.Println(tt.fields.Tasks)
			if got := t.DeQueue(); got != tt.want {
				t1.Errorf("DeQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularQueue_EnQueue(t1 *testing.T) {
	type fields struct {
		Tasks []int
		Full  bool
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"Test 1", fields{Tasks: []int{1, 2, 3}}, args{5}, true},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &ringbuffer.MyCircularQueue{
				Tasks: tt.fields.Tasks,
				Full:  tt.fields.Full,
			}
			if got := t.EnQueue(tt.args.value); got != tt.want {
				t1.Errorf("EnQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularQueue_Front(t1 *testing.T) {
	type fields struct {
		Tasks []int
		Full  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &ringbuffer.MyCircularQueue{
				Tasks: tt.fields.Tasks,
				Full:  tt.fields.Full,
			}
			if got := t.Front(); got != tt.want {
				t1.Errorf("Front() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularQueue_IsEmpty(t1 *testing.T) {
	type fields struct {
		Tasks []int
		Full  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &ringbuffer.MyCircularQueue{
				Tasks: tt.fields.Tasks,
				Full:  tt.fields.Full,
			}
			if got := t.IsEmpty(); got != tt.want {
				t1.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularQueue_IsFull(t1 *testing.T) {
	type fields struct {
		Tasks []int
		Full  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &ringbuffer.MyCircularQueue{
				Tasks: tt.fields.Tasks,
				Full:  tt.fields.Full,
			}
			if got := t.IsFull(); got != tt.want {
				t1.Errorf("IsFull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularQueue_Rear(t1 *testing.T) {
	type fields struct {
		Tasks []int
		Full  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &ringbuffer.MyCircularQueue{
				Tasks: tt.fields.Tasks,
				Full:  tt.fields.Full,
			}
			if got := t.Rear(); got != tt.want {
				t1.Errorf("Rear() = %v, want %v", got, tt.want)
			}
		})
	}
}
