package helper

import (
	"time"
)

func TimeNow() int {
	//init the location
	//location, _ := time.LoadLocation(config.AppConfigDefault.AppTimeZone)
	//set timezone,
	//return int(time.Now().In(location).Unix())
	return int(time.Now().Unix())
}
