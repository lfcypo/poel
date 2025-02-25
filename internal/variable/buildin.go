package variable

import (
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/lfcypo/poel/internal/typing"
)

var vals *buildInVals
var onceGetBuildInVals sync.Once

type buildInVals struct {
	os struct {
		name typing.String
		arch typing.String
	}
	hostname typing.String
}

func getBuildInVals() *buildInVals {
	onceGetBuildInVals.Do(func() {
		vals = &buildInVals{}
		vals.os.name = typing.String(runtime.GOOS)
		vals.os.arch = typing.String(runtime.GOARCH)

		hostname, err := os.Hostname()
		if err != nil {
			// TODO: handle error
		}
		vals.hostname = typing.String(hostname)

	})
	return vals
}

type buildInVars struct {
	datetime  typing.String
	timestamp typing.Integer
	date      struct {
		year    typing.Integer
		month   typing.Integer
		day     typing.Integer
		weekday typing.Integer
	}
	time struct {
		hour   typing.Integer
		minute typing.Integer
		second typing.Integer
		//microsecond typing.Integer
		//nanosecond  typing.Integer
	}
}

func getBuildinVars() *buildInVars {
	timeNow := time.Now()

	vars := &buildInVars{}
	vars.datetime = typing.String(timeNow.Format(time.RFC3339))
	vars.timestamp = typing.Integer(timeNow.Unix())
	vars.date.year = typing.Integer(timeNow.Year())
	vars.date.month = typing.Integer(int(timeNow.Month()))
	vars.date.day = typing.Integer(timeNow.Day())
	vars.date.weekday = typing.Integer(int(timeNow.Weekday()))
	vars.time.hour = typing.Integer(timeNow.Hour())
	vars.time.minute = typing.Integer(timeNow.Minute())
	vars.time.second = typing.Integer(timeNow.Second())
	return vars
}

type BuildInVars struct {
	*buildInVars
	*buildInVals
}

func GetBuildInVars() *BuildInVars {
	return &BuildInVars{
		getBuildinVars(),
		getBuildInVals(),
	}
}
