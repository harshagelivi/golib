package work

type HttpSink struct {
	Url           string
	AuthRefresher AuthRefresher
	ClientPool    ClientPool
}

func (hs *HttpSink) Do(p []interface{}) error {
	return nil
}
