package main

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"aaa", 1},
		{"abc", 3},
		{"abdawet", 6},

		//Edge cases
		{"", 0},
		{"v", 1},

		//chinese support
		{"你好啊", 3},
		{"这z里s是上海!！", 9},
	}

	for _, tt := range tests {
		if maxLen := lengthOfNonRepeatingSubStr(tt.s); maxLen != tt.ans {
			t.Errorf("lengthOfNonRepeatingSubStr(%s): got %d; expected %d", tt.s, maxLen, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s, ans := "这z里s是上海!！", 9
	for i := 0; i < 13; i++ {
		s += s
	}

	b.Logf("len(s) = %d", len(s))
	b.ResetTimer() //从这开始计时

	for i := 0; i < b.N; i++ {
		if maxLen := lengthOfNonRepeatingSubStr(s); maxLen != ans {
			b.Errorf("lengthOfNonRepeatingSubStr(%s): got %d; expected %d", s, maxLen, ans)
		}
	}
}
