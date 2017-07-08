package spec

type ProxySpec struct {
	HandlerSpec HandlerSpec `json:"handler_spec"`
	ProcessSpec ProcessSpec `json:"process_spec"`
}

type HandlerSpec struct {
	WorkerCount  int         `json:"worker_count"`
	Port         string      `json:"port"`
	QueueConfigs []QueueSpec `json:"queue_spec"`
}

type QueueSpec struct {
	Tag      string `json:"tag"`
	Priority int    `json:"priority"`
}

type ProcessSpec struct {
	Timeout int `json:"timeout"`
}
