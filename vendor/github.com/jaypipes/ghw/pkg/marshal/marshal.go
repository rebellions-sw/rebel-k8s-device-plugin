//
// Use and distribution licensed under the Apache license version 2.
//
// See the COPYING file in the root project directory for full text.
//

package marshal

import (
	"encoding/json"

	"github.com/jaypipes/ghw/pkg/context"
	yaml "gopkg.in/yaml.v3"
)

// SafeYAML returns a string after marshalling the supplied parameter into YAML.
func SafeYAML(ctx *context.Context, p interface{}) string {
	b, err := json.Marshal(p)
	if err != nil {
		ctx.Warn("error marshalling JSON: %s", err)
		return ""
	}

	var jsonObj interface{}
	if err := yaml.Unmarshal(b, &jsonObj); err != nil {
		ctx.Warn("error converting JSON to YAML: %s", err)
		return ""
	}

	yb, err := yaml.Marshal(jsonObj)
	if err != nil {
		ctx.Warn("error marshalling YAML: %s", err)
		return ""
	}

	return string(yb)
}

// SafeJSON returns a string after marshalling the supplied parameter into
// JSON. Accepts an optional argument to trigger pretty/indented formatting of
// the JSON string.
func SafeJSON(ctx *context.Context, p interface{}, indent bool) string {
	var b []byte
	var err error
	if !indent {
		b, err = json.Marshal(p)
	} else {
		b, err = json.MarshalIndent(&p, "", "  ")
	}
	if err != nil {
		ctx.Warn("error marshalling JSON: %s", err)
		return ""
	}
	return string(b)
}
