package service_test

import (
	"bytes"
	"github.com/go-kit/kit/log"
	"github.com/golang/mock/gomock"
	mqdispatchmock "github.com/seagullbird/headr-common/mq/dispatch/mock"
	"github.com/seagullbird/headr-repoctl/service"
	"github.com/seagullbird/headr-repoctl/servicetest"
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
	svctest := servicetest.New(svc, &buf)

	servicetest.RunTests(t, svctest)
}
