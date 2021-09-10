package date

import "time"

// HumanDateRange 常用时间区间
type HumanDateRange string

const (
	Today          HumanDateRange = "today"          // 今天
	ThisWeek       HumanDateRange = "thisWeek"       // 本周
	ThisMouth      HumanDateRange = "thisMonth"      // 本月
	LastSevenDays  HumanDateRange = "lastSevenDays"  // 近7天
	LastThirtyDays HumanDateRange = "lastThirtyDays" // 近30天
)

// GetYoYRange 获取同比日期
func (h HumanDateRange) GetYoYRange() (ys, ye time.Time) {
	var YoYStart, YoYEnd time.Time
	start, end := h.GetRange()
	var offset time.Duration
	switch h {
	case Today: // 计算周周期的日同比
		offset = -(7 * 24 * time.Hour)
	case LastSevenDays: // 计算月周期的周同比
		offset = -(30 * 24 * time.Hour)
	case LastThirtyDays: // 计算年周期的月同比
		offset = -(12 * 30 * 24 * time.Hour)
	}

	YoYStart = start.Add(offset)
	YoYEnd = end.Add(offset)

	return YoYStart, YoYEnd
}

// GetMoMRange 获取环比日期
func (h HumanDateRange) GetMoMRange() (ys, ye time.Time) {
	var YoYStart, YoYEnd time.Time
	start, end := h.GetRange()
	var offset time.Duration
	switch h {
	case Today: // 昨天为环比
		offset = -(24 * time.Hour)
	case LastSevenDays: // 上周为环比
		offset = -(7 * 24 * time.Hour)
	case LastThirtyDays: // 上月为环比
		offset = -(30 * 24 * time.Hour)
	}

	YoYStart = start.Add(offset)
	YoYEnd = end.Add(offset)

	return YoYStart, YoYEnd
}

// GetRange 获取时间区间
func (h HumanDateRange) GetRange() (start, end time.Time) {
	now := time.Now()
	year, month, day := now.Date()
	startOfDay := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	endOfDay := time.Date(year, month, day, 23, 59, 59, 0, now.Location())

	switch h {
	case Today:
		start = startOfDay
		end = endOfDay
		break
	case LastSevenDays:
		start = startOfDay.Add(-6 * 24 * time.Hour)
		end = endOfDay
		break
	case LastThirtyDays:
		start = startOfDay.Add(-29 * 24 * time.Hour)
		end = endOfDay
		break
	default:
		start = startOfDay
		end = endOfDay
		break
	}
	return
}
