package alias

import (
	"fmt"

	"github.com/dnephin/dobi/config"
	"github.com/dnephin/dobi/tasks/common"
	"github.com/dnephin/dobi/tasks/iface"
)

// GetTaskConfig returns a new TaskConfig for the action
func GetTaskConfig(name, act string, conf *config.AliasConfig) (iface.TaskConfig, error) {
	switch act {
	case "", "run":
		return iface.NewTaskConfig(
			common.NewTaskName(name, "run"),
			conf,
			RunDeps(conf),
			NewTask,
		), nil
	case "remove", "rm":
		return iface.NewTaskConfig(
			common.NewTaskName(name, "rm"),
			conf,
			RemoveDeps(conf),
			NewTask,
		), nil
	default:
		return nil, fmt.Errorf("Invalid alias action %q for task %q", act, name)
	}
}

// NewTask creates a new Task object
func NewTask(name common.TaskName, conf config.Resource) iface.Task {
	// TODO: cleaner way to avoid this cast?
	return &Task{name: name, config: conf.(*config.AliasConfig)}
}

// RunDeps returns the dependencies for the run action
func RunDeps(conf config.Resource) func() []string {
	return func() []string {
		return conf.Dependencies()
	}
}

// RemoveDeps returns the dependencies for the remove action
func RemoveDeps(conf config.Resource) func() []string {
	return func() []string {
		confDeps := conf.Dependencies()
		deps := []string{}
		for i := len(confDeps); i > 0; i-- {
			taskname := common.ParseTaskName(confDeps[i-1])
			deps = append(deps, taskname.Resource()+":"+"rm")
		}
		return deps
	}
}
