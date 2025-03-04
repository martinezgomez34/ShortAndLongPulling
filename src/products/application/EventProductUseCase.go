package application

import (
	"sync"
)

type ProductEvent struct {
	mu        sync.Mutex
	listeners []chan struct{}
}

var (
	productEvent       = &ProductEvent{}
	productActionEvent = &ProductEvent{}
)

func (e *ProductEvent) Notify() {
	e.mu.Lock()
	defer e.mu.Unlock()
	for _, ch := range e.listeners {
		ch <- struct{}{}
	}
	e.listeners = nil
}

func (e *ProductEvent) Wait() <-chan struct{} {
	e.mu.Lock()
	defer e.mu.Unlock()
	ch := make(chan struct{}, 1)
	e.listeners = append(e.listeners, ch)
	return ch
}

func NotifyProductCreated() {
	productEvent.Notify()
}

func WaitForProductCreated() <-chan struct{} {
	return productEvent.Wait()
}

func NotifyProductAction() {
	productActionEvent.Notify()
}

func WaitForProductAction() <-chan struct{} {
	return productActionEvent.Wait()
}