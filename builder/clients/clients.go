package clients

// Client is the clients base struct containing the cleint's data.
type Client struct {
	version            string
	baseURL            string
	maxRetries         int
	allowedStatusCodes []int
}

// GetVersion returns the client's version.
func (c *Client) GetVersion() string {
	return c.version
}

// GetBaseURL returns the client's base URL.
func (c *Client) GetBaseURL() string {
	return c.baseURL
}

// GetMaxRetries returns the client's max retries.
func (c *Client) GetMaxRetries() int {
	return c.maxRetries
}

// GetAllowedStatusCodes returns the client's allowed status codes.
func (c *Client) GetAllowedStatusCodes() []int {
	return c.allowedStatusCodes
}

// ClientBuilder is the client builder to be implemented.
type ClientBuilder interface {
	setVersion()
	setBaseURL()
	setMaxRetries()
	setAllowedStatusCodes()
	GetClient() Client
}

// Director is the struct that contains the core methods to build the client.
type Director struct {
	builder ClientBuilder
}

// SetBuilder set the builder into the director.
func (d *Director) SetBuilder(builder ClientBuilder) {
	d.builder = builder
}

// BuildClient builds and return the client.
func (d *Director) BuildClient() Client {
	d.builder.setVersion()
	d.builder.setBaseURL()
	d.builder.setMaxRetries()
	d.builder.setAllowedStatusCodes()

	return d.builder.GetClient()
}
