package consumer

type AckFn func() error
type HandlerFn func(messageID string, values map[string]interface{}, ack AckFn)

type Offset string

const (
	// OffsetOldest from beginning of the time
	OffsetOldest = Offset("0")
	// OffsetNewest from last consumer stop
	OffsetNewest = Offset(">")
)
