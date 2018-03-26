package transport_test

import (
	"github.com/golang/mock/gomock"
	svcmock "github.com/seagullbird/headr-repoctl/service/mock"
	"testing"

	"context"
	"errors"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-repoctl/endpoint"
	"github.com/seagullbird/headr-repoctl/pb"
	"github.com/seagullbird/headr-repoctl/service"
	"github.com/seagullbird/headr-repoctl/transport"
	"google.golang.org/grpc"
	"net"
)

const (
	port = ":1234"
)

func startServer(t *testing.T, svc service.Service) {
	logger := log.NewNopLogger()
	endpoints := endpoint.New(svc, logger)
	grpcServer := transport.NewGRPCServer(endpoints, logger)
	grpcListener, err := net.Listen("tcp", port)
	if err != nil {
		t.Fatal(err)
	}
	baseServer := grpc.NewServer()
	pb.RegisterRepoctlServer(baseServer, grpcServer)

	baseServer.Serve(grpcListener)
}

// This Test tests only application level error handling;
// Application error means an error that is returned by the service itself;
// For example, if the client requested information of an unknown user, the service will respond with a
// "User Not Found" error, which is usually customized by the service programmer.
// So this Test mocks the service, which returns two types of response: No Error and Dummy Error;
// I want to make sure that what the client receives is exactly the same as what the service returns. (nil or dummy error)
func TestGRPCApplication(t *testing.T) {
	// Mocking service.Service
	mockctrl := gomock.NewController(t)
	defer mockctrl.Finish()
	mockSvc := svcmock.NewMockService(mockctrl)
	// Set mock service expectations
	dummyError := errors.New("dummy error")
	for _, rets := range []map[string][]interface{}{
		{
			"NewSite":    {nil},
			"DeleteSite": {nil},
			"WritePost":  {nil},
			"RemovePost": {nil},
			"ReadPost":   {"string", nil},
		},
		{
			"NewSite":    {dummyError},
			"DeleteSite": {dummyError},
			"WritePost":  {dummyError},
			"RemovePost": {dummyError},
			"ReadPost":   {"", dummyError},
		},
	} {
		times := 2
		mockSvc.EXPECT().NewSite(gomock.Any(), gomock.Any()).Return(rets["NewSite"]...).Times(times)
		mockSvc.EXPECT().DeleteSite(gomock.Any(), gomock.Any()).Return(rets["DeleteSite"]...).Times(times)
		mockSvc.EXPECT().WritePost(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(rets["WritePost"]...).Times(times)
		mockSvc.EXPECT().RemovePost(gomock.Any(), gomock.Any(), gomock.Any()).Return(rets["RemovePost"]...).Times(times)
		mockSvc.EXPECT().ReadPost(gomock.Any(), gomock.Any(), gomock.Any()).Return(rets["ReadPost"]...).Times(times)
	}
	// Start GRPC server with the mock service
	go startServer(t, mockSvc)

	// Start GRPC client
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}

	client := transport.NewGRPCClient(conn, nil)

	// testcases
	tests := []struct {
		name   string
		judger func(err1, err2 error) bool
	}{
		{
			"No Error",
			func(err1, err2 error) bool {
				if err1 != nil || err2 != nil {
					return false
				}
				return true
			},
		},
		{
			"Dummy Error",
			func(err1, err2 error) bool {
				if err1.Error() != "dummy error" || err2.Error() != "dummy error" {
					return false
				}
				return true
			},
		},
	}

	// Start tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("NewSite", func(t *testing.T) {
				siteID := uint(1)
				ctx := context.Background()
				clientErr := client.NewSite(ctx, siteID)
				svcErr := mockSvc.NewSite(ctx, siteID)
				if !tt.judger(clientErr, svcErr) {
					t.Fatal("\nclientErr: ", clientErr, "\nsvcErr: ", svcErr)
				}
			})
			t.Run("DeleteSite", func(t *testing.T) {
				ctx := context.Background()
				siteID := uint(1)
				clientErr := client.DeleteSite(ctx, siteID)
				svcErr := mockSvc.DeleteSite(ctx, siteID)
				if !tt.judger(clientErr, svcErr) {
					t.Fatal("\nclientErr: ", clientErr, "\nsvcErr: ", svcErr)
				}
			})
			t.Run("WritePost", func(t *testing.T) {
				ctx := context.Background()
				siteID := uint(1)
				filename := "filename"
				content := "content"
				clientErr := client.WritePost(ctx, siteID, filename, content)
				svcErr := mockSvc.WritePost(ctx, siteID, filename, content)
				if !tt.judger(clientErr, svcErr) {
					t.Fatal("\nclientErr: ", clientErr, "\nsvcErr: ", svcErr)
				}
			})
			t.Run("RemovePost", func(t *testing.T) {
				ctx := context.Background()
				siteID := uint(1)
				filename := "filename"
				clientErr := client.RemovePost(ctx, siteID, filename)
				svcErr := mockSvc.RemovePost(ctx, siteID, filename)
				if !tt.judger(clientErr, svcErr) {
					t.Fatal("\nclientErr: ", clientErr, "\nsvcErr: ", svcErr)
				}
			})
			t.Run("ReadPost", func(t *testing.T) {
				ctx := context.Background()
				siteID := uint(1)
				filename := "filename"
				clientOutput, clientErr := client.ReadPost(ctx, siteID, filename)
				svcOutput, svcErr := mockSvc.ReadPost(ctx, siteID, filename)
				if clientOutput != svcOutput || !tt.judger(clientErr, svcErr) {
					t.Fatal("\nclientOutput: ", clientOutput, "\nclientErr: ", clientErr, "\nsvcOutput: ", svcOutput, "\nsvcErr: ", svcErr)
				}
			})
		})
	}
}
