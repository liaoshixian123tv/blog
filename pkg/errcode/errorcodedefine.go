package errorcode

var (
	Success                   = NewError(0, "Sucess!")
	ServerError               = NewError(10000000, "Server error")
	InvaildParams             = NewError(10000001, "Invaild parameters")
	NotFound                  = NewError(10000002, "Information can not found")
	UnauthorizedAuthNotExist  = NewError(10000003, "Failed to authorized, can not found match Appkey and AppSecret")
	UnauthorizedToken         = NewError(10000004, "Failed to authorized, token error")
	UnauthorizedTokenTimeout  = NewError(10000005, "Failed to authorized, token timeout")
	UnauthorizedTokenGenerate = NewError(10000006, "Failed to authorized, token generated failed")
	TooManyRequest            = NewError(10000007, "Too many requests")
)
