// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NEFilterDataProvider] class.
var (
	NEFilterDataProviderClass     _NEFilterDataProviderClass
	NEFilterDataProviderClassOnce sync.Once
)

func getNEFilterDataProviderClass() _NEFilterDataProviderClass {
	NEFilterDataProviderClassOnce.Do(func() {
		NEFilterDataProviderClass = _NEFilterDataProviderClass{objc.GetClass("NEFilterDataProvider")}
	})
	return NEFilterDataProviderClass
}

type _NEFilterDataProviderClass struct {
	class objc.Class
}

// An interface definition for the [NEFilterDataProvider] class.
type INEFilterDataProvider interface {
	INEFilterProvider
	ApplySettingsCompletionHandler(settings unsafe.Pointer, completionHandler unsafe.Pointer)
	HandleNewFlow(flow unsafe.Pointer) unsafe.Pointer
	ResumeFlowWithVerdict(flow unsafe.Pointer, verdict unsafe.Pointer)
	UpdateFlowUsingVerdictForDirection(flow unsafe.Pointer, verdict unsafe.Pointer, direction unsafe.Pointer)
}

// The principal class for a filter data provider extension.
//
// Network content is delivered to the Filter Data Provider in the form of objects. Each object corresponds to a network connection opened by an application running on the device. The Filter Data Provider can choose to pass or block the data when it receives a new flow, or it can ask the system to see more of the flow’s data in either the outbound or inbound direction before making a pass or block decision. In addition to passing or blocking network data, the Filter Data Provider can tell the system that it needs more information before it can make a decision about a particular flow of data. The system will then ask the Filter Control Provider to update the current set of rules and place them in a location on disk that is readable from the Filter Data Provider extension. When a object is originated from a WebKit browser object, the Filter Data Provider can affect the user experience in the following ways: If the Filter Data Provider chooses to block the web page, then a special “block” page is displayed in the WebKit browser object informing the user that their attempt to access the content was blocked. The Filter Data Provider can choose to add a link to this block page, giving the user the option of requesting access to the content. If the Filter Data Provider chooses to allow the web page, then it can also specify that a string be appended to the web page URL. This allows the Filter Data Provider to direct the WebKit browser object to a “safe” version of the web page. To protect the user’s privacy, the Filter Data Provider extension sandbox prevents the extension from moving network content outside of its address space.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataProvider
type NEFilterDataProvider struct {
	NEFilterProvider
}

// NEFilterDataProviderFrom constructs a [NEFilterDataProvider] from an unsafe.Pointer.
//
// The principal class for a filter data provider extension.
func NEFilterDataProviderFrom(ptr unsafe.Pointer) NEFilterDataProvider {
	return NEFilterDataProvider{
		NEFilterProvider: NEFilterProviderFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NEFilterDataProviderClass) Alloc() NEFilterDataProvider {
	rv := objc.Send[NEFilterDataProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEFilterDataProviderClass) New() NEFilterDataProvider {
	rv := objc.Send[NEFilterDataProvider](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterDataProvider) Init() NEFilterDataProvider {
	rv := objc.Send[NEFilterDataProvider](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterDataProvider) Autorelease() NEFilterDataProvider {
	rv := objc.Send[NEFilterDataProvider](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterDataProvider creates a new NEFilterDataProvider instance.
func NewNEFilterDataProvider() NEFilterDataProvider {
	return getNEFilterDataProviderClass().New()
}


// Applies a set of filtering rules associated with the provider and changes the default filtering action.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataProvider/apply(_:completionHandler:)
func (n_ NEFilterDataProvider) ApplySettingsCompletionHandler(settings unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("applySettings:completionHandler:"), settings, completionHandler)
}

// Make a filtering decision for a newly-created flow of network content.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataProvider/handleNewFlow(_:)
func (n_ NEFilterDataProvider) HandleNewFlow(flow unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("handleNewFlow:"), flow)
	return rv
}

// Resumes a previously-paused flow.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataProvider/resumeFlow(_:with:)
func (n_ NEFilterDataProvider) ResumeFlowWithVerdict(flow unsafe.Pointer, verdict unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("resumeFlow:withVerdict:"), flow, verdict)
}

// Updates the verdict for a flow outside the context of any filter data provider callback.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataProvider/update(_:using:for:)
func (n_ NEFilterDataProvider) UpdateFlowUsingVerdictForDirection(flow unsafe.Pointer, verdict unsafe.Pointer, direction unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("updateFlow:usingVerdict:forDirection:"), flow, verdict, direction)
}



