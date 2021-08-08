package triangle

import (
	"testing" //测试包
)

func TestTriangle(t *testing.T) {
	//测试数据
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 1},
		{13, 35, 37},
		{30000, 40000, 50000},
	}

	//测试逻辑
	for _, tt := range tests {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d, %d); "+"got %d; expectd %d", tt.a, tt.b, actual, tt.c) //输出错误信息
		}
	}
}
