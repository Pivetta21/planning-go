package core

import "time"

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

// Cookie Duration.
//
// Type used to designate cookie duration constants.
type cookieDurationKind = time.Duration

var (
	CookieDurationAuthSession cookieDurationKind = 3600 * time.Second
)
