// Copyright © 2020 Martin Whittington <mwhittington79@googlemail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file

package yaml2hcl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	assert := assert.New(t)
	v := make(map[interface{}]interface{})
	v["subnet_name"] = "subnet-1"
	v["subnet_private"] = true
	tf := Convert(v)
	assert.Contains(tf.Attributes(), "subnet_name")
	assert.Contains(tf.Attributes(), "subnet_private")
	assert.Nil(tf.GetAttribute("some_other_var"))
}

func TestConvertToString(t *testing.T) {
	assert := assert.New(t)
	v := make(map[interface{}]interface{})
	v["subnet_name"] = "subnet-1"
	v["subnet_private"] = true
	tf := ConvertToString(v)
	assert.Contains(tf, "subnet_name")
	assert.Contains(tf, "subnet-1")
	assert.Contains(tf, "subnet_private")
	assert.Contains(tf, "true")
}

func TestGetValue(t *testing.T) {
	assert := assert.New(t)
	v := interface{}("string")
	val := getValue(v)
	assert.True(val.Type().IsPrimitiveType())

	v = true
	val = getValue(v)
	assert.True(val.Type().IsPrimitiveType())

	v = 10
	val = getValue(v)
	assert.True(val.Type().IsPrimitiveType())

	m := make(map[interface{}]interface{})
	k := interface{}("key")
	m[k] = interface{}("value")
	val = getValue(m)
	assert.True(val.Type().IsObjectType())

	l := make([]interface{}, 0)
	val = getValue(l)
	assert.True(val.Type().IsListType())

	li := interface{}("value")
	l = append(l, li)
	val = getValue(l)
	assert.True(val.Type().IsListType())

	v = 3.4 // not handling floats
	val = getValue(v)
	assert.True(val.IsNull())
}
