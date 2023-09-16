package lib

import "time"

type RawReq struct {
	Id      uint32
	Payload []byte
}

type RawResp struct {
	Payload []byte
	Id      uint32
	elapse  time.Duration
	err     error
}

type CallResult struct {
	Id       uint32
	Request  RawReq
	Response RawResp
	Code     RetCode
	Detail   string
}
