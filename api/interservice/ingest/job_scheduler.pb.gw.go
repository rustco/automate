// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: interservice/ingest/job_scheduler.proto

/*
Package ingest is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package ingest

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = descriptor.ForMessage
var _ = metadata.Join

func request_JobSchedulerService_GetStatusJobScheduler_0(ctx context.Context, marshaler runtime.Marshaler, client JobSchedulerServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq JobSchedulerStatusRequest
	var metadata runtime.ServerMetadata

	msg, err := client.GetStatusJobScheduler(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_JobSchedulerService_GetStatusJobScheduler_0(ctx context.Context, marshaler runtime.Marshaler, server JobSchedulerServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq JobSchedulerStatusRequest
	var metadata runtime.ServerMetadata

	msg, err := server.GetStatusJobScheduler(ctx, &protoReq)
	return msg, metadata, err

}

func request_JobSchedulerService_MarkNodesMissing_0(ctx context.Context, marshaler runtime.Marshaler, client JobSchedulerServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq MarkNodesMissingRequest
	var metadata runtime.ServerMetadata

	msg, err := client.MarkNodesMissing(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_JobSchedulerService_MarkNodesMissing_0(ctx context.Context, marshaler runtime.Marshaler, server JobSchedulerServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq MarkNodesMissingRequest
	var metadata runtime.ServerMetadata

	msg, err := server.MarkNodesMissing(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_JobSchedulerService_ConfigureNodesMissingScheduler_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_JobSchedulerService_ConfigureNodesMissingScheduler_0(ctx context.Context, marshaler runtime.Marshaler, client JobSchedulerServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq JobSettings
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_JobSchedulerService_ConfigureNodesMissingScheduler_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.ConfigureNodesMissingScheduler(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_JobSchedulerService_ConfigureNodesMissingScheduler_0(ctx context.Context, marshaler runtime.Marshaler, server JobSchedulerServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq JobSettings
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_JobSchedulerService_ConfigureNodesMissingScheduler_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.ConfigureNodesMissingScheduler(ctx, &protoReq)
	return msg, metadata, err

}

func request_JobSchedulerService_StartNodesMissingScheduler_0(ctx context.Context, marshaler runtime.Marshaler, client JobSchedulerServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq StartNodesMissingSchedulerRequest
	var metadata runtime.ServerMetadata

	msg, err := client.StartNodesMissingScheduler(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_JobSchedulerService_StartNodesMissingScheduler_0(ctx context.Context, marshaler runtime.Marshaler, server JobSchedulerServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq StartNodesMissingSchedulerRequest
	var metadata runtime.ServerMetadata

	msg, err := server.StartNodesMissingScheduler(ctx, &protoReq)
	return msg, metadata, err

}

func request_JobSchedulerService_StopNodesMissingScheduler_0(ctx context.Context, marshaler runtime.Marshaler, client JobSchedulerServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq StopNodesMissingSchedulerRequest
	var metadata runtime.ServerMetadata

	msg, err := client.StopNodesMissingScheduler(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_JobSchedulerService_StopNodesMissingScheduler_0(ctx context.Context, marshaler runtime.Marshaler, server JobSchedulerServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq StopNodesMissingSchedulerRequest
	var metadata runtime.ServerMetadata

	msg, err := server.StopNodesMissingScheduler(ctx, &protoReq)
	return msg, metadata, err

}

func request_JobSchedulerService_DeleteMarkedNodes_0(ctx context.Context, marshaler runtime.Marshaler, client JobSchedulerServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq DeleteMarkedNodesRequest
	var metadata runtime.ServerMetadata

	msg, err := client.DeleteMarkedNodes(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_JobSchedulerService_DeleteMarkedNodes_0(ctx context.Context, marshaler runtime.Marshaler, server JobSchedulerServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq DeleteMarkedNodesRequest
	var metadata runtime.ServerMetadata

	msg, err := server.DeleteMarkedNodes(ctx, &protoReq)
	return msg, metadata, err

}

func request_JobSchedulerService_StartDeleteNodesScheduler_0(ctx context.Context, marshaler runtime.Marshaler, client JobSchedulerServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq StartDeleteNodesSchedulerRequest
	var metadata runtime.ServerMetadata

	msg, err := client.StartDeleteNodesScheduler(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_JobSchedulerService_StartDeleteNodesScheduler_0(ctx context.Context, marshaler runtime.Marshaler, server JobSchedulerServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq StartDeleteNodesSchedulerRequest
	var metadata runtime.ServerMetadata

	msg, err := server.StartDeleteNodesScheduler(ctx, &protoReq)
	return msg, metadata, err

}

func request_JobSchedulerService_StopDeleteNodesScheduler_0(ctx context.Context, marshaler runtime.Marshaler, client JobSchedulerServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq StopDeleteNodesSchedulerRequest
	var metadata runtime.ServerMetadata

	msg, err := client.StopDeleteNodesScheduler(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_JobSchedulerService_StopDeleteNodesScheduler_0(ctx context.Context, marshaler runtime.Marshaler, server JobSchedulerServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq StopDeleteNodesSchedulerRequest
	var metadata runtime.ServerMetadata

	msg, err := server.StopDeleteNodesScheduler(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_JobSchedulerService_ConfigureDeleteNodesScheduler_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_JobSchedulerService_ConfigureDeleteNodesScheduler_0(ctx context.Context, marshaler runtime.Marshaler, client JobSchedulerServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq JobSettings
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_JobSchedulerService_ConfigureDeleteNodesScheduler_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.ConfigureDeleteNodesScheduler(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_JobSchedulerService_ConfigureDeleteNodesScheduler_0(ctx context.Context, marshaler runtime.Marshaler, server JobSchedulerServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq JobSettings
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_JobSchedulerService_ConfigureDeleteNodesScheduler_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.ConfigureDeleteNodesScheduler(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterJobSchedulerServiceHandlerServer registers the http handlers for service JobSchedulerService to "mux".
// UnaryRPC     :call JobSchedulerServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterJobSchedulerServiceHandlerFromEndpoint instead.
func RegisterJobSchedulerServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server JobSchedulerServiceServer) error {

	mux.Handle("GET", pattern_JobSchedulerService_GetStatusJobScheduler_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_JobSchedulerService_GetStatusJobScheduler_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_JobSchedulerService_GetStatusJobScheduler_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_JobSchedulerService_MarkNodesMissing_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_JobSchedulerService_MarkNodesMissing_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_JobSchedulerService_MarkNodesMissing_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_JobSchedulerService_ConfigureNodesMissingScheduler_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_JobSchedulerService_ConfigureNodesMissingScheduler_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_JobSchedulerService_ConfigureNodesMissingScheduler_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_JobSchedulerService_StartNodesMissingScheduler_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_JobSchedulerService_StartNodesMissingScheduler_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_JobSchedulerService_StartNodesMissingScheduler_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_JobSchedulerService_StopNodesMissingScheduler_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_JobSchedulerService_StopNodesMissingScheduler_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_JobSchedulerService_StopNodesMissingScheduler_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_JobSchedulerService_DeleteMarkedNodes_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_JobSchedulerService_DeleteMarkedNodes_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_JobSchedulerService_DeleteMarkedNodes_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_JobSchedulerService_StartDeleteNodesScheduler_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_JobSchedulerService_StartDeleteNodesScheduler_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_JobSchedulerService_StartDeleteNodesScheduler_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_JobSchedulerService_StopDeleteNodesScheduler_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_JobSchedulerService_StopDeleteNodesScheduler_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_JobSchedulerService_StopDeleteNodesScheduler_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_JobSchedulerService_ConfigureDeleteNodesScheduler_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_JobSchedulerService_ConfigureDeleteNodesScheduler_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_JobSchedulerService_ConfigureDeleteNodesScheduler_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterJobSchedulerServiceHandlerFromEndpoint is same as RegisterJobSchedulerServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterJobSchedulerServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterJobSchedulerServiceHandler(ctx, mux, conn)
}

// RegisterJobSchedulerServiceHandler registers the http handlers for service JobSchedulerService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterJobSchedulerServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterJobSchedulerServiceHandlerClient(ctx, mux, NewJobSchedulerServiceClient(conn))
}

// RegisterJobSchedulerServiceHandlerClient registers the http handlers for service JobSchedulerService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "JobSchedulerServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "JobSchedulerServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "JobSchedulerServiceClient" to call the correct interceptors.
func RegisterJobSchedulerServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client JobSchedulerServiceClient) error {

	mux.Handle("GET", pattern_JobSchedulerService_GetStatusJobScheduler_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_JobSchedulerService_GetStatusJobScheduler_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_JobSchedulerService_GetStatusJobScheduler_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_JobSchedulerService_MarkNodesMissing_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_JobSchedulerService_MarkNodesMissing_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_JobSchedulerService_MarkNodesMissing_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_JobSchedulerService_ConfigureNodesMissingScheduler_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_JobSchedulerService_ConfigureNodesMissingScheduler_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_JobSchedulerService_ConfigureNodesMissingScheduler_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_JobSchedulerService_StartNodesMissingScheduler_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_JobSchedulerService_StartNodesMissingScheduler_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_JobSchedulerService_StartNodesMissingScheduler_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_JobSchedulerService_StopNodesMissingScheduler_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_JobSchedulerService_StopNodesMissingScheduler_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_JobSchedulerService_StopNodesMissingScheduler_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_JobSchedulerService_DeleteMarkedNodes_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_JobSchedulerService_DeleteMarkedNodes_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_JobSchedulerService_DeleteMarkedNodes_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_JobSchedulerService_StartDeleteNodesScheduler_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_JobSchedulerService_StartDeleteNodesScheduler_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_JobSchedulerService_StartDeleteNodesScheduler_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_JobSchedulerService_StopDeleteNodesScheduler_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_JobSchedulerService_StopDeleteNodesScheduler_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_JobSchedulerService_StopDeleteNodesScheduler_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_JobSchedulerService_ConfigureDeleteNodesScheduler_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_JobSchedulerService_ConfigureDeleteNodesScheduler_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_JobSchedulerService_ConfigureDeleteNodesScheduler_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_JobSchedulerService_GetStatusJobScheduler_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4}, []string{"api", "v0", "job", "job-scheduler", "status"}, "", runtime.AssumeColonVerbOpt(true)))

	pattern_JobSchedulerService_MarkNodesMissing_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4}, []string{"api", "v0", "job", "missing-nodes", "mark-nodes-missing"}, "", runtime.AssumeColonVerbOpt(true)))

	pattern_JobSchedulerService_ConfigureNodesMissingScheduler_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4}, []string{"api", "v0", "job", "missing-nodes", "config"}, "", runtime.AssumeColonVerbOpt(true)))

	pattern_JobSchedulerService_StartNodesMissingScheduler_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4}, []string{"api", "v0", "job", "missing-nodes", "start"}, "", runtime.AssumeColonVerbOpt(true)))

	pattern_JobSchedulerService_StopNodesMissingScheduler_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4}, []string{"api", "v0", "job", "missing-nodes", "stop"}, "", runtime.AssumeColonVerbOpt(true)))

	pattern_JobSchedulerService_DeleteMarkedNodes_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4}, []string{"api", "v0", "job", "delete-nodes", "delete-marked-nodes"}, "", runtime.AssumeColonVerbOpt(true)))

	pattern_JobSchedulerService_StartDeleteNodesScheduler_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4}, []string{"api", "v0", "job", "delete-nodes", "start"}, "", runtime.AssumeColonVerbOpt(true)))

	pattern_JobSchedulerService_StopDeleteNodesScheduler_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4}, []string{"api", "v0", "job", "delete-nodes", "stop"}, "", runtime.AssumeColonVerbOpt(true)))

	pattern_JobSchedulerService_ConfigureDeleteNodesScheduler_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4}, []string{"api", "v0", "job", "delete-nodes", "config"}, "", runtime.AssumeColonVerbOpt(true)))
)

var (
	forward_JobSchedulerService_GetStatusJobScheduler_0 = runtime.ForwardResponseMessage

	forward_JobSchedulerService_MarkNodesMissing_0 = runtime.ForwardResponseMessage

	forward_JobSchedulerService_ConfigureNodesMissingScheduler_0 = runtime.ForwardResponseMessage

	forward_JobSchedulerService_StartNodesMissingScheduler_0 = runtime.ForwardResponseMessage

	forward_JobSchedulerService_StopNodesMissingScheduler_0 = runtime.ForwardResponseMessage

	forward_JobSchedulerService_DeleteMarkedNodes_0 = runtime.ForwardResponseMessage

	forward_JobSchedulerService_StartDeleteNodesScheduler_0 = runtime.ForwardResponseMessage

	forward_JobSchedulerService_StopDeleteNodesScheduler_0 = runtime.ForwardResponseMessage

	forward_JobSchedulerService_ConfigureDeleteNodesScheduler_0 = runtime.ForwardResponseMessage
)
