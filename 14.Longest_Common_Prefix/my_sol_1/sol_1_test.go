package my_sol_1

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{[]string{"flower", "flow", "flight"}}, "fl"},
		{"", args{[]string{"dog", "racecar", "car"}}, ""},
		{"", args{[]string{}}, ""},
		{"", args{[]string{"aa", "a"}}, "a"},
		{"", args{[]string{"aaa", "aa", "aaa"}}, "aa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
