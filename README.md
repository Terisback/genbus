# GenBus

Generics applied to PubSub.

## Installation

Go >=1.21

```sh
go get -u github.com/Terisback/genbus
```

## Usage

```go
type Event struct {ID string}

bus := genbus.New()

// Subscribe to topic via type (types are used as topic)
sub := genbus.Subscribe[Event](bus)
	
// Send instance of type
go bus.Publish(Event{ID: "ready"})

// Receive typed message from channel 
result := <-res
fmt.Println(result.ID) // `ready`
```

### Use with aliases

```go
type Event struct {ID string}

type AliasedEvent Event

bus := genbus.New()

// Subscribe to topic via aliased type (types are used as topic)
sub := genbus.SubscribeAs[Event, AliasedEvent](bus)
	
// Send instance of type
go bus.Publish(AliasedEvent{ID: "ready"})

// Receive typed message from channel 
result := <-res
fmt.Println(result.ID) // `ready`
```

### Use as untyped `any` eventbus

```go
type Event struct {ID string}

bus := New()

// Subscribe to topic via type instance (types are used as topic)
sub := bus.Subscribe(new(Event))

// Send as any
go bus.Publish(Event{ID: "ready"})

// Receive from channel and typecast
result := <-sub
fmt.Println(result.(Event).ID) // `ready`
```

## Features

- pubsub.Publish is synchronous, it will lock until all subscribers consume message from channel.
- To alter topic but consume same type use [genbus.SubscribeAs](./subscribe.go)