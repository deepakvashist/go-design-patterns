package clients

// GcpBuilder is the GCP client builder.
type GcpBuilder struct {
	client Client
}

func (g *GcpBuilder) setVersion() {
	g.client.version = "v1"
}

func (g *GcpBuilder) setBaseURL() {
	g.client.baseURL = "https://samplegcp.com"
}

func (g *GcpBuilder) setMaxRetries() {
	g.client.maxRetries = 3
}

func (g *GcpBuilder) setAllowedStatusCodes() {
	g.client.allowedStatusCodes = []int{200, 201, 203}
}

// GetClient returns the clients instance.
func (g *GcpBuilder) GetClient() Client {
	return g.client
}
