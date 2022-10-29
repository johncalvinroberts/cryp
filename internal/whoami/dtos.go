package whoami

type StartWhoamiChallengeDTO struct {
	Email string `json:"email"`
}

type TryWhoamiChallengeRequestDTO struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

type TryWhoamiChallengeResponseDTO struct {
	jwt string `json:"jwt"`
}
