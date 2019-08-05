package operation

import (
	"github.com/abg/go-patch/patch"
)

// The method takes a yml path in the form of a unix path and
// any arbitrary value. The null check for o.job is explicitly
// omitted for brevity.
//
// Ex. AddJobProperty("gemfire/tls/enabled", true)
// gemfire:
//   tls:
//      enabled: true
//
// Ex. AddJobProperty("gemfire/name=tls/enabled", true)
// gemfire:
//   - name: tls
//     enabled: true
func (o *Operation) AddJobProperty(path string, value interface{}) *Operation {
	if o.error != nil {
		return o
	}

	for _, job := range o.jobs {
		if job.Properties == nil {
			job.Properties = make(map[string]interface{})
		}

		_, o.error = patch.ReplaceOp{
			Path:  patch.MustNewPointerFromString(path),
			Value: value,
		}.Apply(job.Properties)
	}

	return o
}
