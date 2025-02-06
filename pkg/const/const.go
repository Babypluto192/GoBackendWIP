package _const

type Errors string

const (
	Decode_Error Errors = "Decode_Error"
	Server_Error Errors = "Server_Error"
	Bad_Request  Errors = "Bad_Request"
	No_Error     Errors = "No_Error"
)
