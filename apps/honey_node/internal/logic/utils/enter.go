package utils

import (
	"honey_node/internal/service"
)

type (
	sExample struct{}
)

var (
	logger = service.Logs().Cat("utils")
)

func init() {
	service.RegisterExample(Example())
	logger.Info("Init Examples success")
}

var insExample = sExample{}

func Example() *sExample {
	return &insExample
}

func (s *sExample) Example() {

}
