package rule

import (
	"fmt"
	"regexp"
)

// required: v => !!v || "必要字段",
// atleast8characters: v => (!!v && v.length >= 8) || "至少8个字符",
// validemail: v => (!!v && /.+@.+\..+/.test(v)) || "请输入一个有效的 email",
// validphone: v =>
// 	(!!v && !!/^1[345789][0-9]{9}$/.test(v)) || "请输入正确的电话号码哦",
// validcode: v => (!!v && /^[0-9]{6}$/.test(v)) || "请输入正确的验证码",

var EmailRegex *regexp.Regexp
var PhoneRegex *regexp.Regexp
var CodeRegex *regexp.Regexp

func init() {
	PhoneRegex = regexp.MustCompile(`^1[345789][0-9]{9}$`)
	EmailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	CodeRegex = regexp.MustCompile(`^[0-9]{6}$`)
}

type Rule func(string) error

func Check(v string, rules []Rule) error {
	for _, r := range rules {
		if err := r(v); err != nil {
			return err
		}
	}

	return nil
}

func Required(v string) error {
	if len(v) == 0 {
		return fmt.Errorf("必要字段")
	}

	return nil
}

func AtLeast8Characters(v string) error {
	if len(v) < 8 {
		return fmt.Errorf("至少8个字符")
	}

	return nil
}

func AtMost64Characters(v string) error {
	if len(v) >= 64 {
		return fmt.Errorf("至多64个字符")
	}

	return nil
}

func ValidEmail(v string) error {
	if !EmailRegex.MatchString(v) {
		return fmt.Errorf("无效的邮箱")
	}

	return nil
}

func ValidPhone(v string) error {
	if !PhoneRegex.MatchString(v) {
		return fmt.Errorf("无效的电话号码")
	}

	return nil
}

func ValidCode(v string) error {
	if !CodeRegex.MatchString(v) {
		return fmt.Errorf("无效的验证码")
	}

	return nil
}
