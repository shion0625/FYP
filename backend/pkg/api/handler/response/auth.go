package response

type TokenResponse struct {
	AccessToken string `json:"accessToken"`
}

type OTPResponse struct {
	OtpID string `json:"otpId"`
}
