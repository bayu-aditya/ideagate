# Variable

Object variable from request

- `.Req.Header`
- `.Req.Query`
- `.Req.Json`

Object variable from different step

- `.Step.{StepId}.Var`
- `.Step.{StepId}.Data`
- `.Step.{StepId}.Out`

Object variable from current step

- `.Var`
- `.Data.Body`: for http response
- `.Data.Query.{QueryId}`: for query result
- `.Data.StatusCode`: for http status code response
