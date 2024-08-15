package main

import (
	"errors"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestCheckBalanceOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	numbers := [2]int{500, 100}
	var expected error
	c := NewMockcalc(ctrl)

	c.EXPECT().Sub(500, 100).Return(400)

	result := checkBalance(c, numbers)

	if result != expected {
		t.Errorf("Неверный результат, должно быть %v, а по факту %v", expected, result)
	}
}

func TestCheckBalanceFail(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	numbers := [2]int{10, 100}
	var expected error = errors.New("not enough balance")
	c := NewMockcalc(ctrl)

	c.EXPECT().Sub(gomock.Any(), gomock.Any()).Return(-1000)

	result := checkBalance(c, numbers)

	if result.Error() != expected.Error() {
		t.Errorf("Неверный результат, должно быть %v, а по факту %v", expected, result)
	}
}
