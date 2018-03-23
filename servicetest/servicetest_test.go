package servicetest

import (
	"bytes"
	"github.com/go-kit/kit/log"
	"github.com/golang/mock/gomock"
	mqdispatchmock "github.com/seagullbird/headr-common/mq/dispatch/mock"
	"github.com/seagullbird/headr-repoctl/service"
	"testing"
)

func TestServiceTest(t *testing.T) {
	mockctrl := gomock.NewController(t)
	defer mockctrl.Finish()
	mockDispatcher := mqdispatchmock.NewMockDispatcher(mockctrl)
	mockDispatcher.EXPECT().DispatchMessage(gomock.Any(), gomock.Any()).Return(nil).Times(3)

	var buf bytes.Buffer
	logger := log.NewLogfmtLogger(&buf)
	svc := service.New(mockDispatcher, logger)
	svctest := New(svc, &buf)

	RunTests(t, svctest)
}
