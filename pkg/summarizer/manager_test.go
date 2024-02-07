package summarizer_test

import (
	"context"
	"errors"
	"github.com/fischettij/account-summarizer/pkg/summarizer"
	"github.com/fischettij/account-summarizer/pkg/summarizer/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks -source=manager.go

type ManagerSuite struct {
	suite.Suite
	mockCtrl *gomock.Controller
}

func TestHandlersSuite(t *testing.T) {
	suite.Run(t, new(ManagerSuite))
}

func (suite *ManagerSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
}

func (suite *ManagerSuite) TearDownTest() {
}

func (suite *ManagerSuite) TestNew() {
	suite.Run("given_a_summarizer_and_email_sender_when_creates_new_manager_then_return_manager_and_no_error", func() {
		manager, err := summarizer.NewManager(mocks.NewMockSummarizer(suite.mockCtrl), mocks.NewMockEmailSender(suite.mockCtrl))
		suite.Require().NoError(err)
		suite.Require().NotNil(manager)
	})

	suite.Run("given_an_email_sender_and_nil_summarizer_when_creates_new_manager_then_return_error", func() {
		manager, err := summarizer.NewManager(nil, mocks.NewMockEmailSender(suite.mockCtrl))
		suite.Require().Error(err)
		suite.Require().ErrorContains(err, "summarizer can not be nil")
		suite.Require().Nil(manager)
	})

	suite.Run("given_a_summarizer_and_nil_email_sender_when_creates_new_manager_then_return_error", func() {
		manager, err := summarizer.NewManager(mocks.NewMockSummarizer(suite.mockCtrl), nil)
		suite.Require().Error(err)
		suite.Require().ErrorContains(err, "emailSender can not be nil")
		suite.Require().Nil(manager)
	})
}

func (suite *ManagerSuite) TestStart() {
	suite.Run("given_a_valid_email_and_valid_file_name_when_SendSummary_then_return_no_error", func() {
		fileName := "file_test.csv"
		emailTo := "some@email.com"
		fileSummarizer := mocks.NewMockSummarizer(suite.mockCtrl)
		emailSender := mocks.NewMockEmailSender(suite.mockCtrl)
		summary := mocks.NewMockSummary(suite.mockCtrl)

		manager, err := summarizer.NewManager(fileSummarizer, emailSender)
		suite.Require().NoError(err)
		suite.Require().NotNil(manager)

		summary.EXPECT().Balance().Return(1.0)
		summary.EXPECT().AverageDebitAmount().Return(0.0)
		summary.EXPECT().AverageCreditAmount().Return(1.0)
		summary.EXPECT().TransactionsByMonth().Return(map[time.Month]int{}).AnyTimes()

		fileSummarizer.EXPECT().GenerateReport(fileName).Return(summary, nil)
		emailSender.EXPECT().SendEmail(emailTo, gomock.Any()).Return(nil)

		err = manager.SendSummary(context.Background(), emailTo, fileName)
		suite.Require().NoError(err)
	})

	suite.Run("given_a_valid_email_and_file_with_issues_when_file_summary_processor_fails_then_return_error", func() {
		fileName := "file_test.csv"
		emailTo := "some@email.com"
		fileSummarizer := mocks.NewMockSummarizer(suite.mockCtrl)
		emailSender := mocks.NewMockEmailSender(suite.mockCtrl)

		manager, err := summarizer.NewManager(fileSummarizer, emailSender)
		suite.Require().NoError(err)
		suite.Require().NotNil(manager)

		fileSummarizer.EXPECT().GenerateReport(fileName).Return(nil, errors.New("some-error"))

		err = manager.SendSummary(context.Background(), emailTo, fileName)
		suite.Require().Error(err)
	})

	suite.Run("given_a_email__and_file_when_email_sender_fails_then_return_error", func() {
		fileName := "file_test.csv"
		emailTo := "some@email.com"
		fileSummarizer := mocks.NewMockSummarizer(suite.mockCtrl)
		emailSender := mocks.NewMockEmailSender(suite.mockCtrl)
		summary := mocks.NewMockSummary(suite.mockCtrl)

		manager, err := summarizer.NewManager(fileSummarizer, emailSender)
		suite.Require().NoError(err)
		suite.Require().NotNil(manager)

		summary.EXPECT().Balance().Return(1.0)
		summary.EXPECT().AverageDebitAmount().Return(0.0)
		summary.EXPECT().AverageCreditAmount().Return(1.0)
		summary.EXPECT().TransactionsByMonth().Return(map[time.Month]int{}).AnyTimes()

		fileSummarizer.EXPECT().GenerateReport(fileName).Return(summary, nil)
		emailSender.EXPECT().SendEmail(emailTo, gomock.Any()).Return(errors.New("some-error"))

		err = manager.SendSummary(context.Background(), emailTo, fileName)
		suite.Require().Error(err)
	})
}
