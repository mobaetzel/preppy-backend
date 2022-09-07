package models

type ListResponse[T any] struct {
	Total uint64 `json:"total"`
	Items []T    `json:"items"`
}
