package http

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"

	"avenger/structs"
	"avenger/tools/coding"
)

type client struct {
	config *structs.RequestConfig
}

func New() *client {
	return &client{
		config: structs.NewRequestConfig(),
	}
}

func (c *client) AddHeader(key, value string) *client {
	c.config.Header[key] = value
	return c
}

func (c *client) AddHeaders(headers structs.SSM) *client {
	c.config.Header = headers
	return c
}

func (c *client) AddQuery(key, value string) *client {
	c.config.Query[key] = value
	return c
}

func (c *client) AddQuerys(pramas structs.SSM) *client {
	c.config.Query = pramas
	return c
}

func (c *client) SetCookie(key, value, path, domain string, age int, httpOnly bool) *client {
	c.config.SetCookie(key, value, path, domain, age, httpOnly)
	return c
}

func (c *client) SetProxy(proxy string) *client {
	c.config.Proxy = proxy
	return c
}

func (c *client) SetTimeout(timeout time.Duration) *client {
	c.config.TimeOut = timeout
	return c
}

func (c *client) build() *resty.Request {
	_client := resty.New().SetTimeout(c.config.TimeOut)
	//代理设置
	if len(c.config.Proxy) > 0 {
		_client.SetProxy(c.config.Proxy)
	}
	request := _client.R()
	//Header设置
	if len(c.config.Header) > 0 {
		request.SetHeaders(c.config.Header)
	}
	//Url查询参数
	if len(c.config.Query) > 0 {
		request.SetQueryParams(c.config.Query)
	}
	//Cookie
	if len(c.config.Cookies) > 0 {
		request.SetCookies(c.config.Cookies)
	}
	return request
}

func (c *client) Get(url string) (string, error) {
	resp, err := c.build().Get(url)
	if err != nil {
		return "", err
	} else if resp.StatusCode() != 200 {
		return "", fmt.Errorf("http failed status:%d", resp.StatusCode())
	}
	return resp.String(), nil
}

func (c *client) GetFile(url, path string) error {
	resp, err := c.build().SetOutput(path).Get(url)
	if err != nil {
		return err
	} else if resp.StatusCode() != 200 {
		return fmt.Errorf("http failed status:%d", resp.StatusCode())
	}
	return nil
}

func (c *client) PostRaw(url, body string) (string, error) {
	resp, err := c.build().SetBody(body).Post(url)
	if err != nil {
		return "", err
	} else if resp.StatusCode() != 200 {
		return "", fmt.Errorf("http failed status:%d", resp.StatusCode())
	}
	return resp.String(), nil
}

func (c *client) PostJson(url string, data structs.M) (string, error) {
	body := coding.JSONEncode(data)
	resp, err := c.AddHeader("Content-Type", "application/json").build().SetBody(body).Post(url)
	if err != nil {
		return "", err
	} else if resp.StatusCode() != 200 {
		return "", fmt.Errorf("http failed status:%d", resp.StatusCode())
	}
	return resp.String(), nil
}

func (c *client) PostForm(url string, data structs.SSM) (string, error) {
	resp, err := c.build().SetFormData(data).Post(url)
	if err != nil {
		return "", err
	} else if resp.StatusCode() != 200 {
		return "", fmt.Errorf("http failed status:%d", resp.StatusCode())
	}
	return resp.String(), nil
}
