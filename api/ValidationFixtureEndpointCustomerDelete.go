package api

type ValidationFixtureEndpointCustomerDeleteParameters struct {
	Customer string `url:"-"`
	Id       string `url:"id,omitempty"`
}

type ValidationFixtureEndpointCustomerDeleteRequestBody struct{}

type ValidationFixtureEndpointCustomerDeleteResponse interface {
	GetStatus() int
}

type ValidationFixtureEndpointCustomerDelete204Response struct{}

func (r ValidationFixtureEndpointCustomerDelete204Response) GetStatus() int {
	return 204
}
