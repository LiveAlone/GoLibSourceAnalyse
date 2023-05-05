package office

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func TestCal(t *testing.T) {
	// 生成mock对象
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//// 创建mock对象
	//mockCalculator := NewMockCalculator(ctrl)
	//
	//// 预期Calculator的Add方法将被调用一次，参数为2和3，返回值为5
	//mockCalculator.EXPECT().Add(2, 3).Return(5)
	//
	//// 预期Calculator的Subtract方法将被调用一次，参数为2和3，返回值为-1
	//mockCalculator.EXPECT().Subtract(2, 3).Return(-1)
	//
	//// 调用DoubleCalculator，传入mockCalculator作为参数
	//result := DoubleCalculator(mockCalculator, 2, 3)
	//
	//// 检查DoubleCalculator的返回值是否为预期的4
	//if result != 4 {
	//	t.Errorf("Expected 4, but got %d", result)
	//}
}
