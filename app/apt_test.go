package app

import (
	"fmt"
	"strings"
	"testing"
)

func TestLastIndex(t *testing.T) {
	s := "/data/apt-mirror/mirror/mirrors.ustc.edu.cn/ubuntu/pool/main/a/apt/libapt-pkg5.0_1.2.32_amd64.deb"
	fmt.Println(strings.LastIndex(s, "/"))
	fmt.Println(s[:strings.LastIndex(s, "/")])
}
