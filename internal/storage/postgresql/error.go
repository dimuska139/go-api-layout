package postgresql

import "fmt"

type ErrQueryExecutionFailed struct {
	Query string
	Args  []any
	Err   error
}

func (e *ErrQueryExecutionFailed) Error() string {
	return fmt.Sprintf("query execution failed: %s", e.Err.Error())
}

func NewErrQueryExecutionFailed(query string, args []any, err error) *ErrQueryExecutionFailed {
	return &ErrQueryExecutionFailed{
		Query: query,
		Args:  args,
		Err:   err,
	}
}
