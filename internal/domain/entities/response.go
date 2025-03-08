package entities

type Response struct {
	ResponseCode string `json:"responseCode"`
	ResponseDesc string `json:"responseDesc"`
	ResponseData any    `json:"responseData,omitempty"`
}
