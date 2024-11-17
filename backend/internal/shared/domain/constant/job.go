package constant

var StepIdRequest = "request"

type JobType string

var (
	JobTypeStart         JobType = "start"
	JobTypeEnd           JobType = "end"
	JobTypeTransformerJS JobType = "transformerJS"
	JobTypeCondition     JobType = "condition"
	JobTypeRest          JobType = "rest"
	JobTypeMysql         JobType = "mysql"
	JobTypePostgresql    JobType = "postgresql"
	JobTypeRedis         JobType = "redis"
)
