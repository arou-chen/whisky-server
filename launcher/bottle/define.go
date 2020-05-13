package bottle

type LifeCycle interface {
	Start() error
	Stop() error
}