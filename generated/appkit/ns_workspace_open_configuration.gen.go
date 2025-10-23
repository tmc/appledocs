// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	Activates() bool /* primitive/slice/pointer. */
	SetActivates(value bool /* primitive/slice/pointer. */)
	AddsToRecentItems() bool /* primitive/slice/pointer. */
	SetAddsToRecentItems(value bool /* primitive/slice/pointer. */)
	AllowsRunningApplicationSubstitution() bool /* primitive/slice/pointer. */
	SetAllowsRunningApplicationSubstitution(value bool /* primitive/slice/pointer. */)
	AppleEvent() AppleEventDescriptor /* not a class type */
	SetAppleEvent(value AppleEventDescriptor /* not a class type */)
	Architecture() unsafe.Pointer
	SetArchitecture(value unsafe.Pointer)
	Arguments() []string /* primitive/slice/pointer. */
	SetArguments(value []string /* primitive/slice/pointer. */)
	CreatesNewApplicationInstance() bool /* primitive/slice/pointer. */
	SetCreatesNewApplicationInstance(value bool /* primitive/slice/pointer. */)
	Environment() foundation.IDictionary /* already interface */
	SetEnvironment(value foundation.IDictionary /* already interface */)
	Hides() bool /* primitive/slice/pointer. */
	SetHides(value bool /* primitive/slice/pointer. */)
	HidesOthers() bool /* primitive/slice/pointer. */
	SetHidesOthers(value bool /* primitive/slice/pointer. */)
	ForPrinting() bool /* primitive/slice/pointer. */
	SetForPrinting(value bool /* primitive/slice/pointer. */)
	PromptsUserIfNeeded() bool /* primitive/slice/pointer. */
	SetPromptsUserIfNeeded(value bool /* primitive/slice/pointer. */)
	RequiresUniversalLinks() bool /* primitive/slice/pointer. */
	SetRequiresUniversalLinks(value bool /* primitive/slice/pointer. */)
	IsForPrinting() bool /* primitive/slice/pointer. */
	SetIsForPrinting(value bool /* primitive/slice/pointer. */)
	// methods:
}

// The configuration options for opening URLs or launching apps.
//
// Create an object before launching an app or opening a URL using the shared object. Use the properties of this object to customize the behavior of the launched app or the handling of the URLs. For example, you might tell the app to hide itself immediately after launch.


// The configuration options for opening URLs or launching apps.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspaceOpenConfiguration/configuration
func (wc _WorkspaceOpenConfigurationClass) Configuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(wc.class), objc.Sel("configuration"))
	return rv
}


// A Boolean value indicating whether the system activates the app and brings it to the foreground.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/activates
func (w_ WorkspaceOpenConfiguration) Activates() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("activates"))
	return rv
}


// A Boolean value indicating whether the system activates the app and brings it to the foreground.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/activates
func (w_ WorkspaceOpenConfiguration) SetActivates(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setActivates:"), value)
}


// A Boolean value indicating whether to add the app or documents to the Recent Items menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/addsToRecentItems
func (w_ WorkspaceOpenConfiguration) AddsToRecentItems() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("addsToRecentItems"))
	return rv
}


// A Boolean value indicating whether to add the app or documents to the Recent Items menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/addsToRecentItems
func (w_ WorkspaceOpenConfiguration) SetAddsToRecentItems(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAddsToRecentItems:"), value)
}


// A Boolean value that indicates whether to use a running instance of an application even if it’s at a different URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/allowsRunningApplicationSubstitution
func (w_ WorkspaceOpenConfiguration) AllowsRunningApplicationSubstitution() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsRunningApplicationSubstitution"))
	return rv
}


// A Boolean value that indicates whether to use a running instance of an application even if it’s at a different URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/allowsRunningApplicationSubstitution
func (w_ WorkspaceOpenConfiguration) SetAllowsRunningApplicationSubstitution(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsRunningApplicationSubstitution:"), value)
}


// The first Apple event to send to the new app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/appleEvent
func (w_ WorkspaceOpenConfiguration) AppleEvent() AppleEventDescriptor /* not a class type */ {
	rv := objc.Send[AppleEventDescriptor](w_.ID, objc.Sel("appleEvent"))
	return rv
}


// The first Apple event to send to the new app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/appleEvent
func (w_ WorkspaceOpenConfiguration) SetAppleEvent(value AppleEventDescriptor /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAppleEvent:"), value)
}


// The architecture version of the app to launch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/architecture
func (w_ WorkspaceOpenConfiguration) Architecture() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("architecture"))
	return rv
}


// The architecture version of the app to launch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/architecture
func (w_ WorkspaceOpenConfiguration) SetArchitecture(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setArchitecture:"), value)
}


// The set of command-line arguments to pass to a new app instance at launch time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/arguments
func (w_ WorkspaceOpenConfiguration) Arguments() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](w_.ID, objc.Sel("arguments"))
	return rv
}


// The set of command-line arguments to pass to a new app instance at launch time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/arguments
func (w_ WorkspaceOpenConfiguration) SetArguments(value []string /* primitive/slice/pointer. */) {
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/createsNewApplicationInstance
func (w_ WorkspaceOpenConfiguration) CreatesNewApplicationInstance() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("createsNewApplicationInstance"))
	return rv
}


// A Boolean value indicating whether you want the system to launch a new instance of the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/createsNewApplicationInstance
func (w_ WorkspaceOpenConfiguration) SetCreatesNewApplicationInstance(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCreatesNewApplicationInstance:"), value)
}


// The set of environment variables to set in a new app instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/environment
func (w_ WorkspaceOpenConfiguration) Environment() foundation.IDictionary /* already interface */ {
	rv := objc.Send[foundation.IDictionary](w_.ID, objc.Sel("environment"))
	return rv
}


// The set of environment variables to set in a new app instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/environment
func (w_ WorkspaceOpenConfiguration) SetEnvironment(value foundation.IDictionary /* already interface */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setEnvironment:"), value)
}


// A Boolean value indicating whether you want the app to hide itself after it launches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/hides
func (w_ WorkspaceOpenConfiguration) Hides() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("hides"))
	return rv
}


// A Boolean value indicating whether you want the app to hide itself after it launches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/hides
func (w_ WorkspaceOpenConfiguration) SetHides(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHides:"), value)
}


// A Boolean value indicating whether you want to hide all apps except the one that launched.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/hidesOthers
func (w_ WorkspaceOpenConfiguration) HidesOthers() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("hidesOthers"))
	return rv
}


// A Boolean value indicating whether you want to hide all apps except the one that launched.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/hidesOthers
func (w_ WorkspaceOpenConfiguration) SetHidesOthers(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHidesOthers:"), value)
}


// A Boolean value indicating whether you want to print the contents of documents and URLs instead of opening them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/isForPrinting
func (w_ WorkspaceOpenConfiguration) ForPrinting() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("forPrinting"))
	return rv
}


// A Boolean value indicating whether you want to print the contents of documents and URLs instead of opening them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/isForPrinting
func (w_ WorkspaceOpenConfiguration) SetForPrinting(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setForPrinting:"), value)
}


// A Boolean value indicating whether to display errors, authentication requests, or other UI elements to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/promptsUserIfNeeded
func (w_ WorkspaceOpenConfiguration) PromptsUserIfNeeded() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("promptsUserIfNeeded"))
	return rv
}


// A Boolean value indicating whether to display errors, authentication requests, or other UI elements to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/promptsUserIfNeeded
func (w_ WorkspaceOpenConfiguration) SetPromptsUserIfNeeded(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPromptsUserIfNeeded:"), value)
}


// A Boolean value indicating whether you require the URL to have an associated universal link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/requiresUniversalLinks
func (w_ WorkspaceOpenConfiguration) RequiresUniversalLinks() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("requiresUniversalLinks"))
	return rv
}


// A Boolean value indicating whether you require the URL to have an associated universal link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/requiresUniversalLinks
func (w_ WorkspaceOpenConfiguration) SetRequiresUniversalLinks(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setRequiresUniversalLinks:"), value)
}


// A Boolean value indicating whether you want to print the contents of documents and URLs instead of opening them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/openconfiguration/isforprinting
func (w_ WorkspaceOpenConfiguration) IsForPrinting() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("isForPrinting"))
	return rv
}


// A Boolean value indicating whether you want to print the contents of documents and URLs instead of opening them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/openconfiguration/isforprinting
func (w_ WorkspaceOpenConfiguration) SetIsForPrinting(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsForPrinting:"), value)
}



