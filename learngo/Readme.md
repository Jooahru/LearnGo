Run 오류 발생시
1. 오류 
main.go:6:2: no required module provides package github.com/will/learngo/banking: go.mod file not found in current directory or any parent directory; see 'go help modules'

1. 해결
go env -w GO111MODULE=auto