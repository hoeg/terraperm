package terraform

import "io"

type Tracer struct {
	exe Executor
}

func NewTracer(exe Executor) Tracer {
	return Tracer{
		exe: exe,
	}
}

func (t Tracer) MakeTrace(w io.Writer) error {
	err := t.exe.Init()
	if err != nil {
		return err
	}

	defer t.exe.Destroy()
	err = t.exe.Apply(w)
	if err != nil {
		return err
	}

	return nil
}
