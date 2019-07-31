package api

import (
	"encoding/json"
)

type internalTNAPIResponse struct {
	ErrorCode string `json:"error_code"`
	Result json.RawMessage `json:"result"`
}

// CaptchaDetails contains relevant data to handle a captcha
type CaptchaDetails struct {
	Token string `json:"captcha_token"`
	Link string `json:"captcha_link"`
}

// TNAPIResponse is the standard envelope returned by the TN API
type TNAPIResponse struct {
	ErrorCode string `json:"error_code"`
	Result interface{} `json:""`
}

// UnmarshalJSON is a custom JSON deserializer for the TN API response
func (r *TNAPIResponse) UnmarshalJSON(data []byte) error {
	iResp := &internalTNAPIResponse{}
	err := json.Unmarshal(data, iResp)
	if err != nil {
		return err
	}

	if len(iResp.ErrorCode) > 0 {
		switch iResp.ErrorCode {
		case "CAPTCHA_REQUIRED":
			r.Result = new(CaptchaDetails)
		// This is where we can add additional custom error type handling
		default:
			r.Result = new(string)
		}
	}
	err = json.Unmarshal(iResp.Result, r.Result)
	if err != nil {
		return err
	}
	r.ErrorCode = iResp.ErrorCode
	return nil
}
