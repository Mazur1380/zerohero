package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Создаем новый читатель Kafka
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"}, // Укажите адрес вашего Kafka брокера
		Topic:     "test",                     // Укажите ваш топик
		Partition: 0,
	})

	defer r.Close()

	// Чтение сообщений
	fmt.Println("Начинаем чтение сообщений...")
	for {
		// Читаем сообщение
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("Ошибка чтения сообщения: %v", err)
		}
		// Обрабатываем сообщение
		fmt.Printf("Сообщение получено: ключ=%s, значение=%s, раздел=%d, смещение=%d\n",
			string(m.Key), string(m.Value), m.Partition, m.Offset)
	}
}

// package main

// import (
// 	"fmt"
// 	"log"

// 	"github.com/confluentinc/confluent-kafka-go/kafka"
// )

// const (
// 	KafkaServer  = "localhost:9092"
// 	KafkaTopic   = "test"
// 	KafkaGroupId = "grupa1"
// )

// func main() {
// 	c, err := kafka.NewConsumer(&kafka.ConfigMap{
// 		"bootstrap.servers": KafkaServer,
// 		"group.id":          "",
// 		"auto.offset.reset": kafka.OffsetBeginning,
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer c.Close()

// 	topic := KafkaTopic
// 	c.Assign([]kafka.TopicPartition{{Topic: &topic, Partition: int32(0), Offset: kafka.OffsetBeginning}})

// 	// c.SubscribeTopics([]string{topic}, nil)

// 	for {
// 		msg, err := c.ReadMessage(-1)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println(string(msg.Value))

// 	}
// }
