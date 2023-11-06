package input

type MonitorServicePort interface {
	Monitor() error
}
