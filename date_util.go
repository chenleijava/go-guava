package guava

import "time"

//format layout
const (
	date          = "2006-01-02"                 //yyyy-MM-dd
	dateShortTime = "2006-01-02 15:04:00"        //yyyy-MM-dd hh:mm:00
	datetime      = "2006-01-02 15:04:05"        //yyyy-MM-dd hh:mm:ss
	dateTimeStamp = "2006-01-02 15:04:05.000000" ////yyyy-MM-dd hh:mm:ss.ms
	minute        = "15:04"                      //hh:mm
)

//Get minutes
func GetMinute(t *time.Time) string {
	return t.Format(minute)
}

//Get now time ,format  yyyy-MM-dd hh:mm:ss
func GetNowDateTime() string {
	return time.Now().Format(datetime)
}

//Get now time ,format  yyyy-MM-dd hh:mm:00
func GetNowDateShortimeZero() string {
	return time.Now().Format(dateShortTime)
}

//Get now time ,format  yyyy-MM-dd hh:mm:00
func GetDateShortTimeZeroBy(timeStamp int64) string {
	return time.Unix(timeStamp, 0).Format(dateShortTime)
}

//format yyyy-MM-dd hh:mm:ss
func GetDateTimeBy(timeStamp uint64) string {
	return time.Unix(int64(timeStamp), 0).Format(datetime)
}

//
func GetDateTime(timeStamp uint64) time.Time {
	return time.Unix(int64(timeStamp), 0)
}

//format yyyy-MM-dd
func GetNowDate() string {
	return time.Now().Format(date)
}

//format
func FormatDate(t *time.Time) string {
	return t.Format(date)
}

//format yyyy-MM-dd
func GetDateBy(timeStamp uint64) string {
	return time.Unix(int64(timeStamp), 0).Format(date)
}

//format
func GetYesterdayDate() string {
	return time.Now().AddDate(0, 0, -1).Format(date)
}

//format
func GetYesterdayDateTime() time.Time {
	return time.Now().AddDate(0, 0, -1)
}

//format
func GetNowUtc() time.Time {
	return time.Now().UTC()
}

//format
func ParseTime(value string) time.Time {
	t, _ := time.ParseInLocation(datetime, value, time.Local)
	return t
}

//parse time to stamp
func ParseTimeStamp(value string) int64 {
	t, _ := time.ParseInLocation(datetime, value, time.Local)
	return t.Unix()
}

//Get now timestamp
func GetTimeStampSeconds(value string) int64 {
	return time.Now().Unix()
}

//Get mills second
func GetTimeStampTimeMillis() int64 {
	return time.Now().UnixNano() / 1e6
}
