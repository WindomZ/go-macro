package captcha

import "testing"

func TestVerifyCaptchaWith(t *testing.T) {
	c := GenerateCaptchaWith("t", 6)
	t.Log(c)
	if !VerifyCaptchaWith("t", c) {
		t.Fatal("invalid captcha")
	}
}
