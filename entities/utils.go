package entities

import "google.golang.org/protobuf/proto"

func ToWorkerEntity(origin []byte) (*Worker, error) {
	ans := &Worker{}
	err := proto.Unmarshal(origin, ans)
	return ans, err
}
