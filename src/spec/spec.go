package spec

type HandlerSpec struct {
	WorkerCount  int         `json:"worker_count"`
	Port         int         `json:"port"`
	QueueConfigs []QueueSpec `json:"queue_spec"`
}

type QueueSpec struct {
	Tag      string `json:"tag"`
	Priority int    `json:"priority"`
}
