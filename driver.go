package view

import (
	. "github.com/chefsgo/base"
)

type (

	// LogDriver view驱动
	Driver interface {
		// 连接到驱动
		Connect(config Config) (Connect, error)
	}
	// LogConnect 日志连接
	Connect interface {
		Open() error
		Health() (Health, error)
		Close() error

		Parse(Body) (string, error)
	}

	Health struct {
		Workload int64
	}

	Helper struct {
		Name   string   `json:"name"`
		Desc   string   `json:"desc"`
		Alias  []string `json:"alias"`
		Action Any      `json:"-"`
	}
)

//Driver 为view模块注册驱动
func (this *Module) Driver(name string, driver Driver, override bool) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if driver == nil {
		panic("Invalid view driver: " + name)
	}

	if override {
		this.drivers[name] = driver
	} else {
		if this.drivers[name] == nil {
			this.drivers[name] = driver
		}
	}
}

func (this *Module) Helper(name string, config Helper, override bool) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	alias := make([]string, 0)
	if name != "" {
		alias = append(alias, name)
	}
	if config.Alias != nil {
		alias = append(alias, config.Alias...)
	}

	for _, key := range alias {
		if override {
			this.helpers[key] = config
		} else {
			if _, ok := this.helpers[key]; ok == false {
				this.helpers[key] = config
			}
		}

	}
}
