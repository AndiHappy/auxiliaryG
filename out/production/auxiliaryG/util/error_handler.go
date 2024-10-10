package util

import (
	"errors"
	"log"
	"net"
	"syscall"
)

func CheckErr(err error) {
	if err == nil {
		return
	}
	log.Fatal(err)
	var netError net.Error
	var netOpError net.OpError
	var syscallErr syscall.Errno
	switch {
	case errors.As(err, &netError):
		{
			if netError.Timeout() {
				println("netError: Timeout")
				return
			}
		}
	case errors.As(err, netOpError):
		if netOpError.Op == "dial" {
			println("Unknown host")
		} else if netOpError.Op == "read" {
			println("Connection refused")
		}
	case errors.As(err, syscallErr):
		if errors.Is(syscallErr, syscall.ECONNREFUSED) {
			println("Connection refused")
		}
	default:
		println(err)
	}

}
