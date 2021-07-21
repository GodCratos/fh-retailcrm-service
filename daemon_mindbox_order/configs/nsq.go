package configs

import "os"

const (
	NsqHost              = "NSQ_HOST"
	NsqPort              = "NSQ_PORT"
	NsqTopicMindboxOrder = "NSQ_TOPIC_MINDBOX_ORDER"
	NsqChannel           = "NSQ_CHANNEL"
)

func NsqGetUrl() string {
	return os.Getenv(NsqHost) + ":" + os.Getenv(NsqPort)
}

func NsqGetTopicMindboxOrder() string {
	return os.Getenv(NsqTopicMindboxOrder)
}
func NsqGetChannel() string {
	return os.Getenv(NsqChannel)
}
