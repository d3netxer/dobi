package iface

import (
	"github.com/dnephin/dobi/logging"
	"github.com/dnephin/dobi/tasks/context"
)

// Task is an interface implemented by all tasks
type Task interface {
	logging.LogRepresenter
	Name() string
	// TODO: possibly remove? currently unused
	Prepare(*context.ExecuteContext) error
	Run(*context.ExecuteContext) error
	Stop(*context.ExecuteContext) error
}