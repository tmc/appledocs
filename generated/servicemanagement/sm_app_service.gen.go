// Code generated from Apple documentation for ServiceManagement. DO NOT EDIT.

package servicemanagement

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SMAppService */

/* debug [class_header]: Header for SMAppService */
// The class instance for the [AppService] class.
var (
	AppServiceClass     _AppServiceClass
	AppServiceClassOnce sync.Once
)

func getAppServiceClass() _AppServiceClass {
	AppServiceClassOnce.Do(func() {
		AppServiceClass = _AppServiceClass{objc.GetClass("SMAppService")}
	})
	return AppServiceClass
}

type _AppServiceClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for AppService */
// An interface definition for the [AppService] class.
type IAppService interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for AppService */
	// properties:
	Status() AppServiceStatus
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for AppService */
	// methods:
	RegisterAndReturnError(error_ unsafe.Pointer) bool
	UnregisterAndReturnError(error_ unsafe.Pointer) bool
	UnregisterWithCompletionHandler(handler unsafe.Pointer)
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for AppService */
// Alloc allocates a new instance without initialization.
func (ac _AppServiceClass) Alloc() AppService {
	rv := objc.Send[AppService](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AppServiceClass) New() AppService {
	rv := objc.Send[AppService](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AppService) Init() AppService {
	rv := objc.Send[AppService](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AppService) Autorelease() AppService {
	rv := objc.Send[AppService](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAppService creates a new AppService instance.
func NewAppService() AppService {
	return getAppServiceClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for AppService */
// An object the framework uses to control helper executables that live inside an app’s main bundle.
//
// In macOS 13 and later, use to register and control , , and as helper executables for your app. When converting code from earlier versions of macOS, use an object and select one of the following methods depending on the type of service your helper executable provides: For initialized as , the and APIs provide a replacement for . For initialized as , the and methods provide a replacement for installing property lists in or . For initialized as , the and methods provide a replacement for installing property lists in .

// An object the framework uses to control helper executables that live inside an app’s main bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService
type AppService struct {
	objectivec.Object
}

// AppServiceFrom constructs a [AppService] from an unsafe.Pointer.
//
// An object the framework uses to control helper executables that live inside an app’s main bundle.
func AppServiceFrom(ptr unsafe.Pointer) AppService {
	return AppService{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for AppService */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for AppService */

// Initializes an app service object with a launch agent with the property list name you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/agent(plistName:)
func (ac _AppServiceClass) AgentServiceWithPlistName(plistName objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("agentServiceWithPlistName:"), plistName)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=AgentServiceWithPlistName) */

// Initializes an app service object with a launch daemon with the property list name you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/daemon(plistName:)
func (ac _AppServiceClass) DaemonServiceWithPlistName(plistName objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("daemonServiceWithPlistName:"), plistName)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=DaemonServiceWithPlistName) */

// Initializes an app service object for a login item corresponding to the bundle with the identifier you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/loginItem(identifier:)
func (ac _AppServiceClass) LoginItemServiceWithIdentifier(identifier objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("loginItemServiceWithIdentifier:"), identifier)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=LoginItemServiceWithIdentifier) */

// Opens System Settings to the Login Items control panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/openSystemSettingsLoginItems()
func (ac _AppServiceClass) OpenSystemSettingsLoginItems() {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("openSystemSettingsLoginItems"))
} /* debug [class_methods/method]: Class method for%!(EXTRA string=OpenSystemSettingsLoginItems) */

// Check the authorization status of an earlier OS version login item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/statusForLegacyPlist(at:)
func (ac _AppServiceClass) StatusForLegacyURL(url objc.IObject /* cross-framework: NSURL */) AppServiceStatus {
	rv := objc.Send[AppServiceStatus](objc.ID(ac.class), objc.Sel("statusForLegacyURL:"), url)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=StatusForLegacyURL) */

/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for AppService */

// An app service object that corresponds to the main application as a login item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/mainApp
func (ac _AppServiceClass) MainAppService() AppService {
	rv := objc.Send[AppService](objc.ID(ac.class), objc.Sel("mainAppService"))
	return rv
} /* debug [class_properties_class/property]: mainAppService */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for AppService */

// Registers the service so it can begin launching subject to user approval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/register()
func (a_ AppService) RegisterAndReturnError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("registerAndReturnError:"), error_)
	return rv
} /* debug [instance_methods/method]: RegisterAndReturnError */

// Unregisters the service so the system no longer launches it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/unregister()
func (a_ AppService) UnregisterAndReturnError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("unregisterAndReturnError:"), error_)
	return rv
} /* debug [instance_methods/method]: UnregisterAndReturnError */

// Unregisters the service so the system no longer launches it and calls a completion handler you provide with the resulting error value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/unregister(completionHandler:)
func (a_ AppService) UnregisterWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("unregisterWithCompletionHandler:"), handler)
} /* debug [instance_methods/method]: UnregisterWithCompletionHandler */

/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for AppService */

// An app service object that corresponds to the main application as a login item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/mainApp
func (a_ AppService) MainAppService() ISMAppService {
	rv := objc.Send[AppService](a_.ID, objc.Sel("mainAppService"))
	return rv
} /* debug [instance_properties/getter]: mainAppService */

// A property that describes registration or authorization state of the service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/status-swift.property
func (a_ AppService) Status() AppServiceStatus {
	rv := objc.Send[AppServiceStatus](a_.ID, objc.Sel("status"))
	return rv
} /* debug [instance_properties/getter]: status */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class SMAppService */
