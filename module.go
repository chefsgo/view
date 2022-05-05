package view

import (
	"sync"
	"time"

	. "github.com/chefsgo/base"
	"github.com/chefsgo/chef"
)

func init() {
	chef.Register(NAME, module)
}

var (
	module = &Module{
		config: Config{
			Driver: chef.DEFAULT, Root: "views",
			Shared: "shared", Left: "{%", Right: "%}",
		},
		drivers: make(map[string]Driver, 0),
		helpers: make(map[string]Helper, 0),
	}
)

type (
	// 日志模块定义
	Module struct {
		connected, initialized, launched bool

		mutex   sync.Mutex
		drivers map[string]Driver
		helpers map[string]Helper

		helperActions Map

		config  Config
		connect Connect
	}

	// LogConfig 日志模块配置
	Config struct {
		Driver  string
		Root    string
		Shared  string
		Left    string
		Right   string
		Setting Map
	}

	Body struct {
		View     string
		Site     string
		Language string
		Timezone *time.Location
		Data     Map
		Model    Map
		Helpers  Map
	}
)

func (this *Module) Parse(body Body) (string, error) {
	if this.connect == nil {
		return "", errInvalidConnection
	}

	//view层的helper
	if body.Helpers == nil {
		body.Helpers = Map{}
	}
	for key, helper := range this.helperActions {
		//默认不替换，因为http层携带context的方法，更好用一些
		if _, ok := body.Helpers[key]; !ok {
			body.Helpers[key] = helper
		}
	}

	return this.connect.Parse(body)
}
