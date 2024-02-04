package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"github.com/fischettij/account-summarizer/internal/handlers"
	"github.com/fischettij/account-summarizer/internal/handlers/mocks"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks -source=summarizer.go

type SummarizerHandlerSuite struct {
	suite.Suite
	mockCtrl *gomock.Controller
	manager  *mocks.MockSummarizerManager
	handler  *handlers.SummarizerHandler
}

func TestSummarizerHandlerSuite(t *testing.T) {
	suite.Run(t, new(SummarizerHandlerSuite))
}

func (suite *SummarizerHandlerSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.manager = mocks.NewMockSummarizerManager(suite.mockCtrl)
	var err error
	suite.handler, err = handlers.NewSummarizerHandler(suite.manager)
	suite.Require().NoError(err)
	gin.SetMode(gin.TestMode)
}

func (suite *SummarizerHandlerSuite) TearDownTest() {
}

func (suite *SummarizerHandlerSuite) TestSummarySend() {
	router := gin.New()
	endpoint := "/summary/send"
	router.POST(endpoint, suite.handler.SendEmail)

	suite.Run("given_an_empty_body_when_post_to_SummarySend_then_return_bad_request", func() {
		emptyRequestBodyMap := map[string]interface{}{}
		jsonBody, err := json.Marshal(emptyRequestBodyMap)
		suite.Require().NoError(err)

		req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(jsonBody))
		suite.Require().NoError(err)

		responseRecorder := httptest.NewRecorder()
		router.ServeHTTP(responseRecorder, req)

		suite.Require().Equal(http.StatusBadRequest, responseRecorder.Code)
	})

	suite.Run("given_a_body_whit_invalid_file_extension_when_post_to_SummarySend_then_return_bad_request", func() {
		emptyRequestBodyMap := map[string]interface{}{
			"file_name": "file.docx",
			"email":     "some@email.com",
		}
		jsonBody, err := json.Marshal(emptyRequestBodyMap)
		suite.Require().NoError(err)

		req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(jsonBody))
		suite.Require().NoError(err)

		responseRecorder := httptest.NewRecorder()
		router.ServeHTTP(responseRecorder, req)

		suite.Require().Equal(http.StatusBadRequest, responseRecorder.Code)

		var responseJSON map[string]interface{}
		err = json.Unmarshal(responseRecorder.Body.Bytes(), &responseJSON)
		suite.Require().NoError(err)
		suite.Require().Contains(responseJSON["error"].(string), "invalid file_name. CSV extension expected")
	})

	suite.Run("given_a_body_whit_invalid_email_when_post_to_SummarySend_then_return_bad_request", func() {
		emptyRequestBodyMap := map[string]interface{}{
			"file_name": "file.csv",
			"email":     "bad_email",
		}
		jsonBody, err := json.Marshal(emptyRequestBodyMap)
		suite.Require().NoError(err)

		req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(jsonBody))
		suite.Require().NoError(err)

		responseRecorder := httptest.NewRecorder()
		router.ServeHTTP(responseRecorder, req)

		suite.Require().Equal(http.StatusBadRequest, responseRecorder.Code)

		var responseJSON map[string]interface{}
		err = json.Unmarshal(responseRecorder.Body.Bytes(), &responseJSON)
		suite.Require().NoError(err)
		suite.Require().Contains(responseJSON["error"].(string), "invalid email format.")
	})

	suite.Run("given_a_valid_body_when_post_to_SummarySend_then_return_no_content", func() {
		fileName := "file.csv"
		email := "lionel.messi@gmail.com"
		emptyRequestBodyMap := map[string]interface{}{
			"file_name": fileName,
			"email":     email,
		}
		jsonBody, err := json.Marshal(emptyRequestBodyMap)
		suite.Require().NoError(err)

		suite.manager.EXPECT().SendSummary(gomock.Any(), email, "file.csv").Return(nil)

		req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(jsonBody))
		suite.Require().NoError(err)

		responseRecorder := httptest.NewRecorder()
		router.ServeHTTP(responseRecorder, req)

		suite.Require().Equal(http.StatusNoContent, responseRecorder.Code)
	})
}
