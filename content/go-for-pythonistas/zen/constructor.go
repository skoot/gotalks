package example

import "net/rpc"

type Database struct {
	client *rpc.Client
}

func NewClient(addr string) (*Database, error) {
	client, err := rpc.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &Database{client: client}, nil
}
