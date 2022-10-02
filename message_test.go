package types

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestNewMessage(t *testing.T) {
	m := NewMessage(SourceAutoresponse, "some-id", "Hello, world!")

	for _, test := range []struct {
		name   string
		expect any
		got    any
	}{
		{"Source", SourceAutoresponse, m.Source},
		{"ID", "some-id", m.ID},
		{"Message", "Hello, world!", m.Message},
	} {
		t.Run(test.name, func(t *testing.T) {
			if test.expect != test.got {
				t.Errorf("expected %#v, received %#v", test.expect, test.got)
			}
		})

	}
}

func TestParseMessage(t *testing.T) {
	for _, test := range []struct {
		name      string
		m         map[string]interface{}
		expect    Message
		expectErr bool
	}{
		{"Bad types", map[string]interface{}{"source": 1, "ts": "invalid"}, Message{Source: SourceWhatsapp}, true},
		{"Empty message", map[string]interface{}{}, Message{}, false},
		{"Happy path", map[string]interface{}{"source": 2, "ts": 0, "id": "some-id", "msg": "<3"}, Message{
			Source:    SourceAutoresponse,
			ID:        "some-id",
			Timestamp: time.Time{},
			Message:   "<3",
		}, true},
	} {
		t.Run(test.name, func(t *testing.T) {
			m, err := ParseMessage(test.m)
			if err == nil && test.expectErr {
				t.Errorf("expected error")
			} else if err != nil && !test.expectErr {
				t.Errorf("unexpected error: %+v", err)
			}

			if !cmp.Equal(test.expect, m) {
				t.Error(cmp.Diff(test.expect, m))
			}
		})
	}
}

func TestMessage_Map(t *testing.T) {
	for _, test := range []struct {
		name   string
		m      Message
		expect map[string]any
	}{
		{"Empty message", Message{}, map[string]any{"id": "", "msg": "", "source": Source(0), "ts": map[string]any{}}},
		{"Happy path", Message{
			Source:    SourceAutoresponse,
			ID:        "some-id",
			Timestamp: time.Time{},
			Message:   "<3",
		}, map[string]any{"source": Source(2), "ts": map[string]any{}, "id": "some-id", "msg": "<3"}},
	} {
		t.Run(test.name, func(t *testing.T) {
			m := test.m.Map()

			if !cmp.Equal(test.expect, m) {
				t.Error(cmp.Diff(test.expect, m))
			}
		})
	}
}
