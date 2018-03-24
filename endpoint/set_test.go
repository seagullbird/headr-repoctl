package endpoint_test

import (
	"bytes"
	"context"
	"github.com/go-kit/kit/log"
	"github.com/golang/mock/gomock"
	mqdispatchmock "github.com/seagullbird/headr-common/mq/dispatch/mock"
	"github.com/seagullbird/headr-repoctl/config"
	"github.com/seagullbird/headr-repoctl/endpoint"
	"github.com/seagullbird/headr-repoctl/service"
	"os"
	"testing"
)

// When testing Set, just make sure its output and that of its internal service are consistent.
func TestSet(t *testing.T) {
	mockctrl := gomock.NewController(t)
	defer mockctrl.Finish()
	mockDispatcher := mqdispatchmock.NewMockDispatcher(mockctrl)
	mockDispatcher.EXPECT().DispatchMessage(gomock.Any(), gomock.Any()).Return(nil).Times(4)

	var buf bytes.Buffer
	logger := log.NewLogfmtLogger(&buf)
	svc := service.New(mockDispatcher, logger)

	endpoints := endpoint.New(svc, logger)
	t.Run("NewSite", clearEnvWrapper(t, func(t *testing.T) {
		ctx := context.Background()
		siteID := uint(1)
		setErr := endpoints.NewSite(ctx, siteID)
		svcErr := svc.NewSite(ctx, siteID)
		if setErr != svcErr {
			t.Fatal(setErr)
		}
	}))
	t.Run("DeleteSite", clearEnvWrapper(t, func(t *testing.T) {
		ctx := context.Background()
		siteID := uint(1)
		setErr := endpoints.DeleteSite(ctx, siteID)
		svcErr := svc.DeleteSite(ctx, siteID)
		if setErr != svcErr {
			t.Fatal(setErr)
		}
	}))
	t.Run("WritePost", clearEnvWrapper(t, func(t *testing.T) {
		ctx := context.Background()
		siteID := uint(1)
		filename := "filename"
		content := "content"
		setErr := endpoints.WritePost(ctx, siteID, filename, content)
		svcErr := svc.WritePost(ctx, siteID, filename, content)
		if setErr != svcErr {
			t.Fatal(setErr)
		}
	}))
	t.Run("RemovePost", clearEnvWrapper(t, func(t *testing.T) {
		ctx := context.Background()
		siteID := uint(1)
		filename := "filename"
		setErr := endpoints.RemovePost(ctx, siteID, filename)
		svcErr := svc.RemovePost(ctx, siteID, filename)
		if setErr != svcErr {
			t.Fatal("setErr=", setErr, "svcErr=", svcErr)
		}
	}))
	t.Run("ReadPost", clearEnvWrapper(t, func(t *testing.T) {
		ctx := context.Background()
		siteID := uint(1)
		filename := "filename"
		setOutput, setErr := endpoints.ReadPost(ctx, siteID, filename)
		svcOutput, svcErr := svc.ReadPost(ctx, siteID, filename)
		if setOutput != svcOutput || setErr != svcErr {
			t.Fatal(setOutput, setErr)
		}
	}))
}

func clearEnvWrapper(t *testing.T, tester func(t *testing.T)) func(t *testing.T) {
	if err := os.RemoveAll(config.SITESDIR); !(err == nil || os.IsNotExist(err)) {
		t.Fatalf("Removing SITESDIR failed: %v", err)
	}

	if err := os.MkdirAll(config.SITESDIR, 0644); err != nil {
		t.Fatalf("Creating SITESDIR failed: %v", err)
	}
	return tester
}
