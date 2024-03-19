package repositories

import (
	"io"
	"math/rand"
	"strconv"
	"time"
)

func ReadRequest(body io.ReadCloser) ([]byte, error) {
	data, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ConvertToInt32(id string) (int32, error) {
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return -1, err
	}

	return int32(idInt), nil
}

func RandomFiveDigitNumber() int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	min := 10000
	max := 99999
	return rand.Intn(max-min+1) + min
}
