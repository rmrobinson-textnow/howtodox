package api

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// SampleType is a sample type we expect the API to return
type SampleType struct {
	FieldA string `json:"field_a"`
	FieldB int `json:"field_b"`
}

func TestSuccess(t *testing.T) {
	resp := []byte(`{"error_code":"", "result":{"field_a":"sample string","field_b":42}}`)

	respData := new(SampleType)
	respType := &TNAPIResponse{
		Result: respData,
	}

	err := json.Unmarshal(resp, respType)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "sample string", respData.FieldA)
	assert.Equal(t, 42, respData.FieldB)
}

func TestCaptcha(t *testing.T) {
	resp := []byte(`{"error_code":"CAPTCHA_REQUIRED", "result":{"captcha_link":"http://some_captcha.textnow.com","captcha_token":"asdfasdfasdf"}}`)

	respData := new(SampleType)
	respType := &TNAPIResponse{
		Result: respData,
	}

	err := json.Unmarshal(resp, respType)
	if err != nil {
		t.Fatal(err)
	}

	captcha, captchaOk := respType.Result.(*CaptchaDetails)
	assert.Equal(t, "CAPTCHA_REQUIRED", respType.ErrorCode)
	assert.True(t, captchaOk)
	assert.Equal(t, "http://some_captcha.textnow.com", captcha.Link)
	assert.Equal(t, "asdfasdfasdf", captcha.Token)
}

func TestRandomError(t *testing.T) {
	resp := []byte(`{"error_code":"SOME_OTHER_ERROR", "result":"life's a party"}`)

	respData := new(SampleType)
	respType := &TNAPIResponse{
		Result: respData,
	}

	err := json.Unmarshal(resp, respType)
	if err != nil {
		t.Fatal(err)
	}

	_, captchaOk := respType.Result.(*CaptchaDetails)
	errStr, errStrOk := respType.Result.(*string)
	assert.Equal(t, "SOME_OTHER_ERROR", respType.ErrorCode)
	assert.False(t, captchaOk)
	assert.True(t, errStrOk)
	assert.Equal(t, "life's a party", *errStr)
}
