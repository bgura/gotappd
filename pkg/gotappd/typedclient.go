package gotappd

// Gotappd client wrapper with typed parsing of ressponses
type GotappdTyped struct {
	client  *Gotappd
}

// Create an instance of the Gotappd client
func newTypedClient(client *Gotappd) *GotappdTyped {
	api := GotappdTyped{client}
	
	return &api
}

// Send a 'ping' to the server and test its connection.
func (api GotappdTyped) Ping() (PingResult, error) {
	return api.client.Ping()
}