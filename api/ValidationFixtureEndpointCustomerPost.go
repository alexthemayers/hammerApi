package api

type ValidationFixtureEndpointCustomerPostParameters struct {
	Customer string `url:"-"`
	Id       string `url:"id,omitempty"`
}

type ValidationFixtureEndpointCustomerPostRequestBody ValidationFixtureEndpointCustomerPostRequestBodyModel

type ValidationFixtureEndpointCustomerPostResponse interface {
	GetStatus() int
}

type ValidationFixtureEndpointCustomerPost201Response struct{}

func (r ValidationFixtureEndpointCustomerPost201Response) GetStatus() int {
	return 201
}
