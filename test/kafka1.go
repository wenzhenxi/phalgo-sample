package main

import (
	"github.com/Shopify/sarama"
	"os"
	"os/signal"
	"log"
	"fmt"
)

func main() {

	consumer, err := sarama.NewConsumer([]string{"120.27.97.59:9092"}, nil)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	fmt.Println(1)
	partitionConsumer, err := consumer.ConsumePartition("wenzhenxi", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	fmt.Println(1)
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	fmt.Println(1)
	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	fmt.Println(1)
	consumed := 0
	ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Consumed message offset %d\n", msg.Offset)
			consumed++
		case <-signals:
			break ConsumerLoop
		}
	}
	fmt.Println(1)
	log.Printf("Consumed: %d\n", consumed)

}



// func main() {
//	t := time.Now().Unix()
//
//	config := sarama.NewConfig()
//	config.Producer.Return.Successes = true
//	producer, err := sarama.NewAsyncProducer([]string{"120.27.97.59:9092"}, config)
//	if err != nil {
//		panic(err)
//	}
//
//	// Trap SIGINT to trigger a graceful shutdown.
//	signals := make(chan os.Signal, 1)
//	signal.Notify(signals, os.Interrupt)
//
//	var (
//		wg sync.WaitGroup
//		enqueued, successes, errors int
//	)
//
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		for _ = range producer.Successes() {
//			successes++
//		}
//	}()
//
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		for err := range producer.Errors() {
//			log.Println(err)
//			errors++
//		}
//	}()
//	i := 0
//	ProducerLoop:
//	for {
//		i++
//		message := &sarama.ProducerMessage{Topic: "test123", Value: sarama.StringEncoder("testing 123")}
//		select {
//		case producer.Input() <- message:
//			enqueued++
//
//		case <-signals:
//			producer.AsyncClose() // Trigger a shutdown of the producer.
//			break ProducerLoop
//		}
//
//	}
//	fmt.Println(time.Now().Unix() - t)
//	wg.Wait()
//
//	log.Printf("Successfully produced: %d; errors: %d\n", successes, errors)
//}


/*func main() {

	broker := sarama.NewBroker("115.28.153.5:9092")
	err := broker.Open(nil)
	if err != nil {
		panic(err)
	}

	request := sarama.MetadataRequest{Topics: []string{"test"}}
	response, err := broker.GetMetadata(&request)
	if err != nil {
		_ = broker.Close()
		panic(err)
	}

	fmt.Println("There are", len(response.Topics), "topics active in the cluster.")

	if err = broker.Close(); err != nil {
		panic(err)
	}
}*/
func main2() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition("my_topic", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumed := 0
	ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Consumed message offset %d\n", msg.Offset)
			consumed++
		case <-signals:
			break ConsumerLoop
		}
	}

	log.Printf("Consumed: %d\n", consumed)
}

