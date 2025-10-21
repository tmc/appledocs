// Code generated from Apple documentation for ServiceManagement. DO NOT EDIT.

package servicemanagement

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SMAppService] class.
var (
	SMAppServiceClass     _SMAppServiceClass
	SMAppServiceClassOnce sync.Once
)

func getSMAppServiceClass() _SMAppServiceClass {
	SMAppServiceClassOnce.Do(func() {
		SMAppServiceClass = _SMAppServiceClass{objc.GetClass("SMAppService")}
	})
	return SMAppServiceClass
}

type _SMAppServiceClass struct {
	class objc.Class
}

// An interface definition for the [SMAppService] class.
type ISMAppService interface {
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
type SMAppService struct {
	objectivec.Object
}

// SMAppServiceFrom constructs a [SMAppService] from an unsafe.Pointer.
//
// An object the framework uses to control helper executables that live inside an app’s main bundle.
func SMAppServiceFrom(ptr unsafe.Pointer) SMAppService {
	return SMAppService{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SMAppServiceClass) Alloc() SMAppService {
	rv := objc.Send[SMAppService](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SMAppServiceClass) New() SMAppService {
	rv := objc.Send[SMAppService](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SMAppService) Init() SMAppService {
	rv := objc.Send[SMAppService](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SMAppService) Autorelease() SMAppService {
	rv := objc.Send[SMAppService](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSMAppService creates a new SMAppService instance.
func NewSMAppService() SMAppService {
	return getSMAppServiceClass().New()
}


// Initializes an app service object with a launch agent with the property list name you provide.
//
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/agent(plistName:)
func (sc _SMAppServiceClass) AgentServiceWithPlistName(plistName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("agentServiceWithPlistName:"), objc.String(plistName))
	return rv
}

// Initializes an app service object with a launch daemon with the property list name you provide.
//
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/daemon(plistName:)
func (sc _SMAppServiceClass) DaemonServiceWithPlistName(plistName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("daemonServiceWithPlistName:"), objc.String(plistName))
	return rv
}

// Initializes an app service object for a login item corresponding to the bundle with the identifier you provide.
//
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/loginItem(identifier:)
func (sc _SMAppServiceClass) LoginItemServiceWithIdentifier(identifier string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("loginItemServiceWithIdentifier:"), objc.String(identifier))
	return rv
}

// Opens System Settings to the Login Items control panel.
//
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/openSystemSettingsLoginItems()
func (sc _SMAppServiceClass) OpenSystemSettingsLoginItems() {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("openSystemSettingsLoginItems"))
}

// Check the authorization status of an earlier OS version login item.
//
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/statusForLegacyPlist(at:)
func (sc _SMAppServiceClass) StatusForLegacyURL(url unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("statusForLegacyURL:"), url)
	return rv
}

// An app service object that corresponds to the main application as a login item.
//
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/mainApp
func (sc _SMAppServiceClass) MainAppService() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("mainAppService"))
	return rv
}
// Registers the service so it can begin launching subject to user approval.
//
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/register()
func (s_ SMAppService) RegisterAndReturnError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("registerAndReturnError:"), error_)
	return rv
}

// Unregisters the service so the system no longer launches it.
//
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/unregister()
func (s_ SMAppService) UnregisterAndReturnError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("unregisterAndReturnError:"), error_)
	return rv
}

// Unregisters the service so the system no longer launches it and calls a completion handler you provide with the resulting error value.
//
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/unregister(completionHandler:)
func (s_ SMAppService) UnregisterWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("unregisterWithCompletionHandler:"), handler)
}

// An app service object that corresponds to the main application as a login item.
//
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/mainApp
func (s_ SMAppService) MainAppService() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("mainAppService"))
	return rv
}

// A property that describes registration or authorization state of the service.
//
// [Full Topic]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/status-swift.property
func (s_ SMAppService) Status() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("status"))
	return rv
}




