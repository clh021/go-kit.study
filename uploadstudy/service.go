package main

type uploadService struct {
}

func (u *uploadService) list() []string {
	var lists []string
	lists = append(lists, "1.txt", "2.pdf", "3.jpg")
	return lists
}
