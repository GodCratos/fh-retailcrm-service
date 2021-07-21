package configs

import "os"

const (
	NsqHost              = "NSQ_HOST"
	NsqPort              = "NSQ_PORT"
	NsqTopicRetailOrder  = "NSQ_TOPIC_RETAIL_ORDER"
	NsqTopicMindboxOrder = "NSQ_TOPIC_MINDBOX_ORDER"
	NsqChannel           = "NSQ_CHANNEL"
)

func NsqGetUrl() string {
	return os.Getenv(NsqHost) + ":" + os.Getenv(NsqPort)
}

func NsqGetTopicRetailOrder() string {
	return os.Getenv(NsqTopicRetailOrder)
}

func NsqGetTopicMindboxOrder() string {
	return os.Getenv(NsqTopicMindboxOrder)
}
func NsqGetChannel() string {
	return os.Getenv(NsqChannel)
}
