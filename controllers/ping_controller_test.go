package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_services "github.com/drpaneas/drew/mocks"
	"github.com/drpaneas/drew/services"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func TestPingControllerNoError(t *testing.T) {
	// Preparation for the gin.Context input
	fakeResponseWriter := httptest.NewRecorder()
	fakeGinContext, _ := gin.CreateTestContext(fakeResponseWriter)

	// prepare GoMock
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockPingController := mock_services.NewMockpingServiceInterface(mockCtrl)
	// What the service will return
	mockPingController.EXPECT().PingService().Return("pong", nil)
	// hack
	services.PingServiceVar = mockPingController

	// I am ready to test
	PingController(fakeGinContext)

	// First Test
	// Check if status is 200
	if fakeResponseWriter.Code != http.StatusOK {
		t.Error("response code should be 200")
	}

	// Second Test
	// Check if it returns 'pong'
	if fakeResponseWriter.Body.String() != "pong" {
		t.Error("response string should say 'pong'")
	}
}

func TestPingControllerWithError(t *testing.T) {
	// Preparation for the gin.Context input
	fakeResponseWriter := httptest.NewRecorder()
	fakeGinContext, _ := gin.CreateTestContext(fakeResponseWriter)

	// prepare GoMock
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockPingController := mock_services.NewMockpingServiceInterface(mockCtrl)
	err := fmt.Errorf(http.StatusText(http.StatusInternalServerError))
	mockPingController.EXPECT().PingService().Return("", err)

	// hack
	services.PingServiceVar = mockPingController

	// I am ready to test
	PingController(fakeGinContext)

	// First Test
	// Check if status is not 200
	if fakeResponseWriter.Code == http.StatusOK {
		t.Error("response code should not be 200")
	}

	// Second Test
	// Check if it does not return 'pong'
	if fakeResponseWriter.Body.String() == "pong" {
		t.Error("response string should not say 'pong'")
	}
}
