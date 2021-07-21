package configs

import "os"

const (
	NsqHost              = "NSQ_HOST"
	NsqPort              = "NSQ_PORT"
	NsqTopicRetailClient = "NSQ_TOPIC_RETAIL_CLIENT"
	NsqTopicRetailOrder  = "NSQ_TOPIC_RETAIL_ORDER"
	NsqChannel           = "NSQ_CHANNEL"
)

func NsqGetUrl() string {
	return os.Getenv(NsqHost) + ":" + os.Getenv(NsqPort)
}

func NsqGetTopicRetailClient() string {
	return os.Getenv(NsqTopicRetailClient)
}

func NsqGetTopicRetailOrder() string {
	return os.Getenv(NsqTopicRetailOrder)
}

func NsqGetChannel() string {
	return os.Getenv(NsqChannel)
}
