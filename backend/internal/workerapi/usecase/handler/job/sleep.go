package job

import "time"

type sleep struct {
	Input StartInput
}

func (s *sleep) Start() (output StartOutput, err error) {
	sleepMs := s.Input.Step.Action.TimeoutMs
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)
	return
}
