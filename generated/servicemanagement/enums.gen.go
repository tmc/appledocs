// Code generated from Apple documentation for ServiceManagement. DO NOT EDIT.

package servicemanagement

// Enum types and constants
// SMAppServiceStatus - Constants that describe the registration or authorization status of a helper executable.
//
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/Status-swift.enum
type AppServiceStatus uint

const (
	// AppServiceStatusEnabled - The service has been successfully registered and is eligible to run.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/Status-swift.enum/enabled
	AppServiceStatusEnabled AppServiceStatus = 0
	// AppServiceStatusNotFound - An error occurred and the framework couldn’t find this service.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/Status-swift.enum/notFound
	AppServiceStatusNotFound AppServiceStatus = 0
	// AppServiceStatusNotRegistered - The service hasn’t registered with the Service Management framework, or the service attempted to reregister after it was already registered.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/Status-swift.enum/notRegistered
	AppServiceStatusNotRegistered AppServiceStatus = 0
	// AppServiceStatusRequiresApproval - The service has been successfully registered, but the user needs to take action in System Preferences.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/Status-swift.enum/requiresApproval
	AppServiceStatusRequiresApproval AppServiceStatus = 0
)


