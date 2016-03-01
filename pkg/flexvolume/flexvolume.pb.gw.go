// Code generated by protoc-gen-grpc-gateway
// source: pkg/flexvolume/flexvolume.proto
// DO NOT EDIT!

package flexvolume

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gengo/grpc-gateway/runtime"
	"github.com/gengo/grpc-gateway/utilities"
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"go.pedge.io/pb/go/google/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var _ codes.Code
var _ io.Reader
var _ = runtime.String
var _ = json.Marshal
var _ = utilities.NewDoubleArray

func request_API_Init_0(ctx context.Context, client APIClient, req *http.Request, pathParams map[string]string) (proto.Message, error) {
	var protoReq google_protobuf.Empty

	if err := json.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	return client.Init(ctx, &protoReq)
}

func request_API_Attach_0(ctx context.Context, client APIClient, req *http.Request, pathParams map[string]string) (proto.Message, error) {
	var protoReq AttachRequest

	if err := json.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	return client.Attach(ctx, &protoReq)
}

func request_API_Detach_0(ctx context.Context, client APIClient, req *http.Request, pathParams map[string]string) (proto.Message, error) {
	var protoReq DetachRequest

	if err := json.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	return client.Detach(ctx, &protoReq)
}

func request_API_Mount_0(ctx context.Context, client APIClient, req *http.Request, pathParams map[string]string) (proto.Message, error) {
	var protoReq MountRequest

	if err := json.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	return client.Mount(ctx, &protoReq)
}

func request_API_Unmount_0(ctx context.Context, client APIClient, req *http.Request, pathParams map[string]string) (proto.Message, error) {
	var protoReq UnmountRequest

	if err := json.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	return client.Unmount(ctx, &protoReq)
}

// RegisterAPIHandlerFromEndpoint is same as RegisterAPIHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterAPIHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				glog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				glog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterAPIHandler(ctx, mux, conn)
}

// RegisterAPIHandler registers the http handlers for service API to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterAPIHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	client := NewAPIClient(conn)

	mux.Handle("POST", pattern_API_Init_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		closeNotifier, ok := w.(http.CloseNotifier)
		if ok {
			go func() {
				<-closeNotifier.CloseNotify()
				cancel()
			}()
		}
		resp, err := request_API_Init_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, w, req, err)
			return
		}

		forward_API_Init_0(ctx, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_API_Attach_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		closeNotifier, ok := w.(http.CloseNotifier)
		if ok {
			go func() {
				<-closeNotifier.CloseNotify()
				cancel()
			}()
		}
		resp, err := request_API_Attach_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, w, req, err)
			return
		}

		forward_API_Attach_0(ctx, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_API_Detach_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		closeNotifier, ok := w.(http.CloseNotifier)
		if ok {
			go func() {
				<-closeNotifier.CloseNotify()
				cancel()
			}()
		}
		resp, err := request_API_Detach_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, w, req, err)
			return
		}

		forward_API_Detach_0(ctx, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_API_Mount_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		closeNotifier, ok := w.(http.CloseNotifier)
		if ok {
			go func() {
				<-closeNotifier.CloseNotify()
				cancel()
			}()
		}
		resp, err := request_API_Mount_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, w, req, err)
			return
		}

		forward_API_Mount_0(ctx, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_API_Unmount_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		closeNotifier, ok := w.(http.CloseNotifier)
		if ok {
			go func() {
				<-closeNotifier.CloseNotify()
				cancel()
			}()
		}
		resp, err := request_API_Unmount_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, w, req, err)
			return
		}

		forward_API_Unmount_0(ctx, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_API_Init_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"init"}, ""))

	pattern_API_Attach_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"attach"}, ""))

	pattern_API_Detach_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"detach"}, ""))

	pattern_API_Mount_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"mount"}, ""))

	pattern_API_Unmount_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"unmount"}, ""))
)

var (
	forward_API_Init_0 = runtime.ForwardResponseMessage

	forward_API_Attach_0 = runtime.ForwardResponseMessage

	forward_API_Detach_0 = runtime.ForwardResponseMessage

	forward_API_Mount_0 = runtime.ForwardResponseMessage

	forward_API_Unmount_0 = runtime.ForwardResponseMessage
)
