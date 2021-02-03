package calendar

import (
	"errors"
	"unicode/utf8"

	"github.com/guozhe001/go-learn/headfirstgo/ch9/calendar"
)

// Event 事件
type Event struct {
	title string
	// 内嵌Date
	calendar.Date
}

// SetTitle 设置标题
func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("invalid title")
	}
	e.title = title
	return nil
}

// Title 获取title
func (e *Event) Title() string {
	return e.title
}
