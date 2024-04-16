package backendmanager

import (
	"fmt"

	"gorm.io/gorm"
)

type QueueManager struct {
	DB *gorm.DB
}

type Status string

const (
	QUEUED     Status = "QUEUED"
	RUNNING    Status = "RUNNING"
	COMPLETE   Status = "COMPLETE"
	FAILED     Status = "INCOMPLETE"
	SUCCESSFUL Status = "SUCCESSFUL"
)

type Queue struct {
	gorm.Model
	Username   string `json:"username"`
	StartTime  int64  `json:"startTime"`
	EndTime    int64  `json:"endTime"`
	UseProxies bool   `json:"useProxies"`
	Status     Status `json:"status"`
}

func NewQueueManager() *QueueManager {
	return &QueueManager{}
}

func (qm *QueueManager) CreateQueue(queue Queue) error {
	fmt.Println("Creating queue")
	fmt.Println(queue)

	if err := qm.DB.Create(&queue).Error; err != nil {
		return err
	}
	return nil
}

func (qm *QueueManager) GetQueues() ([]Queue, error) {
	var queues []Queue
	if err := qm.DB.Find(&queues).Error; err != nil {
		return nil, err
	}
	return queues, nil
}

func (qm *QueueManager) DeleteQueue(username string) error {
	if err := qm.DB.Where("username = ?", username).Delete(&Queue{}).Error; err != nil {
		return err
	}
	return nil
}