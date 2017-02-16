# Protobuf Stream

[![GoDoc](https://godoc.org/github.com/haochi/protobuf-stream-go?status.svg)](https://godoc.org/github.com/haochi/protobuf-stream-go)

This is used for reading and writing protobuf streams, specifically for messages that implements the following interface:

```golang
type Message interface {
	proto.Marshaler
	proto.Unmarshaler
	proto.Sizer
}
```

The official code generator [golang/protobuf](https://github.com/golang/protobuf) doesn't support `proto.Sizer`,
but you can check out [gogo/protobuf](https://github.com/gogo/protobuf).

## Write / WriteWithLock
```golang
if err := stream.Write(writer, message); err != nil {
    log.Printf("Can't write: %v", err)
}
```

## Read

```golang
for err != io.EOF {
    var message Message
    if err = stream.Read(reader, &message); err == nil {
        log.Println(message)
    } else if err != nil && err != io.EOF {
        log.Println(err)
    }
}
```
