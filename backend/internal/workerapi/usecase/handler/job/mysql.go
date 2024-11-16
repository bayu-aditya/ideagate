package job

import "github.com/bayu-aditya/ideagate/backend/pkg/utils/errors"

type mysql struct {
	Input StartInput
}

func (j *mysql) Start() (output StartOutput, err error) {
	var (
		ctx        = j.Input.Ctx
		ctxData    = j.Input.DataCtx
		dataSource = j.Input.DataSource
		step       = j.Input.Step
	)

	if j.Input.DataSource.MysqlConn == nil {
		err = errors.New("empty mysql connection")
		return
	}

	session := dataSource.MysqlConn.WithContext(ctx)

	if len(step.Action.Queries) > 1 {
		session.Begin()
		defer session.Commit()
	}

	for _, queryItem := range step.Action.Queries {
		// run template for query template
		query, errQuery := queryItem.Query.GetValueString(step.Id, ctxData)
		if errQuery != nil {
			err = errQuery
			return
		}

		// run template for parameters
		var paramsParsed []interface{}
		for _, param := range queryItem.Parameters {
			paramParsed, errParsed := param.GetValue(step.Id, ctxData)
			if errParsed != nil {
				err = errParsed
				return
			}

			paramsParsed = append(paramsParsed, paramParsed)
		}

		// execute query
		var rows []map[string]interface{}
		if err = session.Raw(query, paramsParsed...).Scan(&rows).Error; err != nil {
			return
		}

		// write into context data
		ctxData.SetStepDataBody(step.Id, rows)
	}

	return
}
