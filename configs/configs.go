package configs

const (
	ACCEPT        = "Accept"
	AUTHORIZATION = "Authorization"
	XCTO          = "X-Content-Type-Options"
	XCTO_VALUE    = "nosniff"
	HSTS          = "Strict-Transport-Security"
	HSTS_VALUE    = "max-age=31536000; includeSubDomains"
	CC            = "Cache-Control"
	CC_VALUE      = "no-store"
	ACAO          = "Access-Control-Allow-Origin"
	ACAO_VALUE    = "*"
	ACAM          = "Access-Control-Allow-Methods"
	ACAM_VALUE    = "GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS"
	ACAH          = "Access-Control-Allow-Headers"
	ACAH_VALUE    = "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, Origin, Cookie, Timestamp, Grpc-Metadata-Timestamp, Grpc-Metadata-Client, Grpc-Metadata-Secret, Grpc-Metadata-Device"
	ACAC          = "Access-Control-Allow-Credentials"
	ACAC_VALUE    = "false"
	GRPC_METHOD   = "Grpc-Metadata-Method"

	// MAX 10 MB REQUEST AND RESPONSE GRPC
	MAX_SIZE_GRPC = 1024 * 1024 * 50

	// HTTP
	CONTENT_TYPE     = "Content-Type"
	APPLICATION_JSON = "application/json"

	IdentifierId = "identifierId"

	HTTPGet   = "GET"
	HTTPPost  = "POST"
	HTTPPut   = "PUT"
	HTTPPatch = "PATCH"
	HTTPDel   = "DELETE"
)
