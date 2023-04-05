package internal

var queue = make(map[string][]any)

const (
	UserQueue        = "queue.user.verify"
	TransactionQueue = "queue.transaction.process"
)

func PushToQueue(key string, data any) {
	queue[key] = append(queue[key], data)
}

func GetAllFromQueue(key string) []any {
	snapshot := queue[key]
	queue[key] = []any{}
	return snapshot
}
