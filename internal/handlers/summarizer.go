package handlers

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SummarizerManager interface {
	SendSummary(ctx context.Context, email, filaPath string) error
}

type SummarizerHandler struct {
	manager SummarizerManager
}

func NewSummarizerHandler(manager SummarizerManager) (*SummarizerHandler, error) {
	if manager == nil {
		return nil, errors.New("summarizer manager cannot be nil")
	}
	return &SummarizerHandler{
		manager: manager,
	}, nil
}

type SendSummaryRequestBody struct {
	Email    string `json:"email"`
	FileName string `json:"file_name"`
}

func (d SummarizerHandler) SendEmail(c *gin.Context) {
	var requestBody SendSummaryRequestBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !isCSVFilePath(requestBody.FileName) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid file_name. CSV extension expected"})
		return
	}

	if !isValidEmail(requestBody.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email format."})
		return
	}

	err := d.manager.SendSummary(context.Background(), requestBody.Email, requestBody.FileName)

	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}
