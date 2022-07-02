package token

import (
    "time"
    "errors"
    "github.com/google/uuid"
)

var (

    ErrInvalidToken = errors.New("token is invalid")
    ErrExpiredToken = errors.New("token has expired")
)

// Payload contains the payload data of the token
type Payload struct {
    ID uuid.UUID `json:"id"`
    Username string `json:"username"`
    IssuedAt time.Time `json:"issued_at"`
    ExpiredAt time.Time `json:"expired_at"`
}


func NewPayload(username string, duration time.Duration) (*Payload, error) {
    tokenID, err := uuid.NewRandom()

    if err != nil {
        return nil, err
    }

    payload := &Payload{
        ID: tokenID,
        Username: username,
        IssuedAt: time.Now(),
        ExpiredAt: time.Now().Add(duration),
    }

    return payload, nil
}

// Valid checks if the token payload is valid or not
func (payload *Payload) Valid() error {
    if time.Now().After(payload.ExpiredAt) {
        return ErrExpiredToken
    }
    return nil
}
