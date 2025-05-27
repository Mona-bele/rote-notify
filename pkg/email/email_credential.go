package email

import "errors"

type EmailCredential struct {
	From     string
	Password string
	Port     int
	Host     string
}

type CredentialInterface interface {
	GetPassword() string
	GetPort() int
	GetHost() string
	GetFrom() string
}

func NewEmailCredential(from, password string, port int, host string) (*EmailCredential, error) {
	eCredential := &EmailCredential{
		From:     from,
		Password: password,
		Port:     port,
		Host:     host,
	}
	if err := eCredential.validate(); err != nil {
		return nil, err
	}

	return eCredential, nil
}

func (e EmailCredential) validate() error {
	if e.From == "" {
		return errors.New("from is required")
	}
	if e.Password == "" {
		return errors.New("password is required")
	}

	if e.Port == 0 {
		return errors.New("port is required")
	}

	if e.Host == "" {
		return errors.New("host is required")
	}

	return nil
}

func (e EmailCredential) GetFrom() string {
	return e.From
}

func (e EmailCredential) GetPassword() string {
	return e.Password
}

func (e EmailCredential) GetPort() int {
	return e.Port
}

func (e EmailCredential) GetHost() string {
	return e.Host
}
