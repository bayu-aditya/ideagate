package job

import (
	"fmt"

	"github.com/bayu-aditya/ideagate/backend/core/model/constant"
)

type ErrActionConfigEmpty struct {
	jobType constant.JobType
	stepId  string
}

func (e *ErrActionConfigEmpty) Error() string {
	return fmt.Sprintf("action config empty, job type: %s, step id: %s", e.jobType, e.stepId)
}
