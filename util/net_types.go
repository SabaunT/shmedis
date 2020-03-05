package util

type Request struct {
	Method string
	Arguments Arguments
}

type Arguments struct {
	Key string
	Value interface{}
}
