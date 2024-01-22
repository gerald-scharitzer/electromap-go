package electricitymaps

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Session struct {
	AuthToken string // protect this
}

func GetSession() Session {
	return Session{}
}

func GetSessionFromYaml(bytes []byte) (*Session, error) {
	session := GetSession()
	err := yaml.Unmarshal(bytes, &session)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func GetSessionFromFile(file *os.File) (*Session, error) {
	var bytes []byte
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	session, err := GetSessionFromYaml(bytes)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func GetSessionFromStdin() (*Session, error) {
	session, err := GetSessionFromFile(os.Stdin)
	if err != nil {
		return nil, err
	}
	return session, nil
}
