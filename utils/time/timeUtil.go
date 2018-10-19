package time

import "time"

func NsToMs(duration time.Duration)  time.Duration {
	return duration*10*1000000*1000

}
