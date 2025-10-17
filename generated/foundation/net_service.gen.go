// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [NetService] class.
var NetServiceClass objc.Class

func init() {
	NetServiceClass = objc.GetClass("NSNetService")
}

type NetService struct {
	objc.ID
}

func NetServiceFrom(ptr unsafe.Pointer) NetService {
	return NetService{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc NetService) Alloc() NetService {
	ret := objc.ID(NetServiceClass).Send(objc.RegisterName("alloc"))
	return NetService{ret}
}

// Init initializes the instance.
func (n_ NetService) Init() NetService {
	ret := n_.ID.Send(objc.RegisterName("init"))
	return NetService{ret}
}
// Returns the receiver, initialized as a network service of a given type and sets the initial host information. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NetService/init(domain:type:name:)
func NewNetServiceWithDomainTypeName(domain string, type_ string, name string) NetService {
	instance := NetService{}.Alloc()
	sel := objc.RegisterName("initWithDomain:type:name:")
	ret := instance.ID.Send(sel, domain, type_, name)
	instance = NetService{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes the receiver for publishing a network service of type   at the socket location specified by  ,  , and  . [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NetService/init(domain:type:name:port:)
func NewNetServiceWithDomainTypeNamePort(domain string, type_ string, name string, port int) NetService {
	instance := NetService{}.Alloc()
	sel := objc.RegisterName("initWithDomain:type:name:port:")
	ret := instance.ID.Send(sel, domain, type_, name, port)
	instance = NetService{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}


// Returns an   object representing a TXT record formed from a given dictionary. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NetService/data(fromTXTRecord:)
func (nc NetService) DataFromTXTRecordDictionary(txtDictionary unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("dataFromTXTRecordDictionary:")
	ret := objc.ID(NetServiceClass).Send(sel, txtDictionary)
	return unsafe.Pointer(ret)
}
// Returns a dictionary representing a TXT record given as an   object. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NetService/dictionary(fromTXTRecord:)
func (nc NetService) DictionaryFromTXTRecordData(txtData unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("dictionaryFromTXTRecordData:")
	ret := objc.ID(NetServiceClass).Send(sel, txtData)
	return unsafe.Pointer(ret)
}
// Creates a pair of input and output streams for the receiver and returns a Boolean value that indicates whether they were retrieved successfully. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NetService/getInputStream(_:outputStream:)
func (n_ NetService) GetInputStreamOutputStream(inputStream unsafe.Pointer, outputStream unsafe.Pointer) bool {
	sel := objc.RegisterName("getInputStream:outputStream:")
	ret := n_.ID.Send(sel, inputStream, outputStream)
	return ret != 0
}
// Attempts to advertise the receiver’s on the network. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NetService/publish()
func (n_ NetService) Publish() {
	sel := objc.RegisterName("publish")
	n_.ID.Send(sel)
}
// Attempts to advertise the receiver on the network, with the given options. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NetService/publish(options:)
func (n_ NetService) PublishWithOptions(options unsafe.Pointer) {
	sel := objc.RegisterName("publishWithOptions:")
	n_.ID.Send(sel, options)
}
// Removes the service from the given run loop for a given mode. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NetService/remove(from:forMode:)
func (n_ NetService) RemoveFromRunLoopForMode(aRunLoop unsafe.Pointer, mode unsafe.Pointer) {
	sel := objc.RegisterName("removeFromRunLoop:forMode:")
	n_.ID.Send(sel, aRunLoop, mode)
}
// Starts a resolve process for the service. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NetService/resolve()
func (n_ NetService) Resolve() {
	sel := objc.RegisterName("resolve")
	n_.ID.Send(sel)
}
// Starts a resolve process of a finite duration for the service. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NetService/resolve(withTimeout:)
func (n_ NetService) ResolveWithTimeout(timeout foundation.TimeInterval) {
	sel := objc.RegisterName("resolveWithTimeout:")
	n_.ID.Send(sel, timeout)
}
// Adds the service to the specified run loop. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NetService/schedule(in:forMode:)
func (n_ NetService) ScheduleInRunLoopForMode(aRunLoop unsafe.Pointer, mode unsafe.Pointer) {
	sel := objc.RegisterName("scheduleInRunLoop:forMode:")
	n_.ID.Send(sel, aRunLoop, mode)
}
// Sets the TXT record for the receiver, and returns a Boolean value that indicates whether the operation was successful. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NetService/setTXTRecord(_:)
func (n_ NetService) SetTXTRecordData(recordData unsafe.Pointer) bool {
	sel := objc.RegisterName("setTXTRecordData:")
	ret := n_.ID.Send(sel, recordData)
	return ret != 0
}
// Starts the monitoring of TXT-record updates for the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NetService/startMonitoring()
func (n_ NetService) StartMonitoring() {
	sel := objc.RegisterName("startMonitoring")
	n_.ID.Send(sel)
}
// Halts a currently running attempt to publish or resolve a service. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NetService/stop()
func (n_ NetService) Stop() {
	sel := objc.RegisterName("stop")
	n_.ID.Send(sel)
}
// Stops the monitoring of TXT-record updates for the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NetService/stopMonitoring()
func (n_ NetService) StopMonitoring() {
	sel := objc.RegisterName("stopMonitoring")
	n_.ID.Send(sel)
}
// Returns the TXT record for the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NetService/txtRecordData()
func (n_ NetService) TXTRecordData() unsafe.Pointer {
	sel := objc.RegisterName("TXTRecordData")
	ret := n_.ID.Send(sel)
	return unsafe.Pointer(ret)
}

