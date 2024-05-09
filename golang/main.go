package main

import (
	"context"
	"fmt"
	"os"
	"path"
	"runtime/pprof"
	"runtime/trace"
	"time"

	"github.com/MohammadAsDev/pub_sub/config"
	"github.com/MohammadAsDev/pub_sub/publisher"
	"github.com/MohammadAsDev/pub_sub/subscriber"
)

const topic = "testing_topic"

func generateSubscribers(nSubsribers int) []subscriber.Subscriber {
	subscribers := []subscriber.Subscriber{}
	subscriber_index := 0
	for subscriber_index < nSubsribers {
		sub, err := subscriber.NewRedisSubscriber(context.Background(), topic)
		if err != nil {
			panic(err)
		}
		subscribers = append(subscribers, sub)
		subscriber_index += 1
	}
	return subscribers
}

func generatePublishers(nPublishers int) []publisher.Publisher {
	publishers := []publisher.Publisher{}
	publisher_index := 0
	for publisher_index < nPublishers {
		pub, err := publisher.NewRedisPublisher(context.Background(), topic)
		if err != nil {
			panic(err)
		}
		publishers = append(publishers, pub)
		publisher_index += 1
	}
	return publishers
}

func publishDumpMessages(publisher publisher.Publisher, nMessages int) {
	// log.Printf("[publisher-%d]: start publishing messages...\n", publisher.Id())

	message_index := 1
	for message_index <= nMessages {
		message := fmt.Sprintf("Message-%d", message_index)
		publisher.Publish(message)
		message_index += 1
	}
}

func waitDumpMessages(subscriber subscriber.Subscriber) {

	// removing printing operation for large number of operations

	message_chan, err_chan := subscriber.Subscribe()
	// log.Printf("[subscriber-%d]: start waiting for messages...\n", subscriber.Id())
	for {
		select {
		case <-message_chan: // case msg := <- message_chan
			// fmt.Printf("[subscriber-%d]: %s\n", subscriber.Id(), string(msg))
		case <-err_chan: // case err := <- err_chan
			// log.Printf("[subscriber-%d]: failed to get messages, error=%v", subscriber.Id(), err)
			return
		}
	}
}

func main() {
	trace_f, _ := os.Create("trace.out")
	trace.Start(trace_f)
	defer trace.Stop()

	cpu_f, _ := os.Create("cpu.out")
	if err := pprof.StartCPUProfile(cpu_f); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	mem_f, _ := os.Create("mem.out")
	defer mem_f.Close()

	global_config, err := config.ReadGlobalConfig(path.Join("..", "config.yaml"))
	if err != nil {
		panic(err)
	}
	n_publishers := global_config.NPubs
	n_subscribers := global_config.NSubs

	publishers := generatePublishers(n_publishers)
	subscribers := generateSubscribers(n_subscribers)

	n_messages := global_config.NMsgs

	for _, pub := range publishers {
		go publishDumpMessages(pub, n_messages)
	}

	for _, sub := range subscribers {
		go waitDumpMessages(sub)
	}

	time.Sleep(time.Duration(25) * time.Second) // random time for testing purposes

	if err := pprof.WriteHeapProfile(mem_f); err != nil {
		panic(err)
	}
}
