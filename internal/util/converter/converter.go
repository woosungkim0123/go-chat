package converter

import (
	"fmt"
	"strconv"
)

func ConvertToInt(value any) (int, error) {
	switch v := value.(type) {
	case int:
		// 이미 int 타입인 경우
		return v, nil
	case string:
		// 문자열인 경우, strconv.Atoi를 사용하여 int로 변환
		intValue, err := strconv.Atoi(v)
		if err != nil {
			// 변환 실패
			return 0, fmt.Errorf("cannot converter string to int: %v", err)
		}
		return intValue, nil
	default:
		// 그 외 타입은 처리할 수 없음
		return 0, fmt.Errorf("unsupported type: %T", value)
	}
}
