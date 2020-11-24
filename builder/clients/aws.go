package clients

// AwsBuilder is the AWS client builder.
type AwsBuilder struct {
	client Client
}

func (a *AwsBuilder) setVersion() {
	a.client.version = "v1"
}

func (a *AwsBuilder) setBaseURL() {
	a.client.baseURL = "https://sampleaws.com"
}

func (a *AwsBuilder) setMaxRetries() {
	a.client.maxRetries = 3
}

func (a *AwsBuilder) setAllowedStatusCodes() {
	a.client.allowedStatusCodes = []int{200, 201, 203}
}

// GetClient returns the clients instance.
func (a *AwsBuilder) GetClient() Client {
	return a.client
}
