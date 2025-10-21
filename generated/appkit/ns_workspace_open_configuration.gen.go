// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [WorkspaceOpenConfiguration] class.
var (
	WorkspaceOpenConfigurationClass     _WorkspaceOpenConfigurationClass
	WorkspaceOpenConfigurationClassOnce sync.Once
)

func getWorkspaceOpenConfigurationClass() _WorkspaceOpenConfigurationClass {
	WorkspaceOpenConfigurationClassOnce.Do(func() {
		WorkspaceOpenConfigurationClass = _WorkspaceOpenConfigurationClass{objc.GetClass("NSWorkspaceOpenConfiguration")}
	})
	return WorkspaceOpenConfigurationClass
}

type _WorkspaceOpenConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [WorkspaceOpenConfiguration] class.
type IWorkspaceOpenConfiguration interface {
	objectivec.IObject
}

// The configuration options for opening URLs or launching apps.
//
// Create an object before launching an app or opening a URL using the shared object. Use the properties of this object to customize the behavior of the launched app or the handling of the URLs. For example, you might tell the app to hide itself immediately after launch.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration
type WorkspaceOpenConfiguration struct {
	objectivec.Object
}

// WorkspaceOpenConfigurationFrom constructs a [WorkspaceOpenConfiguration] from an unsafe.Pointer.
//
// The configuration options for opening URLs or launching apps.
func WorkspaceOpenConfigurationFrom(ptr unsafe.Pointer) WorkspaceOpenConfiguration {
	return WorkspaceOpenConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (wc _WorkspaceOpenConfigurationClass) Alloc() WorkspaceOpenConfiguration {
	rv := objc.Send[WorkspaceOpenConfiguration](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _WorkspaceOpenConfigurationClass) New() WorkspaceOpenConfiguration {
	rv := objc.Send[WorkspaceOpenConfiguration](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WorkspaceOpenConfiguration) Init() WorkspaceOpenConfiguration {
	rv := objc.Send[WorkspaceOpenConfiguration](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WorkspaceOpenConfiguration) Autorelease() WorkspaceOpenConfiguration {
	rv := objc.Send[WorkspaceOpenConfiguration](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWorkspaceOpenConfiguration creates a new WorkspaceOpenConfiguration instance.
func NewWorkspaceOpenConfiguration() WorkspaceOpenConfiguration {
	return getWorkspaceOpenConfigurationClass().New()
}


// Creates and returns a new workspace configuration object containing default values.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspaceOpenConfiguration/configuration
func (wc _WorkspaceOpenConfigurationClass) Configuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(wc.class), objc.Sel("configuration"))
	return rv
}

// A Boolean value indicating whether the system activates the app and brings it to the foreground.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/activates
func (w_ WorkspaceOpenConfiguration) Activates() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("activates"))
	return rv
}


// SetActivates sets the value of the activates property.
// A Boolean value indicating whether the system activates the app and brings it to the foreground.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/activates
func (w_ WorkspaceOpenConfiguration) SetActivates(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setActivates:"), value)
}

// A Boolean value indicating whether to add the app or documents to the Recent Items menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/addsToRecentItems
func (w_ WorkspaceOpenConfiguration) AddsToRecentItems() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("addsToRecentItems"))
	return rv
}


// SetAddsToRecentItems sets the value of the addsToRecentItems property.
// A Boolean value indicating whether to add the app or documents to the Recent Items menu.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/addsToRecentItems
func (w_ WorkspaceOpenConfiguration) SetAddsToRecentItems(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAddsToRecentItems:"), value)
}

// A Boolean value that indicates whether to use a running instance of an application even if it’s at a different URL.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/allowsRunningApplicationSubstitution
func (w_ WorkspaceOpenConfiguration) AllowsRunningApplicationSubstitution() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsRunningApplicationSubstitution"))
	return rv
}


// SetAllowsRunningApplicationSubstitution sets the value of the allowsRunningApplicationSubstitution property.
// A Boolean value that indicates whether to use a running instance of an application even if it’s at a different URL.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/allowsRunningApplicationSubstitution
func (w_ WorkspaceOpenConfiguration) SetAllowsRunningApplicationSubstitution(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsRunningApplicationSubstitution:"), value)
}

// The architecture version of the app to launch.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/architecture
func (w_ WorkspaceOpenConfiguration) Architecture() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("architecture"))
	return rv
}


// SetArchitecture sets the value of the architecture property.
// The architecture version of the app to launch.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/architecture
func (w_ WorkspaceOpenConfiguration) SetArchitecture(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setArchitecture:"), value)
}

// The set of command-line arguments to pass to a new app instance at launch time.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/arguments
func (w_ WorkspaceOpenConfiguration) Arguments() []string {
	rv := objc.Send[[]string](w_.ID, objc.Sel("arguments"))
	return rv
}


// SetArguments sets the value of the arguments property.
// The set of command-line arguments to pass to a new app instance at launch time.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/arguments
func (w_ WorkspaceOpenConfiguration) SetArguments(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](w_.ID, objc.Sel("setArguments:"), nsArray)
}

// A Boolean value indicating whether you want the system to launch a new instance of the app.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/createsNewApplicationInstance
func (w_ WorkspaceOpenConfiguration) CreatesNewApplicationInstance() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("createsNewApplicationInstance"))
	return rv
}


// SetCreatesNewApplicationInstance sets the value of the createsNewApplicationInstance property.
// A Boolean value indicating whether you want the system to launch a new instance of the app.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/createsNewApplicationInstance
func (w_ WorkspaceOpenConfiguration) SetCreatesNewApplicationInstance(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCreatesNewApplicationInstance:"), value)
}

// The set of environment variables to set in a new app instance.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/environment
func (w_ WorkspaceOpenConfiguration) Environment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("environment"))
	return rv
}


// SetEnvironment sets the value of the environment property.
// The set of environment variables to set in a new app instance.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/environment
func (w_ WorkspaceOpenConfiguration) SetEnvironment(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setEnvironment:"), value)
}

// A Boolean value indicating whether you want the app to hide itself after it launches.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/hides
func (w_ WorkspaceOpenConfiguration) Hides() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hides"))
	return rv
}


// SetHides sets the value of the hides property.
// A Boolean value indicating whether you want the app to hide itself after it launches.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/hides
func (w_ WorkspaceOpenConfiguration) SetHides(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHides:"), value)
}

// A Boolean value indicating whether you want to hide all apps except the one that launched.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/hidesOthers
func (w_ WorkspaceOpenConfiguration) HidesOthers() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hidesOthers"))
	return rv
}


// SetHidesOthers sets the value of the hidesOthers property.
// A Boolean value indicating whether you want to hide all apps except the one that launched.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/hidesOthers
func (w_ WorkspaceOpenConfiguration) SetHidesOthers(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHidesOthers:"), value)
}

// A Boolean value indicating whether you want to print the contents of documents and URLs instead of opening them.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/isForPrinting
func (w_ WorkspaceOpenConfiguration) ForPrinting() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("forPrinting"))
	return rv
}


// SetForPrinting sets the value of the forPrinting property.
// A Boolean value indicating whether you want to print the contents of documents and URLs instead of opening them.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/isForPrinting
func (w_ WorkspaceOpenConfiguration) SetForPrinting(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setForPrinting:"), value)
}

// A Boolean value indicating whether you require the URL to have an associated universal link.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/requiresUniversalLinks
func (w_ WorkspaceOpenConfiguration) RequiresUniversalLinks() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("requiresUniversalLinks"))
	return rv
}


// SetRequiresUniversalLinks sets the value of the requiresUniversalLinks property.
// A Boolean value indicating whether you require the URL to have an associated universal link.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/requiresUniversalLinks
func (w_ WorkspaceOpenConfiguration) SetRequiresUniversalLinks(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setRequiresUniversalLinks:"), value)
}

// The first Apple event to send to the new app.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/openconfiguration/appleevent
func (w_ WorkspaceOpenConfiguration) AppleEvent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("appleEvent"))
	return rv
}


// SetAppleEvent sets the value of the appleEvent property.
// The first Apple event to send to the new app.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/openconfiguration/appleevent
func (w_ WorkspaceOpenConfiguration) SetAppleEvent(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAppleEvent:"), value)
}

// A Boolean value indicating whether you want to print the contents of documents and URLs instead of opening them.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/openconfiguration/isforprinting
func (w_ WorkspaceOpenConfiguration) IsForPrinting() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isForPrinting"))
	return rv
}


// SetIsForPrinting sets the value of the isForPrinting property.
// A Boolean value indicating whether you want to print the contents of documents and URLs instead of opening them.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/openconfiguration/isforprinting
func (w_ WorkspaceOpenConfiguration) SetIsForPrinting(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsForPrinting:"), value)
}

// A Boolean value indicating whether to display errors, authentication requests, or other UI elements to the user.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/openconfiguration/promptsuserifneeded
func (w_ WorkspaceOpenConfiguration) PromptsUserIfNeeded() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("promptsUserIfNeeded"))
	return rv
}


// SetPromptsUserIfNeeded sets the value of the promptsUserIfNeeded property.
// A Boolean value indicating whether to display errors, authentication requests, or other UI elements to the user.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/openconfiguration/promptsuserifneeded
func (w_ WorkspaceOpenConfiguration) SetPromptsUserIfNeeded(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPromptsUserIfNeeded:"), value)
}



