package gotappd


func NewClient(clientid string, clientsecret string) *Gotappd {
	client := newClient(clientid, clientsecret)
	return client
}

func NewTypedClient(clientid string, clientsecret string) *GotappdTyped {
	client := newTypedClient(NewClient(clientid, clientsecret))
	return client
}
