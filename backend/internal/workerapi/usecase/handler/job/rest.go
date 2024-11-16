package job

import (
	"encoding/json"
	"io"
	"net/http"
)

type rest struct {
	Input StartInput
}

func (j *rest) Start() (StartOutput, error) {
	var (
		ctx        = j.Input.Ctx
		dataCtx    = j.Input.DataCtx
		dataSource = j.Input.DataSource
		step       = j.Input.Step
		output     = StartOutput{}
	)

	//construct request url, path is getting by template
	path, err := step.Action.Path.GetValueString(step.Id, dataCtx)
	if err != nil {
		return output, err
	}
	reqUrl := dataSource.Config.Host + path

	// doing request
	req, err := http.NewRequestWithContext(ctx, step.Action.Method, reqUrl, nil)
	if err != nil {
		return output, err
	}

	// construct request header
	for headerKey, headerVar := range step.Action.Headers {
		headerVal, err := headerVar.GetValueString(step.Id, dataCtx)
		if err != nil {
			return output, err
		}

		req.Header.Set(headerKey, headerVal)
	}

	// doing http call
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return output, err
	}

	// set status code into context data
	dataCtx.SetStepStatusCode(step.Id, resp.StatusCode)

	// parse response json into context data
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return output, err
	}

	var respData any
	if len(respBytes) > 0 {
		if err = json.Unmarshal(respBytes, &respData); err != nil {
			return output, err
		}
	}
	dataCtx.SetStepDataBody(step.Id, respData)

	// Determine next step
	for _, returnItem := range step.Returns {
		output.NextStepIds = append(output.NextStepIds, returnItem.NextStepId)
	}

	return output, nil
}
