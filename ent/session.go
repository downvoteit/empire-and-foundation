// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
)

// Session is the model entity for the Session schema.
type Session struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Token holds the value of the "token" field.
	Token string `json:"token,omitempty"`
	// Data holds the value of the "data" field.
	Data []byte `json:"data,omitempty"`
	// Expiry holds the value of the "expiry" field.
	Expiry time.Time `json:"expiry,omitempty"`
}

// FromRows scans the sql response data into Session.
func (s *Session) FromRows(rows *sql.Rows) error {
	var scans struct {
		ID     int
		Token  sql.NullString
		Data   []byte
		Expiry sql.NullTime
	}
	// the order here should be the same as in the `session.Columns`.
	if err := rows.Scan(
		&scans.ID,
		&scans.Token,
		&scans.Data,
		&scans.Expiry,
	); err != nil {
		return err
	}
	s.ID = scans.ID
	s.Token = scans.Token.String
	s.Data = scans.Data
	s.Expiry = scans.Expiry.Time
	return nil
}

// Update returns a builder for updating this Session.
// Note that, you need to call Session.Unwrap() before calling this method, if this Session
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Session) Update() *SessionUpdateOne {
	return (&SessionClient{s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Session) Unwrap() *Session {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Session is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Session) String() string {
	var builder strings.Builder
	builder.WriteString("Session(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", token=")
	builder.WriteString(s.Token)
	builder.WriteString(", data=")
	builder.WriteString(fmt.Sprintf("%v", s.Data))
	builder.WriteString(", expiry=")
	builder.WriteString(s.Expiry.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Sessions is a parsable slice of Session.
type Sessions []*Session

// FromRows scans the sql response data into Sessions.
func (s *Sessions) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		scans := &Session{}
		if err := scans.FromRows(rows); err != nil {
			return err
		}
		*s = append(*s, scans)
	}
	return nil
}

func (s Sessions) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
