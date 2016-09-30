// Code generated by thriftrw
// @generated

// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package thrifttest

import (
	"errors"
	"fmt"
	"go.uber.org/yarpc/crossdock/thrift/gauntlet"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type TestEnumArgs struct {
	Thing *gauntlet.Numberz `json:"thing,omitempty"`
}

func (v *TestEnumArgs) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Thing != nil {
		w, err = v.Thing.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Numberz_Read(w wire.Value) (gauntlet.Numberz, error) {
	var v gauntlet.Numberz
	err := v.FromWire(w)
	return v, err
}

func (v *TestEnumArgs) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TI32 {
				var x gauntlet.Numberz
				x, err = _Numberz_Read(field.Value)
				v.Thing = &x
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *TestEnumArgs) String() string {
	var fields [1]string
	i := 0
	if v.Thing != nil {
		fields[i] = fmt.Sprintf("Thing: %v", *(v.Thing))
		i++
	}
	return fmt.Sprintf("TestEnumArgs{%v}", strings.Join(fields[:i], ", "))
}

func (v *TestEnumArgs) MethodName() string {
	return "testEnum"
}

func (v *TestEnumArgs) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

type TestEnumResult struct {
	Success *gauntlet.Numberz `json:"success,omitempty"`
}

func (v *TestEnumResult) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("TestEnumResult should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *TestEnumResult) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TI32 {
				var x gauntlet.Numberz
				x, err = _Numberz_Read(field.Value)
				v.Success = &x
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("TestEnumResult should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *TestEnumResult) String() string {
	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", *(v.Success))
		i++
	}
	return fmt.Sprintf("TestEnumResult{%v}", strings.Join(fields[:i], ", "))
}

func (v *TestEnumResult) MethodName() string {
	return "testEnum"
}

func (v *TestEnumResult) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}

var TestEnumHelper = struct {
	IsException    func(error) bool
	Args           func(thing *gauntlet.Numberz) *TestEnumArgs
	WrapResponse   func(gauntlet.Numberz, error) (*TestEnumResult, error)
	UnwrapResponse func(*TestEnumResult) (gauntlet.Numberz, error)
}{}

func init() {
	TestEnumHelper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	TestEnumHelper.Args = func(thing *gauntlet.Numberz) *TestEnumArgs {
		return &TestEnumArgs{Thing: thing}
	}
	TestEnumHelper.WrapResponse = func(success gauntlet.Numberz, err error) (*TestEnumResult, error) {
		if err == nil {
			return &TestEnumResult{Success: &success}, nil
		}
		return nil, err
	}
	TestEnumHelper.UnwrapResponse = func(result *TestEnumResult) (success gauntlet.Numberz, err error) {
		if result.Success != nil {
			success = *result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}
