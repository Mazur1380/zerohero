package main

import "testing"

func TestSumm(t *testing.T) {
	tests := []struct {
		Input    []int
		Expected int
	}{
		{
			Input:    []int{1, 2, 3},
			Expected: 6,
		}, {
			Input:    make([]int, 5),
			Expected: 0,
		}, {
			Input:    nil,
			Expected: 0,
		},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			Actual := summ(test.Input)
			if Actual != test.Expected {
				t.Errorf("Не правильный результат, ожидали %v, получили %v ", test.Expected, Actual)
			}
		})
	}

	// 	actualRes := summ(x)
	// 	if actualRes != expectedRes {
	// 		t.Errorf("Неверный результат, ожидали %v, получили %v", expectedRes, actualRes)
	// 	}

	// 	x = make([]int, 5)
	// 	expectedRes = 0

	// 	actualRes = summ(x)
	// 	if actualRes != expectedRes {
	// 		t.Errorf("Неверный результат, ожидали %v, получили %v", expectedRes, actualRes)
	// 	}
	// 	x = nil
	// 	expectedRes = 0

	// 	actualRes = summ(x)
	// 	if actualRes != expectedRes {
	// 		t.Errorf("Неверный результат, ожидали %v, получили %v", expectedRes, actualRes)
	// 	}

	// 	x = []int{1}
	// 	expectedRes = 1

	// 	actualRes = summ(x)
	// 	if actualRes != expectedRes {
	// 		t.Errorf("Неверный результат, ожидали %v, получили %v", expectedRes, actualRes)
	// 	}

}
