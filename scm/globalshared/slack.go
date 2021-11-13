package globalshared

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Bhinneka/candi/tracer"

	"go.uber.org/zap"

	"github.com/Bhinneka/candi/logger"
)

// SlackAttachment model
type SlackAttachment struct {
	Attachments []SlackPayload `json:"attachments"`
}

// SlackPayload model
type SlackPayload struct {
	Text   string       `json:"text"`
	Color  string       `json:"color"`
	Fields []SlackField `json:"fields"`
}

// SlackField model
type SlackField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

// SlackParam : slack notification param
type SlackParam struct {
	Title         string
	Message       string
	OperationName string
	Error         error
}

const (
	successColor = "#36a64f"
	errorColor   = "#f44b42"
)

// SlackSend : send error notification to slack web hook
// requires os environment key values in .env file. eg :
// SLACK_NOTIFIER=true
// SLACK_URL=https://hooks.slack.com/services/SXA232/SADSAVCW321F
func SlackSend(ctx context.Context, param SlackParam) {
	isActive, _ := strconv.ParseBool(os.Getenv("SLACK_NOTIFIER"))
	if !isActive {
		return
	}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				logger.LogEf("%v", r)
			}
		}()

		var message string = fmt.Sprintf("*Trace ID*: ```%s```", tracer.GetTraceID(ctx))
		if strings.TrimSpace(param.Message) != "" {
			message = fmt.Sprintf("*Trace ID*: ```%s```\n*Message*: ```%s```", tracer.GetTraceID(ctx), param.Message)
		}
		messageBody := fmt.Sprintf("*%s *\n\n%s", param.Title, message)

		var slackPayload SlackPayload
		slackPayload.Text = messageBody
		slackPayload.Color = successColor
		if param.Error != nil {
			slackPayload.Color = errorColor
			slackPayload.Text = fmt.Sprintf("%s\n*Error*: ```%s```", messageBody, param.Error.Error())
		}

		hostName, _ := os.Hostname()
		now := time.Now().Format(time.RFC3339)
		slackPayload.Fields = []SlackField{
			{
				Title: "Server",
				Value: hostName,
				Short: true,
			},
			{
				Title: "Environment",
				Value: os.Getenv("ENVIRONMENT"),
				Short: true,
			},
			{
				Title: "Context",
				Value: param.OperationName,
				Short: true,
			},
			{
				Title: "Time",
				Value: now,
				Short: true,
			},
		}

		var slackAttachment SlackAttachment
		slackAttachment.Attachments = append(slackAttachment.Attachments, slackPayload)

		buffer := &bytes.Buffer{}
		encoder := json.NewEncoder(buffer)
		encoder.SetEscapeHTML(true)
		encoder.Encode(slackAttachment)

		url := os.Getenv("SLACK_URL")
		req, _ := http.NewRequest("POST", url, buffer)
		defer req.Body.Close()

		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
		logger.Log(zap.ErrorLevel, string(body), "slack.SendNotification", "send_to_slack")
	}()
}
