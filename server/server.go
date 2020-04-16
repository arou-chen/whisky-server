package server

type Server interface {
	Server() error
	Stop() error
}
