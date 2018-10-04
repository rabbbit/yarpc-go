// Code generated by thriftrw-plugin-yarpc
// @generated

package storeserver

import (
	"context"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/encoding/thrift"
	"go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/atomic"
	"go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/atomic/readonlystoreserver"
)

// Interface is the server-side interface for the Store service.
type Interface interface {
	readonlystoreserver.Interface

	CompareAndSwap(
		ctx context.Context,
		Request *atomic.CompareAndSwap,
	) error

	Forget(
		ctx context.Context,
		Key *string,
	) error

	Increment(
		ctx context.Context,
		Key *string,
		Value *int64,
	) error
}

// New prepares an implementation of the Store service for
// registration.
//
// 	handler := StoreHandler{}
// 	dispatcher.Register(storeserver.New(handler))
func New(impl Interface, opts ...thrift.RegisterOption) []transport.Procedure {
	h := handler{impl}
	service := thrift.Service{
		Name: "Store",
		Methods: []thrift.Method{

			thrift.Method{
				Name: "compareAndSwap",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.CompareAndSwap),
				},
				Signature:    "CompareAndSwap(Request *atomic.CompareAndSwap)",
				ThriftModule: atomic.ThriftModule,
			},

			thrift.Method{
				Name: "forget",
				HandlerSpec: thrift.HandlerSpec{

					Type:   transport.Oneway,
					Oneway: thrift.OnewayHandler(h.Forget),
				},
				Signature:    "Forget(Key *string)",
				ThriftModule: atomic.ThriftModule,
			},

			thrift.Method{
				Name: "increment",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.Increment),
				},
				Signature:    "Increment(Key *string, Value *int64)",
				ThriftModule: atomic.ThriftModule,
			},
		},
	}

	procedures := make([]transport.Procedure, 0, 3)

	procedures = append(
		procedures,
		readonlystoreserver.New(
			impl,
			append(
				opts,
				thrift.Named("Store"),
			)...,
		)...,
	)
	procedures = append(procedures, thrift.BuildProcedures(service, opts...)...)
	return procedures
}

type handler struct{ impl Interface }

func (h handler) CompareAndSwap(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args atomic.Store_CompareAndSwap_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.CompareAndSwap(ctx, args.Request)

	hadError := err != nil
	result, err := atomic.Store_CompareAndSwap_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) Forget(ctx context.Context, body wire.Value) error {
	var args atomic.Store_Forget_Args
	if err := args.FromWire(body); err != nil {
		return err
	}

	return h.impl.Forget(ctx, args.Key)
}

func (h handler) Increment(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args atomic.Store_Increment_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.Increment(ctx, args.Key, args.Value)

	hadError := err != nil
	result, err := atomic.Store_Increment_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}
