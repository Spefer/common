package proto

// no element request
type NoREQ struct {
}

// no element response
type NoRES struct {
}

// ping request
type PingRES struct {
	TimeUnix int64
}

// push single user request
type PushSMsgREQ struct {
	ServerId int32
	Msg      *PushMsg
}

// push multi users
type PushMMsgREQ struct {
	ServerId int32
	UserId   []string
	Msg      *PushMsg
}

// broadcast message
type PushBroadcastREQ struct {
	Msg *PushMsg
}

// broadcast topic message for the users who subscribe the topic
type PushBroadcastTopicREQ struct {
	Topic    string
	ServerId []int32
	Msg      *PushMsg
}
