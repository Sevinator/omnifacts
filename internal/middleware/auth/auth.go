package auth

type Auth interface {
	Connect() string
	Disconnect() bool
}
