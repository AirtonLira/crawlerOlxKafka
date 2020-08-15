package internals

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

// ConnectKafka - Function for conection with the Kafka broker
func ConnectKafka(topico string) (conn *kafka.Conn) {
	configs := initConfigs()
	conn, err := kafka.DialLeader(context.Background(), "tcp", configs.Kafka.Host, topico, configs.Kafka.Partition)
	if err != nil {
		log.Panicln("Error for connect to Kafka - error-> ", err)
	}
	log.Println("Connected with success kafka - ", configs.Kafka.Host)
	return conn
}

// WriteMessageKafka - Function for write a message in Kafka topic
func WriteMessageKafka(conn *kafka.Conn, msg []byte) {
	_, err := conn.WriteMessages(
		kafka.Message{Value: msg},
	)
	if err != nil {
		log.Panicln("Error to write a message in topic - error-> ", err)
	}
}

// ReaderKafka - Function for read a message in Kafka topic
func ReaderKafka(conn *kafka.Conn, topic string) *kafka.Reader {
	configs := initConfigs()

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{configs.Kafka.Host},
		Topic:     topic,
		Partition: configs.Kafka.Partition,
		MinBytes:  configs.Kafka.MinBytes, // 10KB
		MaxBytes:  10e6,                   // 10MB
	})
	r.SetOffset(1)

	return r
}
