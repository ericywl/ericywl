package mocks

import (
	"errors"
	"fmt"
	"testing"
)

type mailSender interface {
	Send(to, subject, body string) error
	SendFrom(from, to, subject, body string) error
}

type mockSender struct {
	SendFunc     func(to, subject, body string) error
	SendFromFunc func(from, to, subject, body string) error
}

func (m mockSender) Send(to, subject, body string) error {
	return m.SendFunc(to, subject, body)
}

func (m mockSender) SendFrom(from, to, subject, body string) error {
	return m.SendFromFunc(from, to, subject, body)
}

func sendWelcomeEmail(sender mailSender, to, subject, body string) {
	// send some email
}

func testWelcomeEmail(t *testing.T) {
	errTest := errors.New("nope")
	var msg string

	sender := mockSender{
		SendFunc: func(to, subject, body string) error {
			msg = fmt.Sprintf("(%s) %s: %s", to, subject, body)
			return nil
		},
		SendFromFunc: func(from, to, subject, body string) error {
			return errTest
		},
	}

	sendWelcomeEmail(sender, "to", "subject", "body")
	if msg != "(to) subject: body" {
		t.Error("sendWelcomeEmail:", msg)
	}
}
