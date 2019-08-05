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
// Ex. GetJobPropertyString("gemfire/tls/enabled")
// gemfire:
//   tls:
//      enabled: "true"
//
// Ex. GetJobPropertyString("gemfire/name=tls/enabled")
// gemfire:
//   - name: tls
//     enabled: true
func (o *Operation) GetJobPropertyString(path string) (string, error) {
	if o.error != nil {
		return "", o.error
	}

	if len(o.jobs) > 1 {
		return "", errors.New("failed to execute 'GetJobPropertyString': not implemented for cases where multiple jobs are retrieved")
	}

	for _, job := range o.jobs {
		result, err := patch.FindOp{
			Path: patch.MustNewPointerFromString(path),
		}.Apply(job.Properties)

		if err != nil {
			return "", err
		}

		strResult, ok := result.(string)
		if !ok {
			return "", fmt.Errorf("failed to find string value at '%s', instead '%v'(%T) was found", path, result, result)
		}

		return strResult, nil
	}

	return "", fmt.Errorf("job property '%s' not found", path)
}
