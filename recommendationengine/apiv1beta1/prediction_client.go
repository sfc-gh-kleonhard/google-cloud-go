// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package recommendationengine

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	recommendationenginepb "google.golang.org/genproto/googleapis/cloud/recommendationengine/v1beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newPredictionClientHook clientHook

// PredictionCallOptions contains the retry settings for each method of PredictionClient.
type PredictionCallOptions struct {
	Predict []gax.CallOption
}

func defaultPredictionGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("recommendationengine.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("recommendationengine.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://recommendationengine.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultPredictionCallOptions() *PredictionCallOptions {
	return &PredictionCallOptions{
		Predict: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalPredictionClient is an interface that defines the methods available from Recommendations AI.
type internalPredictionClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	Predict(context.Context, *recommendationenginepb.PredictRequest, ...gax.CallOption) *PredictResponse_PredictionResultIterator
}

// PredictionClient is a client for interacting with Recommendations AI.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service for making recommendation prediction.
type PredictionClient struct {
	// The internal transport-dependent client.
	internalClient internalPredictionClient

	// The call options for this service.
	CallOptions *PredictionCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *PredictionClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *PredictionClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *PredictionClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// Predict makes a recommendation prediction. If using API Key based authentication,
// the API Key must be registered using the
// PredictionApiKeyRegistry
// service. Learn more (at /recommendations-ai/docs/setting-up#register-key).
func (c *PredictionClient) Predict(ctx context.Context, req *recommendationenginepb.PredictRequest, opts ...gax.CallOption) *PredictResponse_PredictionResultIterator {
	return c.internalClient.Predict(ctx, req, opts...)
}

// predictionGRPCClient is a client for interacting with Recommendations AI over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type predictionGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing PredictionClient
	CallOptions **PredictionCallOptions

	// The gRPC API client.
	predictionClient recommendationenginepb.PredictionServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewPredictionClient creates a new prediction service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service for making recommendation prediction.
func NewPredictionClient(ctx context.Context, opts ...option.ClientOption) (*PredictionClient, error) {
	clientOpts := defaultPredictionGRPCClientOptions()
	if newPredictionClientHook != nil {
		hookOpts, err := newPredictionClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := PredictionClient{CallOptions: defaultPredictionCallOptions()}

	c := &predictionGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		predictionClient: recommendationenginepb.NewPredictionServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *predictionGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *predictionGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *predictionGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *predictionGRPCClient) Predict(ctx context.Context, req *recommendationenginepb.PredictRequest, opts ...gax.CallOption) *PredictResponse_PredictionResultIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).Predict[0:len((*c.CallOptions).Predict):len((*c.CallOptions).Predict)], opts...)
	it := &PredictResponse_PredictionResultIterator{}
	req = proto.Clone(req).(*recommendationenginepb.PredictRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*recommendationenginepb.PredictResponse_PredictionResult, string, error) {
		resp := &recommendationenginepb.PredictResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.predictionClient.Predict(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetResults(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// PredictResponse_PredictionResultIterator manages a stream of *recommendationenginepb.PredictResponse_PredictionResult.
type PredictResponse_PredictionResultIterator struct {
	items    []*recommendationenginepb.PredictResponse_PredictionResult
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*recommendationenginepb.PredictResponse_PredictionResult, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *PredictResponse_PredictionResultIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *PredictResponse_PredictionResultIterator) Next() (*recommendationenginepb.PredictResponse_PredictionResult, error) {
	var item *recommendationenginepb.PredictResponse_PredictionResult
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *PredictResponse_PredictionResultIterator) bufLen() int {
	return len(it.items)
}

func (it *PredictResponse_PredictionResultIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
