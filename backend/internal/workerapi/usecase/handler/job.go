package handler

import (
	"context"
	"encoding/json"
	"github.com/bayu-aditya/ideagate/backend/internal/shared/domain/constant"
	"github.com/bayu-aditya/ideagate/backend/internal/shared/domain/entity/datasource"
	"github.com/bayu-aditya/ideagate/backend/internal/shared/domain/entity/endpoint"
	"github.com/bayu-aditya/ideagate/backend/internal/workerapi/domain/entity"
	"io"
	"net/http"
)

type iJob interface {
	Start(ctx context.Context, ctxData *entity.ContextData, dataSource datasource.DataSource, step endpoint.Step, action endpoint.Action) error
}

func newJob(jobType constant.DataSourceType) iJob {
	switch jobType {
	case constant.DataSourceTypeRest:
		return &jobRest{}
	}

	return nil
}

type jobRedis struct{}

type jobMysql struct{}

type jobPostgres struct{}

type jobRest struct{}

func (j *jobRest) Start(ctx context.Context, ctxData *entity.ContextData, dataSource datasource.DataSource, step endpoint.Step, action endpoint.Action) error {
	// construct request url, path is getting by template
	reqUrl := dataSource.Config.Host + ctxData.GetValues(action.Path)

	// doing request
	req, err := http.NewRequestWithContext(ctx, action.Method, reqUrl, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	// write into context data
	dataAction := entity.ContextStepDataAction{
		StatusCode: resp.StatusCode,
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if len(respBytes) > 0 {
		if err = json.Unmarshal(respBytes, &dataAction.Body); err != nil {
			return err
		}
	}

	ctxData.SetStepActionData(step.Id, action.Id, dataAction)

	return nil
}
