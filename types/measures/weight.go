package measures

import "strconv"

// KG 重量 序列化时保留 3 位小数
type KG float64

// MarshalJSON 千克重量精度三位
func (k KG) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatFloat(float64(k), 'f', 3, 64)), nil
}
