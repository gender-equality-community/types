package types

import (
	"reflect"
	"strconv"
	"time"

	"github.com/mitchellh/mapstructure"
)

// Source signifies the source of a message; whether it's come
// in from whatsapp, slack, some kind of auto-responder, or just
// completely unknown
type Source uint8

const (
	// SourceUnknown is where we simply ust don't know where a message comes from,
	// and is largely only used for zero'd messages, or when errors stop the
	// processing of messages.
	SourceUnknown Source = iota

	// SourceWhatsapp means a message has come from WhatsApp and usually signifies
	// a message from someone seeking advice
	SourceWhatsapp

	// SourceAutoresponse means a message was generated from an application in the
	// processing pipeline somewhere, like the various autoresponses the gec-bot
	// provides
	SourceAutoresponse

	// SourceSlack usually means a message from the GEC _back_ to recipients; though
	// in the future perhaps we'd want to allow slack users to message too...
	// dunno
	SourceSlack
)

// Message is, simply, the message to be passed between recipients
type Message struct {
	Source    Source `mapstructure:"source"`
	ID        string `mapstructure:"id"`
	Timestamp int64  `mapstructure:"ts"`
	Message   string `mapstructure:"msg"`
}

// NewMessage accepts a source, id and string, and returns a new Message
func NewMessage(source Source, id, msg string) Message {
	return Message{
		Source:    source,
		ID:        id,
		Timestamp: time.Now().Unix(),
		Message:   msg,
	}
}

// ParseMessage accepts a map, probably from redis, and turns it into a valid
// Message for processing
func ParseMessage(i map[string]any) (m Message, err error) {
	// If inputs are strings, and we're not expecting strings, such
	// as Source and Timestamp, then cast them properly.
	//
	// This happens when we read from redis, which treats all values
	// as strings
	cast(i, "source")
	cast(i, "ts")

	err = mapstructure.Decode(i, &m)

	return
}

// Map returns a map from the message to be placed on a redis XSTREAM (etc.)
func (m Message) Map() (o map[string]any) {
	_ = mapstructure.Decode(m, &o)

	return
}

func (m Message) GetTimestamp() time.Time {
	return time.Unix(m.Timestamp, 0)
}

func cast(i map[string]any, key string) {
	v, ok := i[key]
	if !ok || reflect.TypeOf(v).String() != "string" {
		return
	}

	i[key], _ = strconv.Atoi(i[key].(string))
}
