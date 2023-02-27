package binary_tree

import (
	"reflect"
	"testing"
)

type args226 struct {
	root *TreeNode
}

var tests226 = []struct {
	name string
	args args226
	want []int
}{
	{
		name: "test1",
		args: args226{
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 9,
					},
				},
			},
		},
		want: []int{4, 7, 2, 9, 6, 3, 1},
	},
	{
		name: "test2",
		args: args226{
			root: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		want: []int{2, 3, 1},
	},
}

func Test_invertTreeV1(t *testing.T) {
	for _, tt := range tests226 {
		t.Run(tt.name, func(t *testing.T) {
			got := invertTreeV1(tt.args.root)
			if gotList := levelTraversal(got); !reflect.DeepEqual(gotList, tt.want) {
				t.Errorf("invertTreeV1() = %v, want %v", gotList, tt.want)
			}
		})
	}
}

func Test_invertTreeV2(t *testing.T) {
	for _, tt := range tests226 {
		t.Run(tt.name, func(t *testing.T) {
			got := invertTreeV2(tt.args.root)
			if gotList := levelTraversal(got); !reflect.DeepEqual(gotList, tt.want) {
				t.Errorf("invertTreeV1() = %v, want %v", gotList, tt.want)
			}
		})
	}
}
