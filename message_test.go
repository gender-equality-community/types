package types

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSource_MarshalBinary(t *testing.T) {
	var empty Source

	for _, test := range []struct {
		name   string
		s      Source
		expect []byte
	}{
		{"Empty Source marshals to 0x0", empty, []byte("0")},
		{"Autoresponse marshals to 0x2", SourceAutoresponse, []byte("2")},
	} {
		t.Run(test.name, func(t *testing.T) {
			received, _ := test.s.MarshalBinary()
			if !cmp.Equal(test.expect, received) {
				t.Error(cmp.Diff(test.expect, received))
			}
		})
	}
}

func TestSource_UnmarshalBinary(t *testing.T) {
	for _, test := range []struct {
		name   string
		s      []byte
		expect Source
	}{
		{"Empty Source unmarshals from 0x0", []byte{}, SourceUnknown},
		{"Garbage/ broken data unmarshals to Unknown", []byte("hello, world!"), SourceUnknown},
		{"Autoresponse unmarshals from 0x2", []byte("2"), SourceAutoresponse},
	} {
		t.Run(test.name, func(t *testing.T) {
			var received Source

			_ = received.UnmarshalBinary(test.s)
			if !cmp.Equal(test.expect, received) {
				t.Error(cmp.Diff(test.expect, received))
			}
		})
	}
}

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
		{"Bad types", map[string]interface{}{"source": true, "ts": "invalid"}, Message{}, true},
		{"Empty message", map[string]interface{}{}, Message{}, false},
		{"Happy path", map[string]interface{}{"source": 2, "ts": 0, "id": "some-id", "msg": "<3"}, Message{
			Source:    SourceAutoresponse,
			ID:        "some-id",
			Timestamp: 0,
			Message:   "<3",
		}, false},
		{"Happy path, from redis", map[string]interface{}{"source": "2", "ts": "0", "id": "some-id", "msg": "<3"}, Message{
			Source:    SourceAutoresponse,
			ID:        "some-id",
			Timestamp: 0,
			Message:   "<3",
		}, false},
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
		{"Empty message", Message{}, map[string]any{"id": "", "msg": "", "source": Source(0), "ts": int64(0)}},
		{"Happy path", Message{
			Source:    SourceAutoresponse,
			ID:        "some-id",
			Timestamp: 0,
			Message:   "<3",
		}, map[string]any{"source": Source(2), "ts": int64(0), "id": "some-id", "msg": "<3"}},
	} {
		t.Run(test.name, func(t *testing.T) {
			m := test.m.Map()

			if !cmp.Equal(test.expect, m) {
				t.Error(cmp.Diff(test.expect, m))
			}
		})
	}
}

func TestMessage_GetTimestamp(t *testing.T) {
	expect := "2022-10-02 18:31:15 +0000 UTC"
	got := Message{Timestamp: 1664735475}.GetTimestamp().UTC().String()

	if expect != got {
		t.Errorf("expected %q, received %q", expect, got)
	}
}
