package terraform

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type Executor struct{}

func NewExecutor() (Executor, error) {
	cmd := exec.Command("which", "terraform")
	err := cmd.Run()
	if err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			fmt.Printf("terraform was not found on the path\n")
		}
		return Executor{}, err
	}
	return Executor{}, nil
}

func (Executor) Init() error {
	out := bytes.NewBuffer(nil)
	return cmd(out, "init")
}

func (Executor) Apply(out io.Writer) error {
	return cmd(out, "apply", "--auto-approve") //No input
}

func (Executor) Destroy() error {
	out := bytes.NewBuffer(nil)
	return cmd(out, "destroy", "--auto-approve")
}

func cmd(out io.Writer, action ...string) error {
	cmd := exec.Command("terraform", action...)
	cmd.Env = append(os.Environ(),
		"TF_LOG=DEBUG",
	)
	cmd.Stderr = out
	return cmd.Run()
}
