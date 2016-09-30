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

type TestExceptionArgs struct {
	Arg *string `json:"arg,omitempty"`
}

func (v *TestExceptionArgs) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Arg != nil {
		w, err = wire.NewValueString(*(v.Arg)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *TestExceptionArgs) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Arg = &x
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *TestExceptionArgs) String() string {
	var fields [1]string
	i := 0
	if v.Arg != nil {
		fields[i] = fmt.Sprintf("Arg: %v", *(v.Arg))
		i++
	}
	return fmt.Sprintf("TestExceptionArgs{%v}", strings.Join(fields[:i], ", "))
}

func (v *TestExceptionArgs) MethodName() string {
	return "testException"
}

func (v *TestExceptionArgs) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

type TestExceptionResult struct {
	Err1 *gauntlet.Xception `json:"err1,omitempty"`
}

func (v *TestExceptionResult) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Err1 != nil {
		w, err = v.Err1.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if i > 1 {
		return wire.Value{}, fmt.Errorf("TestExceptionResult should have at most one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Xception_Read(w wire.Value) (*gauntlet.Xception, error) {
	var v gauntlet.Xception
	err := v.FromWire(w)
	return &v, err
}

func (v *TestExceptionResult) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Err1, err = _Xception_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.Err1 != nil {
		count++
	}
	if count > 1 {
		return fmt.Errorf("TestExceptionResult should have at most one field: got %v fields", count)
	}
	return nil
}

func (v *TestExceptionResult) String() string {
	var fields [1]string
	i := 0
	if v.Err1 != nil {
		fields[i] = fmt.Sprintf("Err1: %v", v.Err1)
		i++
	}
	return fmt.Sprintf("TestExceptionResult{%v}", strings.Join(fields[:i], ", "))
}

func (v *TestExceptionResult) MethodName() string {
	return "testException"
}

func (v *TestExceptionResult) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}

var TestExceptionHelper = struct {
	IsException    func(error) bool
	Args           func(arg *string) *TestExceptionArgs
	WrapResponse   func(error) (*TestExceptionResult, error)
	UnwrapResponse func(*TestExceptionResult) error
}{}

func init() {
	TestExceptionHelper.IsException = func(err error) bool {
		switch err.(type) {
		case *gauntlet.Xception:
			return true
		default:
			return false
		}
	}
	TestExceptionHelper.Args = func(arg *string) *TestExceptionArgs {
		return &TestExceptionArgs{Arg: arg}
	}
	TestExceptionHelper.WrapResponse = func(err error) (*TestExceptionResult, error) {
		if err == nil {
			return &TestExceptionResult{}, nil
		}
		switch e := err.(type) {
		case *gauntlet.Xception:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for TestExceptionResult.Err1")
			}
			return &TestExceptionResult{Err1: e}, nil
		}
		return nil, err
	}
	TestExceptionHelper.UnwrapResponse = func(result *TestExceptionResult) (err error) {
		if result.Err1 != nil {
			err = result.Err1
			return
		}
		return
	}
}
