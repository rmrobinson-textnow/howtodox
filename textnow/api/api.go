package api

import (
	"encoding/json"
)

const (
	captchaRequired = "CAPTCHA_REQUIRED"
)

type internalTNAPIResponse struct {
	ErrorCode string          `json:"error_code"`
	Result    json.RawMessage `json:"result"`
}

// CaptchaDetails contains relevant data to handle a captcha
type CaptchaDetails struct {
	Token string `json:"captcha_token"`
	Link  string `json:"captcha_link"`
}

// TNAPIResponse is the standard envelope returned by the TN API
type TNAPIResponse struct {
	ErrorCode string      `json:"error_code"`
	Result    interface{} `json:""`
}

// UnmarshalJSON is a custom JSON deserializer for the TN API response
func (r *TNAPIResponse) UnmarshalJSON(data []byte) error {
	internalResp := &internalTNAPIResponse{}
	err := json.Unmarshal(data, internalResp)
	if err != nil {
		return err
	}

	if len(internalResp.ErrorCode) > 0 {
		switch internalResp.ErrorCode {
		case captchaRequired:
			r.Result = new(CaptchaDetails)
		// This is where we can add additional custom error type handling
		default:
			r.Result = new(string)
		}
	}
	err = json.Unmarshal(internalResp.Result, r.Result)
	if err != nil {
		return err
	}
	r.ErrorCode = internalResp.ErrorCode
	return nil
}
