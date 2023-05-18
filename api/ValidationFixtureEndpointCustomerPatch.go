package api

type ValidationFixtureEndpointCustomerPatchParameters struct {
	Customer string `url:"-"`
	Id       string `url:"id,omitempty"`
}

type ValidationFixtureEndpointCustomerPatchRequestBody struct{}

type ValidationFixtureEndpointCustomerPatchResponse interface {
	GetStatus() int
}

type ValidationFixtureEndpointCustomerPatch204Response struct{}

func (r ValidationFixtureEndpointCustomerPatch204Response) GetStatus() int {
	return 204
}
