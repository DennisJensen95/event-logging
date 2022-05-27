package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
	"github.com/segmentio/kafka-go"
)

const (
	topic_cpu    = "cpu"
	topic_memory = "memory"
	broker       = "localhost:9092"
	frequency    = 1
)

type publish_func func(*kafka.Writer, context.Context)

func publish_cpu_usage(kafka_writer *kafka.Writer, ctx context.Context) {
	cpu_usage, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	json_payload := "{\"cpu_usage\": " + fmt.Sprintf("%2f", float64(cpu_usage.System+cpu_usage.User)/float64(cpu_usage.Total)*100) +
		", \"cpu_system_usage\": " + fmt.Sprintf("%2f", float64(cpu_usage.System)/float64(cpu_usage.Total)*100) +
		", \"cpu_user_usage\": " + fmt.Sprintf("%2f", float64(cpu_usage.User)/float64(cpu_usage.Total)*100) +
		", \"cpu_cores\": " + fmt.Sprintf("%d", cpu_usage.CPUCount) + "}"

	err_writing := kafka_writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte("cpu"),
		Value: []byte(json_payload),
	})

	if err_writing != nil {
		panic("could not write message " + err.Error())
	}
}

func publish_memory_usage(kafka_writer *kafka.Writer, ctx context.Context) {
	memory_usage, err := memory.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	json_payload := "{\"memory_usage\": " + fmt.Sprintf("%2f", float64(memory_usage.Used)/float64(memory_usage.Total)*100) +
		", \"memory_used\": " + fmt.Sprintf("%2f", float64(memory_usage.Used)/1024/1024) +
		", \"memory_total\": " + fmt.Sprintf("%2f", float64(memory_usage.Total)/1024/1024) + "}"

	err_writing := kafka_writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte("memory"),
		Value: []byte(json_payload),
	})

	if err_writing != nil {
		panic("could not write message " + err.Error())
	}
}

func call_func_on_timer(wg *sync.WaitGroup, fn publish_func, topic string, ctx context.Context) {
	defer wg.Done()

	l := log.New(os.Stdout, "kafka writer: ", 0)

	kafka_writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker},
		Topic:   topic,
		// assign the logger to the writer
		Logger: l,
	})

	defer kafka_writer.Close()

	for range time.Tick(time.Second * frequency) {
		fn(kafka_writer, ctx)
	}
}

func main() {
	var wg sync.WaitGroup
	ctx := context.Background()

	wg.Add(1)
	go call_func_on_timer(&wg, publish_cpu_usage, topic_cpu, ctx)
	wg.Add(1)
	go call_func_on_timer(&wg, publish_memory_usage, topic_memory, ctx)
	fmt.Println("Running CPU usage publisher")
	wg.Wait()
}
