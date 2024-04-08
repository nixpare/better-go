package main

type StringWrapper struct {
	s string
}

func NewStringWrapper(s string) StringWrapper {
	return StringWrapper{ s: s }
}

func (sw StringWrapper) Unwrap() string {
	return sw.s
}

func (sw StringWrapper) Self() *StringWrapper {
	return &sw
}
