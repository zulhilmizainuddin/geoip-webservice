package models

type Retriever interface {
	Query(ipaddress string) interface{}
}