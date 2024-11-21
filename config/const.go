package config

// ErrExitStatus represents the error status in this application.
const ErrExitStatus int = 2

const (
	// AppConfigPath is the path of application.yml.
	AppConfigPath = "config/resources/application.%s.yml"
	// MessagesConfigPath is the path of messages.properties.
	MessagesConfigPath = "config/resources/messages.properties"
	// LoggerConfigPath is the path of zaplogger.yml.
	LoggerConfigPath = "config/resources/zaplogger.%s.yml"
)

// PasswordHashCost is hash cost for a password.
const PasswordHashCost int = 10

const (
	// API represents the group of API.
	API = "/api"
)

const (
	// APIUser represents the group of auth management API.
	APIUser = API + "/auth"
	// APIUserLoginStatus represents the API to get the status of logged in User.
	APIUserLoginStatus = APIUser + "/loginStatus"
	// APIUserLoginUser represents the API to get the logged in User.
	APIUserLoginUser = APIUser + "/loginUser"
	// APIUserLogin represents the API to login by session authentication.
	APIUserLogin = APIUser + "/login"
	// APIUserLogout represents the API to logout.
	APIUserLogout = APIUser + "/logout"
	// APIGetUserInfo represents the API to get the User data.
	APIGetUserInfo = APIUser + "/getUserInfo"
)

const (
	// APIHealth represents the API to get the status of this application.
	APIHealth = API + "/health"
)
