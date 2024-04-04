package apperror

type ErrorCode string

const (
	DataBaseProblem   ErrorCode = "DataBaseProblem"
	FailJsonUnmarshal ErrorCode = "FailJsonUnmarshal"

	NotFoundBucket ErrorCode = "NotFoundBucket"

	NotFoundUserByLoginID ErrorCode = "NotFoundUserByLoginID"
	NotFoundUserByID      ErrorCode = "NotFoundUserByID"
)
