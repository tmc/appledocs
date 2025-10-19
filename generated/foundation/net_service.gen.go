// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NetService] class.
var (
	netServiceClass     _NetServiceClass
	netServiceClassOnce sync.Once
)

func getNetServiceClass() _NetServiceClass {
	netServiceClassOnce.Do(func() {
		netServiceClass = _NetServiceClass{objc.GetClass("NSNetService")}
	})
	return netServiceClass
}

type _NetServiceClass struct {
	class objc.Class
}

// An interface definition for the [NetService] class.
type INetService interface {
	objectivec.IObject
	GetInputStreamOutputStream(inputStream unsafe.Pointer, outputStream unsafe.Pointer) bool
	Publish()
	PublishWithOptions(options unsafe.Pointer)
	RemoveFromRunLoopForMode(aRunLoop unsafe.Pointer, mode unsafe.Pointer)
	Resolve()
	ResolveWithTimeout(timeout TimeInterval)
	ScheduleInRunLoopForMode(aRunLoop unsafe.Pointer, mode unsafe.Pointer)
	SetTXTRecordData(recordData unsafe.Pointer) bool
	StartMonitoring()
	Stop()
	StopMonitoring()
	TXTRecordData() unsafe.Pointer
}

// A network service that broadcasts its availability using multicast DNS.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService
type NetService struct {
	objectivec.Object
}

// NetServiceFrom constructs a [NetService] from an unsafe.Pointer.
//
// A network service that broadcasts its availability using multicast DNS.
func NetServiceFrom(ptr unsafe.Pointer) NetService {
	return NetService{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NetServiceClass) Alloc() NetService {
	rv := objc.Send[NetService](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NetServiceClass) New() NetService {
	rv := objc.Send[NetService](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NetService) Init() NetService {
	rv := objc.Send[NetService](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NetService) Autorelease() NetService {
	rv := objc.Send[NetService](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNetService creates a new NetService instance.
func NewNetService() NetService {
	return getNetServiceClass().New()
}


// Returns the receiver, initialized as a network service of a given type and sets the initial host information.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/init(domain:type:name:)
func NewNetServiceWithDomainTypeName(domain string, type_ string, name string) NetService {
	instance := getNetServiceClass().Alloc()
	rv := objc.Send[NetService](instance.ID, objc.Sel("initWithDomain:type:name:"), objc.String(domain), objc.String(type_), objc.String(name))
	rv.Autorelease()
	return rv
}
// Initializes the receiver for publishing a network service of type at the socket location specified by , , and .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/init(domain:type:name:port:)
func NewNetServiceWithDomainTypeNamePort(domain string, type_ string, name string, port int) NetService {
	instance := getNetServiceClass().Alloc()
	rv := objc.Send[NetService](instance.ID, objc.Sel("initWithDomain:type:name:port:"), objc.String(domain), objc.String(type_), objc.String(name), port)
	rv.Autorelease()
	return rv
}


// Returns an object representing a TXT record formed from a given dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/data(fromTXTRecord:)
func (nc _NetServiceClass) DataFromTXTRecordDictionary(txtDictionary unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("dataFromTXTRecordDictionary:"), txtDictionary)
	return rv
}
// Returns a dictionary representing a TXT record given as an object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/dictionary(fromTXTRecord:)
func (nc _NetServiceClass) DictionaryFromTXTRecordData(txtData unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("dictionaryFromTXTRecordData:"), txtData)
	return rv
}
// Creates a pair of input and output streams for the receiver and returns a Boolean value that indicates whether they were retrieved successfully.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/getInputStream(_:outputStream:)
func (n_ NetService) GetInputStreamOutputStream(inputStream unsafe.Pointer, outputStream unsafe.Pointer) bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("getInputStream:outputStream:"), inputStream, outputStream)
	return rv
}
// Attempts to advertise the receiver’s on the network.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/publish()
func (n_ NetService) Publish() {
	objc.Send[objc.ID](n_.ID, objc.Sel("publish"))
}
// Attempts to advertise the receiver on the network, with the given options.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/publish(options:)
func (n_ NetService) PublishWithOptions(options unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("publishWithOptions:"), options)
}
// Removes the service from the given run loop for a given mode.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/remove(from:forMode:)
func (n_ NetService) RemoveFromRunLoopForMode(aRunLoop unsafe.Pointer, mode unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("removeFromRunLoop:forMode:"), aRunLoop, mode)
}
// Starts a resolve process for the service.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/resolve()
func (n_ NetService) Resolve() {
	objc.Send[objc.ID](n_.ID, objc.Sel("resolve"))
}
// Starts a resolve process of a finite duration for the service.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/resolve(withTimeout:)
func (n_ NetService) ResolveWithTimeout(timeout TimeInterval) {
	objc.Send[objc.ID](n_.ID, objc.Sel("resolveWithTimeout:"), timeout)
}
// Adds the service to the specified run loop.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/schedule(in:forMode:)
func (n_ NetService) ScheduleInRunLoopForMode(aRunLoop unsafe.Pointer, mode unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("scheduleInRunLoop:forMode:"), aRunLoop, mode)
}
// Sets the TXT record for the receiver, and returns a Boolean value that indicates whether the operation was successful.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/setTXTRecord(_:)
func (n_ NetService) SetTXTRecordData(recordData unsafe.Pointer) bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("setTXTRecordData:"), recordData)
	return rv
}
// Starts the monitoring of TXT-record updates for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/startMonitoring()
func (n_ NetService) StartMonitoring() {
	objc.Send[objc.ID](n_.ID, objc.Sel("startMonitoring"))
}
// Halts a currently running attempt to publish or resolve a service.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/stop()
func (n_ NetService) Stop() {
	objc.Send[objc.ID](n_.ID, objc.Sel("stop"))
}
// Stops the monitoring of TXT-record updates for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/stopMonitoring()
func (n_ NetService) StopMonitoring() {
	objc.Send[objc.ID](n_.ID, objc.Sel("stopMonitoring"))
}
// Returns the TXT record for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NetService/txtRecordData()
func (n_ NetService) TXTRecordData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("TXTRecordData"))
	return rv
}

