package reverse

import (
	"reflect"
	"testing"
)

func Test_reverseListV1(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "head = [1,2,3,4,5]",
			args: args{head: CreateList(&[]int{1, 2, 3, 4, 5})},
			want: CreateList(&[]int{5, 4, 3, 2, 1}),
		},
		{
			name: "head = [1,2]",
			args: args{head: CreateList(&[]int{1, 2})},
			want: CreateList(&[]int{2, 1}),
		},
		{
			name: "head = []",
			args: args{head: CreateList(&[]int{})},
			want: CreateList(&[]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseListV1(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseListV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
