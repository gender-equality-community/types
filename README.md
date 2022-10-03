[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=gender-equality-community_types&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=gender-equality-community_types)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=gender-equality-community_types&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=gender-equality-community_types)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=gender-equality-community_types&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=gender-equality-community_types)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=gender-equality-community_types&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=gender-equality-community_types)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=gender-equality-community_types&metric=sqale_index)](https://sonarcloud.io/summary/new_code?id=gender-equality-community_types)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=gender-equality-community_types&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=gender-equality-community_types)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=gender-equality-community_types&metric=bugs)](https://sonarcloud.io/summary/new_code?id=gender-equality-community_types)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fgender-equality-community%2Ftypes.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fgender-equality-community%2Ftypes?ref=badge_shield)
[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/github.com/gender-equality-community/types)
[![Go Report Card](https://goreportcard.com/badge/github.com/gender-equality-community/types)](https://goreportcard.com/report/github.com/gender-equality-community/types)
---
# types

types holds common data structures used across GEC components.

## Types

### type [Message](https://github.com/gender-equality-community/types/blob/main/message.go#L36)

`type Message struct { ... }`

Message is, simply, the message to be passed between recipients

#### func (Message) [GetTimestamp](https://github.com/gender-equality-community/types/blob/main/message.go#L68)

`func (m Message) GetTimestamp() time.Time`

#### func (Message) [Map](https://github.com/gender-equality-community/types/blob/main/message.go#L62)

`func (m Message) Map() (o map[string]any)`

Map returns a map from the message to be placed on a redis XSTREAM (etc.)

### type [Source](https://github.com/gender-equality-community/types/blob/main/message.go#L12)

`type Source uint8`

Source signifies the source of a message; whether it's come
in from whatsapp, slack, some kind of auto-responder, or just
completely unknown

#### Constants

```golang
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
```

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
