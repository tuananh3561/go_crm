package helper

import (
	"github.com/tuananh3561/go_crm/app/config"
	"time"
)

func TimeNow() int {
	//init the location
	location, _ := time.LoadLocation(config.AppConfigDefault.AppTimeZone)
	//set timezone
	return int(time.Now().In(location).Unix())
}
