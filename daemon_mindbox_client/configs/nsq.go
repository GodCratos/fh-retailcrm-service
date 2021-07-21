package configs

import "os"

const (
	NsqHost               = "NSQ_HOST"
	NsqPort               = "NSQ_PORT"
	NsqTopicMindboxClient = "NSQ_TOPIC_MINDBOX_CLIENT"
	NsqChannel            = "NSQ_CHANNEL"
)

func NsqGetUrl() string {
	return os.Getenv(NsqHost) + ":" + os.Getenv(NsqPort)
}

func NsqGetTopicMindboxClient() string {
	return os.Getenv(NsqTopicMindboxClient)
}
func NsqGetChannel() string {
	return os.Getenv(NsqChannel)
}
