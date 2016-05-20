package captcha

import "testing"

func TestVerifyCaptchaWith(t *testing.T) {
	c := GenerateCaptchaWith("t", 6)
	t.Log(c)
	if !VerifyCaptchaWith("t", c) {
		t.Fatal("invalid captcha")
	}
}

func TestGenerateNumberCaptchaWith(t *testing.T) {
	c := GenerateNumberCaptchaWith("2", 6)
	t.Log(c)
	if !VerifyNumberCaptchaWith("2", c) {
		t.Fatal("invalid captcha")
	}
}
