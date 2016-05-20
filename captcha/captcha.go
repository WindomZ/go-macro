package captcha

import (
	"github.com/WindomZ/go-random/random"
	"strings"
)

func GenerateCaptcha(n int) string {
	return GenerateCaptchaWith("", n)
}

func GenerateCaptchaWith(prefix string, n int) string {
	if n <= 0 {
		return ""
	} else if len(prefix) >= int(n/2) {
		prefix = prefix[:int(n/2)]
	}
	return prefix + random.RandomCaptcha(n-len(prefix))
}

func VerifyCaptchaWith(prefix string, captcha string) bool {
	return strings.HasPrefix(strings.ToLower(captcha), strings.ToLower(prefix))
}

func GenerateNumberCaptcha(n int) string {
	return GenerateCaptchaWith("", n)
}

func GenerateNumberCaptchaWith(prefix string, n int) string {
	if n <= 0 {
		return ""
	} else if len(prefix) >= int(n/2) {
		prefix = prefix[:int(n/2)]
	}
	return prefix + random.RandomNumCaptcha(n-len(prefix))
}

func VerifyNumberCaptchaWith(prefix string, captcha string) bool {
	return strings.HasPrefix(strings.ToLower(captcha), strings.ToLower(prefix))
}
