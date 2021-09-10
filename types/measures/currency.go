package measures

import "strconv"

// Yuan 元
type Yuan float64

// MarshalJSON 按元为单位时,货币精度两位
func (y Yuan) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatFloat(float64(y), 'f', 2, 64)), nil
}

// TODO 银行家算法
