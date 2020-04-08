// Copyright © 2020 Martin Whittington <mwhittington79@googlemail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file

package yaml2hcl

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

func getValue(value interface{}) cty.Value {
	switch t := value.(type) {
	case string:
		return cty.StringVal(t)
	case int:
		return cty.NumberIntVal(int64(t))
	case bool:
		return cty.BoolVal(t)
	case map[interface{}]interface{}:
		m := make(map[string]cty.Value)
		for k, v := range t {
			m[fmt.Sprintf("%v", k)] = getValue(v)
		}
		return cty.ObjectVal(m)
	case []interface{}:
		if len(t) < 1 {
			return cty.ListValEmpty(cty.String)
		}
		vals := []cty.Value{}
		for _, n := range t {
			val := getValue(n)
			vals = append(vals, val)
		}
		return cty.ListVal(vals)
	default:
		// type not handled yet
		return cty.NullVal(cty.String)
	}
}

// Convert converts a map of interfaces (unmarshalled from Yaml) and returns the HCL body
func Convert(vars map[interface{}]interface{}) *hclwrite.Body {
	f := hclwrite.NewEmptyFile()
	for key, value := range vars {
		f.Body().SetAttributeValue(fmt.Sprintf("%v", key), getValue(value))
	}
	return f.Body()
}

// ConvertToString converts a map of interfaces (unmarshalled from Yaml) and returns the HCL body as a string
func ConvertToString(vars map[interface{}]interface{}) string {
	f := hclwrite.NewEmptyFile()
	for key, value := range vars {
		f.Body().SetAttributeValue(fmt.Sprintf("%v", key), getValue(value))
	}
	return string(f.Bytes())
}
