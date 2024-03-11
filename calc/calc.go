package calc

import (
	"errors"
	"time"
)

func Add(x int, y int) (int, error) {
	// xとyが0以下の場合はエラーを返す
	if x < 0 || y < 0 {
		return 0, errors.New("x and y must be greater than 0")
	}
	result := x + y

	time.Sleep(3 * time.Second)
	return result, nil
}
