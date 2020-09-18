package terraform

type Tracer struct {
	exe Executor
}

func (t Tracer) makeTrace() (io.ReadCloser, error) {
	err := exe.Init()
	if err != nil {
		return nil, err
	}

	defer exe.Destroy()
	err = exe.Apply(trace)
	if err != nil {
		return nil, err
	}

	return nil, nil
} 