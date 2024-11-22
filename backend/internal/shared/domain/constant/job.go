package constant

var (
	StepIdStart = "start"
	StepIdEnd   = "end"
)

type JobType string

var (
	JobTypeStart         JobType = "start"
	JobTypeEnd           JobType = "end"
	JobTypeSleep         JobType = "sleep"
	JobTypeTransformerJS JobType = "transformerJS"
	JobTypeCondition     JobType = "condition"
	JobTypeRest          JobType = "rest"
	JobTypeMysql         JobType = "mysql"
	JobTypePostgresql    JobType = "postgresql"
	JobTypeRedis         JobType = "redis"
)
