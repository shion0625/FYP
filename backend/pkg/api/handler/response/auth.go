package response

type TokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type OTPResponse struct {
	OtpID string `json:"otpId"`
}
