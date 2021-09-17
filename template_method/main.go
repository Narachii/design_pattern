package main

import "fmt"

func main() {
	smsOTP := &sms{}
	o := otp{smsOTP}
	o.genAndSendOTP(4)
	fmt.Println("")

	emailOTP := &email{}
	o = otp{emailOTP}
	o.genAndSendOTP(4)
}
