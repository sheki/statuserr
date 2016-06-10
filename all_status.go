package statuserr

var (
	BadRequest = err{StatusCode: 400, Message: "Bad Request"}

	Unauthorized = err{StatusCode: 401, Message: "Unauthorized"}

	PaymentRequired = err{StatusCode: 402, Message: "Payment Required"}

	Forbidden = err{StatusCode: 403, Message: "Forbidden"}

	NotFound = err{StatusCode: 404, Message: "Not Found"}

	MethodNotAllowed = err{StatusCode: 405, Message: "Method Not Allowed"}

	NotAcceptable = err{StatusCode: 406, Message: "Not Acceptable"}

	ProxyAuthenticationRequired = err{StatusCode: 407, Message: "Proxy Authentication Required"}

	RequestTimeout = err{StatusCode: 408, Message: "Request Timeout"}

	Conflict = err{StatusCode: 409, Message: "Conflict"}

	Gone = err{StatusCode: 410, Message: "Gone"}

	LengthRequired = err{StatusCode: 411, Message: "Length Required"}

	PreconditionFailed = err{StatusCode: 412, Message: "Precondition Failed"}

	RequestEntityTooLarge = err{StatusCode: 413, Message: "Request Entity Too Large"}

	UnsupportedMediaType = err{StatusCode: 415, Message: "Unsupported Media Type"}

	RequestedRangeNotSatisfiable = err{StatusCode: 416, Message: "Requested Range Not Satisfiable"}

	ExpectationFailed = err{StatusCode: 417, Message: "Expectation Failed"}

	InternalServerError = err{StatusCode: 500, Message: "Internal Server Error"}

	NotImplemented = err{StatusCode: 501, Message: "Not Implemented"}

	BadGateway = err{StatusCode: 502, Message: "Bad Gateway"}

	ServiceUnavailable = err{StatusCode: 503, Message: "Service Unavailable"}

	GatewayTimeout = err{StatusCode: 504, Message: "Gateway Timeout"}

	HTTPVersionNotSupported = err{StatusCode: 505, Message: "HTTP Version Not Supported"}
)
