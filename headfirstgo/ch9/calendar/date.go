package calendar

import (
	"errors"
)

// Date 日期
type Date struct {
	year  int
	month int
	day   int
}

// Year 获取年
func (d *Date) Year() int {
	return d.year
}

// Month 获取月份
func (d *Date) Month() int {
	return d.month
}

// Day 获取日期
func (d *Date) Day() int {
	return d.day
}

// SetYear 设置年
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

// SetMonth 设置月份
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

// SetDay 设置日期
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}
