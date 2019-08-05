package operation

import (
	"errors"
	"fmt"

	"github.com/abg/go-patch/patch"
)

// The method takes a yml path in the form of a unix path and
// any arbitrary value. The null check for o.job is explicitly
// omitted for brevity.
//
// Ex. GetJobPropertyInt("gemfire/tls/enabled")
// gemfire:
//   tls:
//      enabled: 123
//
// Ex. GetJobPropertyInt("gemfire/name=tls/enabled")
// gemfire:
//   - name: tls
//     enabled: 123
func (o *Operation) GetJobPropertyInt(path string) (int, error) {
	if o.error != nil {
		return 0, o.error
	}

	if len(o.jobs) > 1 {
		return 0, errors.New("failed to execute 'GetJobPropertyInt': not implemented for cases where multiple jobs are retrieved")
	}

	for _, job := range o.jobs {
		result, err := patch.FindOp{
			Path: patch.MustNewPointerFromString(path),
		}.Apply(job.Properties)

		if err != nil {
			return 0, err
		}

		intResult, ok := result.(int)
		if !ok {
			return 0, fmt.Errorf("failed to find int value at '%s', instead '%v'(%T) was found", path, result, result)
		}

		return intResult, nil
	}

	return 0, fmt.Errorf("job property '%s' not found", path)
}