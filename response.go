package main

type Response struct {
	Code        int         `json:"code"`
	Data        interface{} `json:"data"`
	Error       error       `json:"error"`
	UserMessage string      `json:"user_message"`
}
