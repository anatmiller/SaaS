package request_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"
	"time"

	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/client"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/client/metadata"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	v4 "k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/signer/v4"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/awstesting/unit"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/private/protocol/jsonrpc"
)

func BenchmarkTimeoutReadCloser(b *testing.B) {
	resp := `
	{
		"Bar": "qux"
	}
	`

	handlers := request.Handlers{}

	handlers.Send.PushBack(func(r *request.Request) {
		r.HTTPResponse = &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer([]byte(resp))),
		}
	})
	handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	op := &request.Operation{
		Name:       "op",
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	meta := metadata.ClientInfo{
		ServiceName:   "fooService",
		SigningName:   "foo",
		SigningRegion: "foo",
		Endpoint:      "localhost",
		APIVersion:    "2001-01-01",
		JSONVersion:   "1.1",
		TargetPrefix:  "Foo",
	}

	req := request.New(
		*unit.Session.Config,
		meta,
		handlers,
		client.DefaultRetryer{NumMaxRetries: 5},
		op,
		&struct {
			Foo *string
		}{},
		&struct {
			Bar *string
		}{},
	)

	req.ApplyOptions(request.WithResponseReadTimeout(15 * time.Second))
	for i := 0; i < b.N; i++ {
		err := req.Send()
		if err != nil {
			b.Errorf("Expected no error, but received %v", err)
		}
	}
}
