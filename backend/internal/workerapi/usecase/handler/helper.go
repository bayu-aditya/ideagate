package handler

import (
	entityEndpoint "github.com/bayu-aditya/ideagate/backend/internal/shared/domain/entity/endpoint"
	"github.com/bayu-aditya/ideagate/backend/internal/workerapi/domain/entity"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"reflect"
)

func constructRequestCtxData(c *gin.Context, ctxData *entity.ContextData, setting entityEndpoint.SettingRequest) error {
	// construct query parameters
	queries := map[string]interface{}{}
	if err := c.BindQuery(&queries); err != nil {
		return err
	}

	for fieldName, variable := range setting.Query {
		if variable.Required && queries[fieldName] == nil {
			queries[fieldName], _ = castVariableType(variable.Type, variable.Default)
			continue
		}

		queries[fieldName], _ = castVariableType(variable.Type, queries[fieldName])
	}
	ctxData.SetRequestQuery(queries)

	// construct body json
	jsons := map[string]interface{}{}
	if err := c.BindJSON(&jsons); err != nil {
		return err
	}

	for fieldName, variable := range setting.Json {
		if variable.Required && jsons[fieldName] == nil {
			jsons[fieldName], _ = castVariableType(variable.Type, variable.Default)
			continue
		}

		jsons[fieldName], _ = castVariableType(variable.Type, jsons[fieldName])
	}
	ctxData.SetRequestJson(jsons)

	return nil
}

func castVariableType(varType string, value interface{}) (interface{}, error) {
	switch varType {
	case reflect.String.String():
		return cast.ToStringE(value)
	case reflect.Int32.String():
		return cast.ToInt32E(value)
	case reflect.Int64.String():
		return cast.ToInt64E(value)
	case reflect.Float32.String():
		return cast.ToFloat32E(value)
	case reflect.Float64.String():
		return cast.ToFloat64E(value)
	case reflect.Bool.String():
		return cast.ToBoolE(value)
	}

	return nil, nil
}
