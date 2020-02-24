package worker

type Config struct {
    Broker Broker
}

type Broker struct {
    Host     string
    Username string
    Password string
}
