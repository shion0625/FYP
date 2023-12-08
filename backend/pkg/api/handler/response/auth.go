package response

type TokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	UserID       string `json:"userId"`
}

type OTPResponse struct {
	OtpID string `json:"otpId"`
}
