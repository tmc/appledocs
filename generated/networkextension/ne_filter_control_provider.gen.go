// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class NEFilterControlProvider */


/* debug [class_header]: Header for NEFilterControlProvider */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEFilterControlProvider */
// An interface definition for the [NEFilterControlProvider] class.
type INEFilterControlProvider interface {
	INEFilterProvider
	
/* debug [class_interface_properties]: Properties for NEFilterControlProvider */
	// properties:
	NEFilterProviderRemediationMapRemediationButtonTexts() objc.IObject /* cross-framework: NSString */
	NEFilterProviderRemediationMapRemediationURLs() objc.IObject /* cross-framework: NSString */
	NEFilterProviderRemediationURLFlowURL() objc.IObject /* cross-framework: NSString */
	SetNEFilterProviderRemediationURLFlowURL(value objc.IObject /* cross-framework: NSString */)
	NEFilterProviderRemediationURLFlowURLHostname() objc.IObject /* cross-framework: NSString */
	SetNEFilterProviderRemediationURLFlowURLHostname(value objc.IObject /* cross-framework: NSString */)
	NEFilterProviderRemediationURLOrganization() objc.IObject /* cross-framework: NSString */
	SetNEFilterProviderRemediationURLOrganization(value objc.IObject /* cross-framework: NSString */)
	NEFilterProviderRemediationURLUsername() objc.IObject /* cross-framework: NSString */
	SetNEFilterProviderRemediationURLUsername(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEFilterControlProvider */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEFilterControlProvider */
// Alloc allocates a new instance without initialization.
func (nc _NEFilterControlProviderClass) Alloc() NEFilterControlProvider {
	rv := objc.Send[NEFilterControlProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEFilterControlProvider */
// The principal class for a filter control provider extension.
//
// The Filter Control Provider’s primary responsibility is to provide information to the associated Filter Data Provider so that it can perform its task of accurately filtering network content. There are several ways in which the Filter Control Provider provides data to the associated Filter Data Provider: By writing information to disk. For example, the Filter Control Provider can maintain a database of filtering rules on disk in a location where the Filter Data Provider can read from the database. By defining a dictionary that maps keys to sets of customization parameters to be used when generating the block page. The Filter Data Provider gives the system the key for the desired customization parameters, and the system uses that key to get the customization parameters from the Filter Control Provider and generate the customized block page. By defining a dictionary that maps keys to strings to be appended to URLs. The Filter Data Provider gives the system the key for the string to be appended, and the system uses that key to get the string to be appended from the Filter Control Provider and appends the string to the URL.


// The principal class for a filter control provider extension.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEFilterControlProvider *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEFilterControlProvider */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEFilterControlProvider */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEFilterControlProvider */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEFilterControlProvider */

// A key in the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderremediationmapremediationbuttontexts
func (n_ NEFilterControlProvider) NEFilterProviderRemediationMapRemediationButtonTexts() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("NEFilterProviderRemediationMapRemediationButtonTexts"))
	return rv
}/* debug [instance_properties/getter]: NEFilterProviderRemediationMapRemediationButtonTexts */


// A key in the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderremediationmapremediationurls
func (n_ NEFilterControlProvider) NEFilterProviderRemediationMapRemediationURLs() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("NEFilterProviderRemediationMapRemediationURLs"))
	return rv
}/* debug [instance_properties/getter]: NEFilterProviderRemediationMapRemediationURLs */


// This string will be replaced with the full URL of the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderremediationurlflowurl
func (n_ NEFilterControlProvider) NEFilterProviderRemediationURLFlowURL() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("NEFilterProviderRemediationURLFlowURL"))
	return rv
}/* debug [instance_properties/getter]: NEFilterProviderRemediationURLFlowURL */


// This string will be replaced with the full URL of the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderremediationurlflowurl
func (n_ NEFilterControlProvider) SetNEFilterProviderRemediationURLFlowURL(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNEFilterProviderRemediationURLFlowURL:"), value)
}/* debug [instance_properties/setter]: NEFilterProviderRemediationURLFlowURL */


// This string will be replaced with the hostname portion of the flow’s URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderremediationurlflowurlhostname
func (n_ NEFilterControlProvider) NEFilterProviderRemediationURLFlowURLHostname() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("NEFilterProviderRemediationURLFlowURLHostname"))
	return rv
}/* debug [instance_properties/getter]: NEFilterProviderRemediationURLFlowURLHostname */


// This string will be replaced with the hostname portion of the flow’s URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderremediationurlflowurlhostname
func (n_ NEFilterControlProvider) SetNEFilterProviderRemediationURLFlowURLHostname(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNEFilterProviderRemediationURLFlowURLHostname:"), value)
}/* debug [instance_properties/setter]: NEFilterProviderRemediationURLFlowURLHostname */


// This string will be replaced with the value of the organization property set in the filter configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderremediationurlorganization
func (n_ NEFilterControlProvider) NEFilterProviderRemediationURLOrganization() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("NEFilterProviderRemediationURLOrganization"))
	return rv
}/* debug [instance_properties/getter]: NEFilterProviderRemediationURLOrganization */


// This string will be replaced with the value of the organization property set in the filter configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderremediationurlorganization
func (n_ NEFilterControlProvider) SetNEFilterProviderRemediationURLOrganization(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNEFilterProviderRemediationURLOrganization:"), value)
}/* debug [instance_properties/setter]: NEFilterProviderRemediationURLOrganization */


// This string will be replaced with the value of the username property set in the filter configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderremediationurlusername
func (n_ NEFilterControlProvider) NEFilterProviderRemediationURLUsername() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("NEFilterProviderRemediationURLUsername"))
	return rv
}/* debug [instance_properties/getter]: NEFilterProviderRemediationURLUsername */


// This string will be replaced with the value of the username property set in the filter configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderremediationurlusername
func (n_ NEFilterControlProvider) SetNEFilterProviderRemediationURLUsername(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNEFilterProviderRemediationURLUsername:"), value)
}/* debug [instance_properties/setter]: NEFilterProviderRemediationURLUsername */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEFilterControlProvider */


