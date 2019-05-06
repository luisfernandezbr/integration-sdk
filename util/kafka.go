package util

import (
	"context"
	"encoding/binary"
	"sync/atomic"

	dkafka "github.com/dangkaka/go-kafka-avro"
	"github.com/linkedin/goavro"
	kafka "github.com/segmentio/kafka-go"
)

// KafkaProducer is a wrapper for handling kafka
type KafkaProducer struct {
	w     *kafka.Writer
	sc    *dkafka.CachedSchemaRegistryClient
	topic string
	count int64
}

// Close the producer
func (p *KafkaProducer) Close() error {
	return p.w.Close()
}

// Count will return the number of messages that the producer has sent
func (p *KafkaProducer) Count() int64 {
	return atomic.LoadInt64(&p.count)
}

// Send kafka message
func (p *KafkaProducer) Send(ctx context.Context, schema string, key []byte, value []byte) error {
	avroCodec, err := goavro.NewCodec(schema)
	schemaID, err := p.sc.CreateSubject(p.topic+"-value", avroCodec)
	if err != nil {
		return err
	}
	binarySchemaID := make([]byte, 4)
	binary.BigEndian.PutUint32(binarySchemaID, uint32(schemaID))

	native, _, err := avroCodec.NativeFromTextual(value)
	if err != nil {
		return err
	}

	// Convert native Go form to binary Avro data
	binaryValue, err := avroCodec.BinaryFromNative(nil, native)
	if err != nil {
		return err
	}

	var binaryMsg []byte
	// first byte is magic byte, always 0 for now
	binaryMsg = append(binaryMsg, byte(0))
	//4-byte schema ID as returned by the Schema Registry
	binaryMsg = append(binaryMsg, binarySchemaID...)
	//avro serialized data in Avro’s binary encoding
	binaryMsg = append(binaryMsg, binaryValue...)

	msg := kafka.Message{
		Key:   key,
		Value: binaryMsg,
	}

	if err := p.w.WriteMessages(ctx, msg); err != nil {
		return err
	}

	atomic.AddInt64(&p.count, 1)

	return nil
}

// NewKafkaProducer creates a new kakfa producer for a given topic
func NewKafkaProducer(brokerList []string, schemaRegistryServers []string, topic string) *KafkaProducer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:   brokerList,
		Topic:     topic,
		Balancer:  &kafka.Hash{},
		Async:     false,
		BatchSize: 1,
	})
	client := dkafka.NewCachedSchemaRegistryClient(schemaRegistryServers)
	return &KafkaProducer{writer, client, topic, 0}
}
