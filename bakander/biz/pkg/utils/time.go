/*
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package utils

import (
	"saas/biz/pkg/errno"
	"time"
)

// MillTimeStampToTime convert ms timestamp to time.Time
func MillTimeStampToTime(timestamp int64) time.Time {
	second := timestamp / 1000
	nano := timestamp % 1000 * 1000000
	return time.Unix(second, nano)
}

// SecondTimeStampToTime convert s timestamp to time.Time
func SecondTimeStampToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

// 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

// 获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

func GetStringDateTime(d string) (time.Time, error) {
	at, err := time.ParseInLocation(time.DateTime, d, time.Local)
	if err != nil {
		return GetZeroTime(time.Now()), errno.DateErr
	}
	return at, nil
}
func GetStringDateOnlyZeroTime(d string) (time.Time, error) {
	at, err := time.ParseInLocation(time.DateOnly, d, time.Local)
	if err != nil {
		return GetZeroTime(time.Now()), errno.DateErr
	}
	return GetZeroTime(at), nil
}

// 获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// 获取某一天的23:59:59点时间
func GetEndTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, 0, d.Location())
}

// 获取近七天的时间
// 如传入time.Now(), 返回当前时间的前7天的时间
func GetNearlySevenDaysTime(d time.Time) []time.Time {
	var times []time.Time
	for i := 0; i < 7; i++ {
		times = append(times, GetZeroTime(d.AddDate(0, 0, -i)))
	}
	return times
}

// 获取本周周一的日期
func GetFirstDateOfWeek() time.Time {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
}

// GetLastWeekFirstDate 获取上周的周一日期
func GetLastWeekFirstDate() time.Time {
	thisWeekMonday := GetFirstDateOfWeek()
	return thisWeekMonday.AddDate(0, 0, -7)
}

// GetLastWeekFirstDate 获取上周的周日日期
func GetLastWeekEndDate() time.Time {
	thisWeekMonday := GetFirstDateOfWeek()
	return thisWeekMonday.AddDate(0, 0, -1)
}

// 获取本月的起始日期
func GetFirstMonthDaysTime() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
}

// 获取上个月的开始日期
func GetLastFirstMonthDaysTime() time.Time {
	firstDayOfThisMonth := GetFirstMonthDaysTime()
	return firstDayOfThisMonth.AddDate(0, -1, 0)

}

// 获取上个月的结束日期
func GetLastEndMonthDaysTime() time.Time {
	firstDayOfThisMonth := GetFirstMonthDaysTime()
	d := firstDayOfThisMonth.AddDate(0, 0, -1)
	return time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, 0, d.Location())
}

// 获取上周的时间
func GetWeekDaysTime() []time.Time {
	var times []time.Time
	for i := 1; i < 8; i++ {
		times = append(times, GetZeroTime(GetFirstDateOfWeek().AddDate(0, 0, -i)))
	}
	return times
}
