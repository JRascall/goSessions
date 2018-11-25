package sessions

type Session struct {
	ssid      string
	expiry    string
	ip        string
	navigator string
}

func createSession(ssid string) ISession {
	return &Session{
		ssid: ssid,
	}
}

func (s *Session) SSID() string {
	return s.ssid
}

func (s *Session) Expiry() string {
	return s.expiry
}

func (s *Session) IP() string {
	return s.ip
}

func (s *Session) Navigator() string {
	return s.navigator
}

func (s *Session) Add(key string, data interface{}) {

}

func (s *Session) Delete(key string) {

}

func (s *Session) Get(key string) interface{} {
	return nil
}
