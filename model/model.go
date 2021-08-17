package model

type Result struct {
	Total   uint64
	Success uint64
	Failure uint64
	Message []string
}
