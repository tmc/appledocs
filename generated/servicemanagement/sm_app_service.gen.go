// Code generated from Apple documentation for ServiceManagement. DO NOT EDIT.

package servicemanagement

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [AppService] class.
type IAppService interface {
	objectivec.IObject
	RegisterAndReturnError(error_ unsafe.Pointer) bool
	UnregisterAndReturnError(error_ unsafe.Pointer) bool
	UnregisterWithCompletionHandler(handler unsafe.Pointer)
}

// An object the framework uses to control helper executables that live inside an app’s main bundle.
//
// In macOS 13 and later, use to register and control , , and as helper executables for your app. When converting code from earlier versions of macOS, use an object and select one of the following methods depending on the type of service your helper executable provides: For initialized as , the and APIs provide a replacement for . For initialized as , the and methods provide a replacement for installing property lists in or . For initialized as , the and methods provide a replacement for installing property lists in .
//
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

// Alloc allocates a new instance without initialization.
func (ac _AppServiceClass) Alloc() AppService {
	rv := objc.Send[AppService](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Initializes an app service object with a launch agent with the property list name you provide.
//
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/agent(plistName:)
func (ac _AppServiceClass) AgentServiceWithPlistName(plistName appkit.string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("agentServiceWithPlistName:"), plistName)
	return rv
}

// Initializes an app service object with a launch daemon with the property list name you provide.
//
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/daemon(plistName:)
func (ac _AppServiceClass) DaemonServiceWithPlistName(plistName appkit.string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("daemonServiceWithPlistName:"), plistName)
	return rv
}

// Initializes an app service object for a login item corresponding to the bundle with the identifier you provide.
//
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/loginItem(identifier:)
func (ac _AppServiceClass) LoginItemServiceWithIdentifier(identifier appkit.string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("loginItemServiceWithIdentifier:"), identifier)
	return rv
}

// Opens System Settings to the Login Items control panel.
//
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/openSystemSettingsLoginItems()
func (ac _AppServiceClass) OpenSystemSettingsLoginItems() {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("openSystemSettingsLoginItems"))
}

// Check the authorization status of an earlier OS version login item.
//
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/statusForLegacyPlist(at:)
func (ac _AppServiceClass) StatusForLegacyURL(url foundation.IURL) AppServiceStatus {
	rv := objc.Send[AppServiceStatus](objc.ID(ac.class), objc.Sel("statusForLegacyURL:"), url)
	return rv
}

// An app service object that corresponds to the main application as a login item.
//
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/mainApp
func (ac _AppServiceClass) MainAppService() AppService {
	rv := objc.Send[SMAppService](objc.ID(ac.class), objc.Sel("mainAppService"))
	return rv
}
// Registers the service so it can begin launching subject to user approval.
//
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/register()
func (a_ AppService) RegisterAndReturnError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("registerAndReturnError:"), error_)
	return rv
}

// Unregisters the service so the system no longer launches it.
//
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/unregister()
func (a_ AppService) UnregisterAndReturnError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("unregisterAndReturnError:"), error_)
	return rv
}

// Unregisters the service so the system no longer launches it and calls a completion handler you provide with the resulting error value.
//
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/unregister(completionHandler:)
func (a_ AppService) UnregisterWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("unregisterWithCompletionHandler:"), handler)
}

// An app service object that corresponds to the main application as a login item.
//
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/mainApp
func (a_ AppService) MainAppService() SMAppService {
	rv := objc.Send[SMAppService](a_.ID, objc.Sel("mainAppService"))
	return rv
}

// A property that describes registration or authorization state of the service.
//
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/status-swift.property
func (a_ AppService) Status() AppServiceStatus {
	rv := objc.Send[AppServiceStatus](a_.ID, objc.Sel("status"))
	return rv
}




