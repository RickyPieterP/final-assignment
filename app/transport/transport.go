package transport

import "mygram/app/interface/container"

type Tp struct {
	Transport *Transport
}

func SetupTransport(container *container.Container) *Tp {
	transport := NewTransport(container.MygramUsecase, container.Response)
	return &Tp{
		Transport: transport,
	}
}
