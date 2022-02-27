package thinkutils

import (
	"strconv"
	"time"
)

type datetime struct {
}

func (this datetime) Timestamp() int64 {
	now := time.Now()
	return now.Unix()
}

func (this datetime) TimestampMs() int64 {
	now := time.Now()
	return now.UnixMilli()
}

func (this datetime) CurDatetime() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}

func (this datetime) Hour() int {
	t := time.Now()

	nRet, err := strconv.Atoi(t.Format("15"))
	if err != nil {
		return 0
	}

	return nRet
}

func (this datetime) Yesterday() string {
	return this.DiffDate(-1)
}

func (this datetime) Tomorrow() string {
	return this.DiffDate(1)
}

func (this datetime) DiffDate(nDay int) string {
	var nTimestamp = this.Timestamp()
	nTimestamp += 3600 * 24 * int64(nDay)
	t := time.Unix(nTimestamp, 0)

	return t.Format("2006-01-02")
}

func (this datetime) TimeStampToDate(nTimestamp int64) string {
	t := time.Unix(nTimestamp, 0)

	return t.Format("2006-01-02")
}

func (this datetime) DateToTimestamp(szDate string) int64 {
	t, err := time.Parse("2006-01-02", szDate)
	if err != nil {
		return 0
	}

	return t.Unix()
}

func (this datetime) DateTimeToTimestamp(szDate string) int64 {
	t, err := time.Parse("2006-01-02 15:04:05", szDate)
	if err != nil {
		return 0
	}

	return t.Unix()
}

func (this datetime) TimeStampToDateTime(nTimestamp int64) string {
	t := time.Unix(nTimestamp, 0)

	return t.Format("2006-01-02 15:04:05")
}

func (this datetime) FirstDayOfMonth(szDate string) string {
	t, err := time.Parse("2006-01-02", szDate)
	if err != nil {
		return ""
	}

	strMonth := t.Format("2006-01")

	return strMonth + "-01"
}

func (this datetime) LastDayOfMonth(szDate string) string {
	szFirst := this.FirstDayOfMonth(szDate)
	nTimestamp := this.DateToTimestamp(szFirst)
	nTimestamp += 32 * 3600 * 24

	szNextMonth := this.TimeStampToDate(nTimestamp)
	szNextFirst := this.FirstDayOfMonth(szNextMonth)

	nTimestamp = this.DateToTimestamp(szNextFirst)
	nTimestamp -= 24 * 3600

	return this.TimeStampToDate(nTimestamp)
}

func (this datetime) DateBetweenStartEnd(szDateStart string, szDateEnd string) []string {
	lstRet := make([]string, 2)

	nTimestampStart := this.DateToTimestamp(szDateStart)
	nTimestampEnd := this.DateToTimestamp(szDateEnd)

	nStart := nTimestampStart + 3600*24
	for {
		if nStart >= nTimestampEnd {
			break
		}

		lstRet = append(lstRet, this.TimeStampToDate(nStart))

		nStart += 3600 * 24
	}

	return lstRet
}

func (this datetime) StartEndOfWeek(szDate string) (string, string) {
	//calc start
	var szStart string = ""
	var szEnd string = ""

	nTimpstamp := this.DateToTimestamp(szDate)
	tm := time.Unix(nTimpstamp, 0)
	for {
		if time.Monday == tm.Weekday() {
			szStart = tm.Format("2006-01-02")
			break
		}

		nTimpstamp -= 3600 * 24
		tm = time.Unix(nTimpstamp, 0)
	}

	szEnd = time.Unix(nTimpstamp+6*3600*24, 0).Format("2006-01-02")

	return szStart, szEnd
}
