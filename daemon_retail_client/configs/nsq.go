package configs

import "os"

const (
	NsqHost               = "NSQ_HOST"
	NsqPort               = "NSQ_PORT"
	NsqTopicRetailClient  = "NSQ_TOPIC_RETAIL_CLIENT"
	NsqTopicMindboxClient = "NSQ_TOPIC_MINDBOX_CLIENT"
	NsqChannel            = "NSQ_CHANNEL"
)

func NsqGetUrl() string {
	return os.Getenv(NsqHost) + ":" + os.Getenv(NsqPort)
}

func NsqGetTopicRetailClient() string {
	return os.Getenv(NsqTopicRetailClient)
}

func NsqGetTopicMindboxClient() string {
	return os.Getenv(NsqTopicMindboxClient)
}
func NsqGetChannel() string {
	return os.Getenv(NsqChannel)
}
