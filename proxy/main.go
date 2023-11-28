package main

import "fmt"

type Subject interface {
	Do(action string) string
}

type RealSubject struct {
}

func (r *RealSubject) Do(action string) string {
	return "RealSubject action: " + action
}

type Proxy struct {
	realSubject *RealSubject
}

func (p *Proxy) Do(action string) string {
	if p.realSubject == nil {
		p.realSubject = &RealSubject{}
	}
	fmt.Println("Proxy: some work before invoking RealSubject")
	return p.realSubject.Do(action)
}

func main() {
	proxy := &Proxy{}
	result := proxy.Do("some action")
	fmt.Println(result)
}
