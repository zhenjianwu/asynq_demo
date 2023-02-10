package common

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"log"
	"time"
)

const (
	TypeEmailDelivery = "email:deliver"
	TypeEcho          = "email:echo"
)

// EmailDeliveryPayload 异步任务需要传递的数据结构
type EmailDeliveryPayload struct {
	UserID     int
	TemplateID string
	DataStr    string
}

type EchoPayload struct {
	UserID     int
	TemplateID string
	DataStr    string
}

// NewEmailDeliveryTask 异步任务需要传递的数据
func NewEmailDeliveryTask(userID int, tmplID, dataStr string) (*asynq.Task, error) {
	payload, err := json.Marshal(EmailDeliveryPayload{UserID: userID, TemplateID: tmplID, DataStr: dataStr})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return asynq.NewTask(TypeEmailDelivery, payload), nil
}

func NewEchoTask(userID int, tmplId, dataStr string) (*asynq.Task, error) {
	payload, err := json.Marshal(EmailDeliveryPayload{UserID: userID, TemplateID: tmplId, DataStr: dataStr})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return asynq.NewTask(TypeEcho, payload), nil
}

// HandleEmailDeliveryTask 发送email处理逻辑
func HandleEmailDeliveryTask(ctx context.Context, t *asynq.Task) error {
	//接收任务数据
	var p EmailDeliveryPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	//逻辑处理start...
	log.Printf("Sending Email to User: user_id=%d, template_id=%s data_str:%s", p.UserID, p.TemplateID, p.DataStr)
	time.Sleep(3 * time.Second)
	return nil
}

func HandleEchoTask(ctx context.Context, t *asynq.Task) error {
	var p EchoPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	//逻辑处理start...
	log.Printf("Sending echo to User: user_id=%d, template_id=%s data_str:%s", p.UserID, p.TemplateID, p.DataStr)
	time.Sleep(3 * time.Second)
	return nil
}
