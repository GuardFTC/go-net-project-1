// Package base @Author:冯铁城 [17615007230@163.com] 2023-10-13 15:16:53
package base

import (
	"fmt"
	"time"
)

type jsonTime time.Time

// MarshalJSON 实现它的json序列化方法
func (j jsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(j).Format(time.DateTime))
	return []byte(stamp), nil
}
