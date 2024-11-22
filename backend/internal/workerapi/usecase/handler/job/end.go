package job

type end struct {
	input StartInput
}

func (e *end) Start() (StartOutput, error) {
	return StartOutput{}, nil
}
