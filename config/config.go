package config

type Config struct {
	NATS
}

type NATS struct {
	ClusterID string
	ClientID  string
	Subject   string
}
