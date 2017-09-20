package timing

import (
	"strings"
	"time"

	"github.com/robjporter/go-functions/as"
)

var timers map[string]string

func init() {
	timers = make(map[string]string)
}

func Timer(name string) string {
	name = strings.TrimSpace(name)

	if timers[name] == "" {
		timers[name] = as.ToString(time.Now().Round(time.Second))
	} else {
		end := time.Now().Round(time.Second)
		return as.ToString(end.Sub(as.ToTime(timers[name])))
	}
	return ""
}
