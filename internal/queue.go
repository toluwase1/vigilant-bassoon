package internal

var queue map[string][]any

const (
	UserQueue        = "queue.user.verify"
	TransactionQueue = "queue.transaction.process"
)

func PushToQueue(key string, data any) {
	queue[key] = append(queue[key], data)
}

func PopFromQueue(key string) any {
	var value = queue[key][len(queue[key])-1]
	queue[key] = queue[key][:len(queue[key])-1]
	return value
}
