package endpoint_test

import (
	"bytes"
	"context"
	"github.com/go-errors/errors"
	"github.com/go-kit/kit/log"
	"github.com/golang/mock/gomock"
	"github.com/seagullbird/headr-repoctl/config"
	"github.com/seagullbird/headr-repoctl/endpoint"
	svcmock "github.com/seagullbird/headr-repoctl/service/mock"
	"os"
	"testing"
)

// When testing Set, just make sure its output and that of its internal service are consistent.
func TestSet(t *testing.T) {
	// Mocking Service
	mockctrl := gomock.NewController(t)
	defer mockctrl.Finish()
	mockSvc := svcmock.NewMockService(mockctrl)

	var buf bytes.Buffer
	logger := log.NewLogfmtLogger(&buf)
	endpoints := endpoint.New(mockSvc, logger)

	dummyError := errors.New("dummy error")
	tests := []struct {
		name string
		rets map[string][]interface{}
	}{
		{"No Error", map[string][]interface{}{
			"NewSite":    {nil},
			"DeleteSite": {nil},
			"WritePost":  {nil},
			"RemovePost": {nil},
			"ReadPost":   {"string", nil},
		}},
		{"Dummy Error", map[string][]interface{}{
			"NewSite":    {dummyError},
			"DeleteSite": {dummyError},
			"WritePost":  {dummyError},
			"RemovePost": {dummyError},
			"ReadPost":   {"", dummyError},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set EXPECTS
			mockSvc.EXPECT().NewSite(gomock.Any(), gomock.Any()).Return(tt.rets["NewSite"]...).Times(2)
			mockSvc.EXPECT().DeleteSite(gomock.Any(), gomock.Any()).Return(tt.rets["DeleteSite"]...).Times(2)
			mockSvc.EXPECT().WritePost(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.rets["WritePost"]...).Times(2)
			mockSvc.EXPECT().RemovePost(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.rets["RemovePost"]...).Times(2)
			mockSvc.EXPECT().ReadPost(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.rets["ReadPost"]...).Times(2)

			t.Run("NewSite", clearEnvWrapper(t, func(t *testing.T) {
				ctx := context.Background()
				siteID := uint(1)
				setErr := endpoints.NewSite(ctx, siteID)
				svcErr := mockSvc.NewSite(ctx, siteID)
				if setErr != svcErr {
					t.Fatal(setErr)
				}
			}))
			t.Run("DeleteSite", clearEnvWrapper(t, func(t *testing.T) {
				ctx := context.Background()
				siteID := uint(1)
				setErr := endpoints.DeleteSite(ctx, siteID)
				svcErr := mockSvc.DeleteSite(ctx, siteID)
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
				svcErr := mockSvc.WritePost(ctx, siteID, filename, content)
				if setErr != svcErr {
					t.Fatal(setErr)
				}
			}))
			t.Run("RemovePost", clearEnvWrapper(t, func(t *testing.T) {
				ctx := context.Background()
				siteID := uint(1)
				filename := "filename"
				setErr := endpoints.RemovePost(ctx, siteID, filename)
				svcErr := mockSvc.RemovePost(ctx, siteID, filename)
				if setErr != svcErr {
					t.Fatal("setErr=", setErr, "svcErr=", svcErr)
				}
			}))
			t.Run("ReadPost", clearEnvWrapper(t, func(t *testing.T) {
				ctx := context.Background()
				siteID := uint(1)
				filename := "filename"
				setOutput, setErr := endpoints.ReadPost(ctx, siteID, filename)
				svcOutput, svcErr := mockSvc.ReadPost(ctx, siteID, filename)
				if setOutput != svcOutput || setErr != svcErr {
					t.Fatal(setOutput, setErr)
				}
			}))
		})
	}
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
