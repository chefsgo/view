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
