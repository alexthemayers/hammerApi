package api

type ValidationFixtureEndpointCustomerGetParameters struct {
	Customer string `url:"-"`
	Id       string `url:"id,omitempty"`
}

type ValidationFixtureEndpointCustomerGetRequestBody struct{}

type ValidationFixtureEndpointCustomerGetResponse interface {
	GetStatus() int
}

type ValidationFixtureEndpointCustomerGet200Response ValidationFixtureEndpointCustomerGet200ResponseModel

func (r ValidationFixtureEndpointCustomerGet200Response) GetStatus() int {
	return 200
}
