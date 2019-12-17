// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/pdeguing/empire-and-foundation/ent/migrate"

	"github.com/pdeguing/empire-and-foundation/ent/planet"
	"github.com/pdeguing/empire-and-foundation/ent/session"
	"github.com/pdeguing/empire-and-foundation/ent/timer"
	"github.com/pdeguing/empire-and-foundation/ent/user"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Planet is the client for interacting with the Planet builders.
	Planet *PlanetClient
	// Session is the client for interacting with the Session builders.
	Session *SessionClient
	// Timer is the client for interacting with the Timer builders.
	Timer *TimerClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	c := config{log: log.Println}
	c.options(opts...)
	return &Client{
		config:  c,
		Schema:  migrate.NewSchema(c.driver),
		Planet:  NewPlanetClient(c),
		Session: NewSessionClient(c),
		Timer:   NewTimerClient(c),
		User:    NewUserClient(c),
	}
}

// Open opens a connection to the database specified by the driver name and a
// driver-specific data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil

	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug}
	return &Tx{
		config:  cfg,
		Planet:  NewPlanetClient(cfg),
		Session: NewSessionClient(cfg),
		Timer:   NewTimerClient(cfg),
		User:    NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Planet.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true}
	return &Client{
		config:  cfg,
		Schema:  migrate.NewSchema(cfg.driver),
		Planet:  NewPlanetClient(cfg),
		Session: NewSessionClient(cfg),
		Timer:   NewTimerClient(cfg),
		User:    NewUserClient(cfg),
	}
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// PlanetClient is a client for the Planet schema.
type PlanetClient struct {
	config
}

// NewPlanetClient returns a client for the Planet from the given config.
func NewPlanetClient(c config) *PlanetClient {
	return &PlanetClient{config: c}
}

// Create returns a create builder for Planet.
func (c *PlanetClient) Create() *PlanetCreate {
	return &PlanetCreate{config: c.config}
}

// Update returns an update builder for Planet.
func (c *PlanetClient) Update() *PlanetUpdate {
	return &PlanetUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *PlanetClient) UpdateOne(pl *Planet) *PlanetUpdateOne {
	return c.UpdateOneID(pl.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *PlanetClient) UpdateOneID(id int) *PlanetUpdateOne {
	return &PlanetUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Planet.
func (c *PlanetClient) Delete() *PlanetDelete {
	return &PlanetDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PlanetClient) DeleteOne(pl *Planet) *PlanetDeleteOne {
	return c.DeleteOneID(pl.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PlanetClient) DeleteOneID(id int) *PlanetDeleteOne {
	return &PlanetDeleteOne{c.Delete().Where(planet.ID(id))}
}

// Create returns a query builder for Planet.
func (c *PlanetClient) Query() *PlanetQuery {
	return &PlanetQuery{config: c.config}
}

// Get returns a Planet entity by its id.
func (c *PlanetClient) Get(ctx context.Context, id int) (*Planet, error) {
	return c.Query().Where(planet.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PlanetClient) GetX(ctx context.Context, id int) *Planet {
	pl, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return pl
}

// QueryOwner queries the owner edge of a Planet.
func (c *PlanetClient) QueryOwner(pl *Planet) *UserQuery {
	query := &UserQuery{config: c.config}
	id := pl.ID
	step := sql.NewStep(
		sql.From(planet.Table, planet.FieldID, id),
		sql.To(user.Table, user.FieldID),
		sql.Edge(sql.M2O, true, planet.OwnerTable, planet.OwnerColumn),
	)
	query.sql = sql.Neighbors(pl.driver.Dialect(), step)

	return query
}

// QueryTimers queries the timers edge of a Planet.
func (c *PlanetClient) QueryTimers(pl *Planet) *TimerQuery {
	query := &TimerQuery{config: c.config}
	id := pl.ID
	step := sql.NewStep(
		sql.From(planet.Table, planet.FieldID, id),
		sql.To(timer.Table, timer.FieldID),
		sql.Edge(sql.O2M, false, planet.TimersTable, planet.TimersColumn),
	)
	query.sql = sql.Neighbors(pl.driver.Dialect(), step)

	return query
}

// SessionClient is a client for the Session schema.
type SessionClient struct {
	config
}

// NewSessionClient returns a client for the Session from the given config.
func NewSessionClient(c config) *SessionClient {
	return &SessionClient{config: c}
}

// Create returns a create builder for Session.
func (c *SessionClient) Create() *SessionCreate {
	return &SessionCreate{config: c.config}
}

// Update returns an update builder for Session.
func (c *SessionClient) Update() *SessionUpdate {
	return &SessionUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *SessionClient) UpdateOne(s *Session) *SessionUpdateOne {
	return c.UpdateOneID(s.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *SessionClient) UpdateOneID(id int) *SessionUpdateOne {
	return &SessionUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Session.
func (c *SessionClient) Delete() *SessionDelete {
	return &SessionDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SessionClient) DeleteOne(s *Session) *SessionDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SessionClient) DeleteOneID(id int) *SessionDeleteOne {
	return &SessionDeleteOne{c.Delete().Where(session.ID(id))}
}

// Create returns a query builder for Session.
func (c *SessionClient) Query() *SessionQuery {
	return &SessionQuery{config: c.config}
}

// Get returns a Session entity by its id.
func (c *SessionClient) Get(ctx context.Context, id int) (*Session, error) {
	return c.Query().Where(session.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SessionClient) GetX(ctx context.Context, id int) *Session {
	s, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return s
}

// TimerClient is a client for the Timer schema.
type TimerClient struct {
	config
}

// NewTimerClient returns a client for the Timer from the given config.
func NewTimerClient(c config) *TimerClient {
	return &TimerClient{config: c}
}

// Create returns a create builder for Timer.
func (c *TimerClient) Create() *TimerCreate {
	return &TimerCreate{config: c.config}
}

// Update returns an update builder for Timer.
func (c *TimerClient) Update() *TimerUpdate {
	return &TimerUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *TimerClient) UpdateOne(t *Timer) *TimerUpdateOne {
	return c.UpdateOneID(t.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *TimerClient) UpdateOneID(id int) *TimerUpdateOne {
	return &TimerUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Timer.
func (c *TimerClient) Delete() *TimerDelete {
	return &TimerDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TimerClient) DeleteOne(t *Timer) *TimerDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TimerClient) DeleteOneID(id int) *TimerDeleteOne {
	return &TimerDeleteOne{c.Delete().Where(timer.ID(id))}
}

// Create returns a query builder for Timer.
func (c *TimerClient) Query() *TimerQuery {
	return &TimerQuery{config: c.config}
}

// Get returns a Timer entity by its id.
func (c *TimerClient) Get(ctx context.Context, id int) (*Timer, error) {
	return c.Query().Where(timer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TimerClient) GetX(ctx context.Context, id int) *Timer {
	t, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return t
}

// QueryPlanet queries the planet edge of a Timer.
func (c *TimerClient) QueryPlanet(t *Timer) *PlanetQuery {
	query := &PlanetQuery{config: c.config}
	id := t.ID
	step := sql.NewStep(
		sql.From(timer.Table, timer.FieldID, id),
		sql.To(planet.Table, planet.FieldID),
		sql.Edge(sql.M2O, true, timer.PlanetTable, timer.PlanetColumn),
	)
	query.sql = sql.Neighbors(t.driver.Dialect(), step)

	return query
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	return &UserCreate{config: c.config}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	return &UserUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	return c.UpdateOneID(u.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	return &UserUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	return &UserDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	return &UserDeleteOne{c.Delete().Where(user.ID(id))}
}

// Create returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	u, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return u
}

// QueryPlanets queries the planets edge of a User.
func (c *UserClient) QueryPlanets(u *User) *PlanetQuery {
	query := &PlanetQuery{config: c.config}
	id := u.ID
	step := sql.NewStep(
		sql.From(user.Table, user.FieldID, id),
		sql.To(planet.Table, planet.FieldID),
		sql.Edge(sql.O2M, false, user.PlanetsTable, user.PlanetsColumn),
	)
	query.sql = sql.Neighbors(u.driver.Dialect(), step)

	return query
}
