package job

import (
	"fmt"

	"github.com/bayu-aditya/ideagate/backend/internal/shared/domain/constant"
)

type IJob interface {
	Start() (StartOutput, error)
}

func New(jobType constant.JobType, input StartInput) (IJob, error) {
	switch jobType {
	case constant.JobTypeStart:
		return &start{Input: input}, nil

	case constant.JobTypeEnd:
		return &end{input: input}, nil

	case constant.JobTypeSleep:
		return &sleep{Input: input}, nil

	case constant.JobTypeRest:
		return &rest{Input: input}, nil

	case constant.JobTypeMysql:
		return &mysql{Input: input}, nil
	}

	return nil, fmt.Errorf("unknown job type (%s)", jobType)
}
