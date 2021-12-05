package request

import "net/http/cookiejar"

var jar, _ = cookiejar.New(nil)

func DeleteJar() {
	jar, _ = cookiejar.New(nil)
}
