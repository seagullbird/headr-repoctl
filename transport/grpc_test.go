package transport_test

import (
	"github.com/golang/mock/gomock"
	svcmock "github.com/seagullbird/headr-repoctl/service/mock"
	"testing"

	"context"
	//"errors"
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

func TestGRPC(t *testing.T) {
	mockctrl := gomock.NewController(t)
	defer mockctrl.Finish()
	mockSvc := svcmock.NewMockService(mockctrl)
	// TODO:EXPECTS
	//dummyError := errors.New("dummy error")
	for _, rets := range []map[string][]interface{}{
		map[string][]interface{}{
			"NewSite":    {nil},
			"DeleteSite": {nil},
			"WritePost":  {nil},
			"RemovePost": {nil},
			"ReadPost":   {"string", nil},
		},
		//map[string][]interface{}{
		//	"NewSite":    {dummyError},
		//	"DeleteSite": {dummyError},
		//	"WritePost":  {dummyError},
		//	"RemovePost": {dummyError},
		//	"ReadPost":   {"", dummyError},
		//},
	} {
		times := 2
		mockSvc.EXPECT().NewSite(gomock.Any(), gomock.Any()).Return(rets["NewSite"]...).Times(times)
		mockSvc.EXPECT().DeleteSite(gomock.Any(), gomock.Any()).Return(rets["DeleteSite"]...).Times(times)
		mockSvc.EXPECT().WritePost(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(rets["WritePost"]...).Times(times)
		mockSvc.EXPECT().RemovePost(gomock.Any(), gomock.Any(), gomock.Any()).Return(rets["RemovePost"]...).Times(times)
		mockSvc.EXPECT().ReadPost(gomock.Any(), gomock.Any(), gomock.Any()).Return(rets["ReadPost"]...).Times(times)
	}
	go startServer(t, mockSvc)

	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}

	client := transport.NewGRPCClient(conn, nil)

	tests := []struct{ name string }{
		{"No Error"},
		//{"Dummy Error"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("NewSite", func(t *testing.T) {
				siteID := uint(1)
				ctx := context.Background()
				clientErr := client.NewSite(ctx, siteID)
				svcErr := mockSvc.NewSite(ctx, siteID)
				if clientErr != svcErr {
					t.Fatal(clientErr, svcErr)
				}
			})
			t.Run("DeleteSite", func(t *testing.T) {
				ctx := context.Background()
				siteID := uint(1)
				clientErr := client.DeleteSite(ctx, siteID)
				svcErr := mockSvc.DeleteSite(ctx, siteID)
				if clientErr != svcErr {
					t.Fatal(clientErr)
				}
			})
			t.Run("WritePost", func(t *testing.T) {
				ctx := context.Background()
				siteID := uint(1)
				filename := "filename"
				content := "content"
				clientErr := client.WritePost(ctx, siteID, filename, content)
				svcErr := mockSvc.WritePost(ctx, siteID, filename, content)
				if clientErr != svcErr {
					t.Fatal(clientErr)
				}
			})
			t.Run("RemovePost", func(t *testing.T) {
				ctx := context.Background()
				siteID := uint(1)
				filename := "filename"
				clientErr := client.RemovePost(ctx, siteID, filename)
				svcErr := mockSvc.RemovePost(ctx, siteID, filename)
				if clientErr != svcErr {
					t.Fatal("clientErr=", clientErr, "svcErr=", svcErr)
				}
			})
			t.Run("ReadPost", func(t *testing.T) {
				ctx := context.Background()
				siteID := uint(1)
				filename := "filename"
				setOutput, clientErr := client.ReadPost(ctx, siteID, filename)
				svcOutput, svcErr := mockSvc.ReadPost(ctx, siteID, filename)
				if setOutput != svcOutput || clientErr != svcErr {
					t.Fatal(setOutput, clientErr)
				}
			})
		})
	}
}
