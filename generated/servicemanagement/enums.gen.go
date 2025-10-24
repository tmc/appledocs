// Code generated from Apple documentation for ServiceManagement. DO NOT EDIT.

package servicemanagement

/* debug [enums.gen.go]: Generating 1 enums for ServiceManagement */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum SMAppServiceStatus (4 cases) */
// SMAppServiceStatus - Constants that describe the registration or authorization status of a helper executable.
//
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/Status-swift.enum
type SMAppServiceStatus uint

const (
	// SMAppServiceStatusEnabled - The service has been successfully registered and is eligible to run.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/Status-swift.enum/enabled
	SMAppServiceStatusEnabled SMAppServiceStatus = 0
	// SMAppServiceStatusNotFound - An error occurred and the framework couldn’t find this service.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/Status-swift.enum/notFound
	SMAppServiceStatusNotFound SMAppServiceStatus = 0
	// SMAppServiceStatusNotRegistered - The service hasn’t registered with the Service Management framework, or the service attempted to reregister after it was already registered.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/Status-swift.enum/notRegistered
	SMAppServiceStatusNotRegistered SMAppServiceStatus = 0
	// SMAppServiceStatusRequiresApproval - The service has been successfully registered, but the user needs to take action in System Preferences.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/Status-swift.enum/requiresApproval
	SMAppServiceStatusRequiresApproval SMAppServiceStatus = 0
)
