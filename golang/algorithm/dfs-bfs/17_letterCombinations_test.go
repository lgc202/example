package dfs_bfs

import (
	"reflect"
	"testing"
)

func Test_letterCombinationsDFS(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{digits: "23"},
			want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name: "test2",
			args: args{digits: ""},
			want: nil,
		},
		{
			name: "test3",
			args: args{"2"},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterCombinationsDFS(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCombinationsDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_letterCombinationsBFS(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{digits: "23"},
			want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name: "test2",
			args: args{digits: ""},
			want: nil,
		},
		{
			name: "test3",
			args: args{"2"},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterCombinationsBFS(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCombinationsBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
