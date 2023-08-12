package core

import "time"

// Context Key.
//
// Type used to designate context keys constants.
type contextKeyKind string

const (
	ContextKeyLoggedUser contextKeyKind = "LOGGED_USER"
)

// Cookie Name.
//
// Type used to designate cookie names constants.
type cookieNameKind string

const (
	CookieNameAuthSession cookieNameKind = "SESSION_OT"
)

func (c cookieNameKind) String() string {
	return string(c)
}

// Time.
//
// Type used to designate duration time constants. SI unit: s.
type timeKind int

const (
	TimeExpirationAuthSession timeKind = 3600
	TimeExpirationAuthCookie  timeKind = 3600 * 24 * 90
)

func (c timeKind) Duration() time.Duration {
	return time.Duration(c) * time.Second
}

func (c timeKind) MaxAge() int {
	return int(c.Duration().Seconds())
}

func (c timeKind) Future() time.Time {
	return time.Now().UTC().Add(c.Duration())
}

func (c timeKind) Past() time.Time {
	return time.Now().UTC().Add(-c.Duration())
}
