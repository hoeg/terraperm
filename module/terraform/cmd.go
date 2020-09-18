package terraform

import (
	"os"
	"os/exec"
	"io"
)

type Executor struct {}

func NewExecutor() (Executor, error) {
	cmd := exec.Command("which", "terraform")
	err := cmd.Run()
	if err != nil {
		return Executor{}, err
	}
	return Executor{}, nil
}

func (Executor) Init() error {
	return nil
}

func (Executor) Apply() (io.Writer, error) {
	return nil, nil
}

func (Executor) Destroy() error {
	return nil
}

func cmd(action string) (io.ReadCloser, error) {
	run := exec.Command("terraform", action)
	run.Env = append(os.Environ(),
		"TF_LOG=DEBUG",
	)
	return nil, nil
}

