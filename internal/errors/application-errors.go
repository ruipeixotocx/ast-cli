package applicationerrors

const (
	ApplicationDoesntExistOrNoPermission = "Provided application does not exist or user has no permission to the application"
)

const (
	FailedToGetApplication = "Failed to get application"
	RedirectURLNotFound    = "redirect URL not found in response"
	HTTPMethodNotFound     = "HTTP method not found in request"
)
