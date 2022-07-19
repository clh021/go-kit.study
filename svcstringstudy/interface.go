package main

type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}
