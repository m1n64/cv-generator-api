package utils

import (
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"log"
	"os"
	"sync"
	"time"
)

const (
	AnalyticQueueName = "analytics_queue"
	DeleteCVQueueName = "delete_cv_queue"
	PDFGenerateQueue  = "pdf_generate_queue"
)

type CVAnalyticQueueMessage struct {
	UserID   uuid.UUID `json:"user_id"`
	CvID     uuid.UUID `json:"cv_id"`
	Action   string    `json:"action"`
	DateTime time.Time `json:"date_time"`
	Detail   string    `json:"detail"`
}

type RabbitMQConnection struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

var instance *RabbitMQConnection
var once sync.Once

func ConnectRabbitMQ() *RabbitMQConnection {
	once.Do(func() {
		conn, err := amqp.Dial(os.Getenv("RABBITMQ_HOST"))
		if err != nil {
			log.Fatalf("Failed to connect to RabbitMQ: %s", err)
		}

		ch, err := conn.Channel()
		if err != nil {
			log.Fatalf("Failed to open a channel: %s", err)
		}

		instance = &RabbitMQConnection{
			Connection: conn,
			Channel:    ch,
		}
		log.Println("RabbitMQ connection established")
	})
	return instance
}

func InitializeQueues() {
	queues := []string{
		AnalyticQueueName,
		DeleteCVQueueName,
		PDFGenerateQueue,
	}

	for _, queue := range queues {
		_, err := instance.Channel.QueueDeclare(
			queue,
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			log.Fatalf("Failed to declare queue %s: %s", queue, err)
		}
		log.Printf("Queue declared: %s", queue)
	}
}

func GetRabbitMQInstance() *RabbitMQConnection {
	if instance == nil {
		log.Fatalf("RabbitMQ connection is not initialized. Call ConnectRabbitMQ first.")
	}
	return instance
}

func (r *RabbitMQConnection) CloseRabbitMQ() {
	if r.Channel != nil {
		r.Channel.Close()
	}
	if r.Connection != nil {
		r.Connection.Close()
	}
	log.Println("RabbitMQ connection closed")
}
