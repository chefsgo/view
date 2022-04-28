package view

import (
	. "github.com/chefsgo/base"
)

func (this *Module) Register(name string, value Any, override bool) {
	switch obj := value.(type) {
	case Driver:
		this.Driver(name, obj, override)
	}
}
func (this *Module) Configure(value Any) {
	if cfg, ok := value.(Config); ok {
		this.config = cfg
		return
	}

	var config Map
	if global, ok := value.(Map); ok {
		if vvv, ok := global["view"].(Map); ok {
			config = vvv
		}
	}
	if config == nil {
		return
	}

	//设置值
	if driver, ok := config["driver"].(string); ok {
		this.config.Driver = driver
	}
	if vv, ok := config["root"].(string); ok {
		this.config.Root = vv
	}
	if vv, ok := config["shared"].(string); ok {
		this.config.Shared = vv
	}
	if vv, ok := config["left"].(string); ok {
		this.config.Left = vv
	}
	if vv, ok := config["right"].(string); ok {
		this.config.Right = vv
	}

	if setting, ok := config["setting"].(Map); ok {
		this.config.Setting = setting
	}
}
func (this *Module) Initialize() {
	if this.initialized {
		return
	}

	if this.config.Root == "" {
		this.config.Root = "views"
	}
	if this.config.Shared == "" {
		this.config.Shared = "shared"
	}
	if this.config.Left == "" {
		this.config.Left = "{%"
	}
	if this.config.Right == "" {
		this.config.Right = "%}"
	}

	this.helperActions = Map{}
	for key, helper := range this.helpers {
		if helper.Action != nil {
			this.helperActions[key] = helper.Action
		}
	}

	this.initialized = true
}
func (this *Module) Connect() {
	if this.connected {
		return
	}

	driver, ok := this.drivers[this.config.Driver]
	if ok == false {
		panic("Invalid view driver: " + this.config.Driver)
	}

	// 建立连接
	connect, err := driver.Connect(this.config)
	if err != nil {
		panic("Failed to connect to view: " + err.Error())
	}

	// 打开连接
	err = connect.Open()
	if err != nil {
		panic("Failed to open view connect: " + err.Error())
	}

	// 保存连接，设置管道大小
	this.connect = connect

	this.connected = true
}
func (this *Module) Launch() {
	if this.launched {
		return
	}

	this.launched = true
}
func (this *Module) Terminate() {
	if this.connect != nil {
		this.connect.Close()
	}

	this.launched = false
	this.connected = false
	this.initialized = false
}
