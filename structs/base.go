package structs

import "net/http"

import "time"

type M map[string]interface{}

type SSM map[string]string

type RequestConfig struct {
	Proxy   string
	TimeOut time.Duration
	Header  SSM
	Query   SSM
	Form    SSM
	Cookies []*http.Cookie
}

func (r *RequestConfig) SetCookie(key, value, path, domain string, age int, httpOnly bool) {
	cookie := &http.Cookie{
		Name:     key,
		Value:    value,
		Path:     path,
		Domain:   domain,
		MaxAge:   age,
		HttpOnly: httpOnly,
		Secure:   false,
	}
	r.Cookies = append(r.Cookies, cookie)
}

func NewRequestConfig() *RequestConfig {
	return &RequestConfig{
		Header:  SSM{},
		Query:   SSM{},
		Form:    SSM{},
		TimeOut: 15 * time.Second,
		Cookies: []*http.Cookie{},
	}
}
