// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [NEFilterControlProvider] class.
var (
	NEFilterControlProviderClass     _NEFilterControlProviderClass
	NEFilterControlProviderClassOnce sync.Once
)

func getNEFilterControlProviderClass() _NEFilterControlProviderClass {
	NEFilterControlProviderClassOnce.Do(func() {
		NEFilterControlProviderClass = _NEFilterControlProviderClass{objc.GetClass("NEFilterControlProvider")}
	})
	return NEFilterControlProviderClass
}

type _NEFilterControlProviderClass struct {
	class objc.Class
}

// An interface definition for the [NEFilterControlProvider] class.
type INEFilterControlProvider interface {
	INEFilterProvider
}

// The principal class for a filter control provider extension.
//
// The Filter Control Provider’s primary responsibility is to provide information to the associated Filter Data Provider so that it can perform its task of accurately filtering network content. There are several ways in which the Filter Control Provider provides data to the associated Filter Data Provider: By writing information to disk. For example, the Filter Control Provider can maintain a database of filtering rules on disk in a location where the Filter Data Provider can read from the database. By defining a dictionary that maps keys to sets of customization parameters to be used when generating the block page. The Filter Data Provider gives the system the key for the desired customization parameters, and the system uses that key to get the customization parameters from the Filter Control Provider and generate the customized block page. By defining a dictionary that maps keys to strings to be appended to URLs. The Filter Data Provider gives the system the key for the string to be appended, and the system uses that key to get the string to be appended from the Filter Control Provider and appends the string to the URL.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterControlProvider
type NEFilterControlProvider struct {
	NEFilterProvider
}

// NEFilterControlProviderFrom constructs a [NEFilterControlProvider] from an unsafe.Pointer.
//
// The principal class for a filter control provider extension.
func NEFilterControlProviderFrom(ptr unsafe.Pointer) NEFilterControlProvider {
	return NEFilterControlProvider{
		NEFilterProvider: NEFilterProviderFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NEFilterControlProviderClass) Alloc() NEFilterControlProvider {
	rv := objc.Send[NEFilterControlProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEFilterControlProviderClass) New() NEFilterControlProvider {
	rv := objc.Send[NEFilterControlProvider](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterControlProvider) Init() NEFilterControlProvider {
	rv := objc.Send[NEFilterControlProvider](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterControlProvider) Autorelease() NEFilterControlProvider {
	rv := objc.Send[NEFilterControlProvider](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterControlProvider creates a new NEFilterControlProvider instance.
func NewNEFilterControlProvider() NEFilterControlProvider {
	return getNEFilterControlProviderClass().New()
}


// A dictionary containing sets of strings used to customize the remediation portion of the block page.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltercontrolprovider/remediationmap
func (n_ NEFilterControlProvider) RemediationMap() foundation.NSObject {
	rv := objc.Send[foundation.NSObject](n_.ID, objc.Sel("remediationMap"))
	return rv
}


// SetRemediationMap sets the value of the remediationMap property.
// A dictionary containing sets of strings used to customize the remediation portion of the block page.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltercontrolprovider/remediationmap
func (n_ NEFilterControlProvider) SetRemediationMap(value foundation.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setRemediationMap:"), value)
}

// A dictionary containing strings to be appended to URLs.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltercontrolprovider/urlappendstringmap
func (n_ NEFilterControlProvider) UrlAppendStringMap() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("urlAppendStringMap"))
	return rv
}


// SetUrlAppendStringMap sets the value of the urlAppendStringMap property.
// A dictionary containing strings to be appended to URLs.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefiltercontrolprovider/urlappendstringmap
func (n_ NEFilterControlProvider) SetUrlAppendStringMap(value appkit.string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUrlAppendStringMap:"), value)
}

// A key in the
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderremediationmapremediationbuttontexts
func (n_ NEFilterControlProvider) NEFilterProviderRemediationMapRemediationButtonTexts() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("NEFilterProviderRemediationMapRemediationButtonTexts"))
	return rv
}

// A key in the
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderremediationmapremediationurls
func (n_ NEFilterControlProvider) NEFilterProviderRemediationMapRemediationURLs() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("NEFilterProviderRemediationMapRemediationURLs"))
	return rv
}

// This string will be replaced with the full URL of the flow.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderremediationurlflowurl
func (n_ NEFilterControlProvider) NEFilterProviderRemediationURLFlowURL() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("NEFilterProviderRemediationURLFlowURL"))
	return rv
}


// SetNEFilterProviderRemediationURLFlowURL sets the value of the NEFilterProviderRemediationURLFlowURL property.
// This string will be replaced with the full URL of the flow.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderremediationurlflowurl
func (n_ NEFilterControlProvider) SetNEFilterProviderRemediationURLFlowURL(value appkit.string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNEFilterProviderRemediationURLFlowURL:"), value)
}

// This string will be replaced with the hostname portion of the flow’s URL.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderremediationurlflowurlhostname
func (n_ NEFilterControlProvider) NEFilterProviderRemediationURLFlowURLHostname() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("NEFilterProviderRemediationURLFlowURLHostname"))
	return rv
}


// SetNEFilterProviderRemediationURLFlowURLHostname sets the value of the NEFilterProviderRemediationURLFlowURLHostname property.
// This string will be replaced with the hostname portion of the flow’s URL.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderremediationurlflowurlhostname
func (n_ NEFilterControlProvider) SetNEFilterProviderRemediationURLFlowURLHostname(value appkit.string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNEFilterProviderRemediationURLFlowURLHostname:"), value)
}

// This string will be replaced with the value of the organization property set in the filter configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderremediationurlorganization
func (n_ NEFilterControlProvider) NEFilterProviderRemediationURLOrganization() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("NEFilterProviderRemediationURLOrganization"))
	return rv
}


// SetNEFilterProviderRemediationURLOrganization sets the value of the NEFilterProviderRemediationURLOrganization property.
// This string will be replaced with the value of the organization property set in the filter configuration.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderremediationurlorganization
func (n_ NEFilterControlProvider) SetNEFilterProviderRemediationURLOrganization(value appkit.string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNEFilterProviderRemediationURLOrganization:"), value)
}

// This string will be replaced with the value of the username property set in the filter configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderremediationurlusername
func (n_ NEFilterControlProvider) NEFilterProviderRemediationURLUsername() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("NEFilterProviderRemediationURLUsername"))
	return rv
}


// SetNEFilterProviderRemediationURLUsername sets the value of the NEFilterProviderRemediationURLUsername property.
// This string will be replaced with the value of the username property set in the filter configuration.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderremediationurlusername
func (n_ NEFilterControlProvider) SetNEFilterProviderRemediationURLUsername(value appkit.string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNEFilterProviderRemediationURLUsername:"), value)
}



