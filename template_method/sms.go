package main

import "fmt"

type sms struct {
	// otp
}

func (s *sms) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Print("SMS: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *sms) saveOTPCache(otp string) {
	fmt.Printf("SMS: saving opt:%s to chache\n", otp)
}

func (s *sms) getMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *sms) sendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}

func (s *sms) publishMetric()  {
	fmt.Printf("SMS: publishing metrics\n")
}
