// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents_test

import (
	"github.com/tmc/appledocs/generated/intents"
)

// Suppress unused import errors
var _ = intents.NewINStartCallIntent

// ExampleNewINStartCallIntentWithAudioRouteDestinationTypeContactsRecordTypeForRedialingCallCapability demonstrates how to create a INStartCallIntent instance using NewINStartCallIntentWithAudioRouteDestinationTypeContactsRecordTypeForRedialingCallCapability.
// Creates a start call intent object with the specified parameters.
func ExampleNewINStartCallIntentWithAudioRouteDestinationTypeContactsRecordTypeForRedialingCallCapability() {
	_ = intents.NewINStartCallIntentWithAudioRouteDestinationTypeContactsRecordTypeForRedialingCallCapability(
		intents.INCallAudioRoute{},      // audioRoute INCallAudioRoute
		intents.INCallDestinationType{}, // destinationType INCallDestinationType
		[]intents.INPerson{},            // contacts []INPerson
		intents.INCallRecordType{},      // recordTypeForRedialing INCallRecordType
		intents.INCallCapability{},      // callCapability INCallCapability
	)
	// Output:
}

// ExampleNewINStartCallIntentWithCallRecordFilterCallRecordToCallBackAudioRouteDestinationTypeContactsCallCapability demonstrates how to create a INStartCallIntent instance using NewINStartCallIntentWithCallRecordFilterCallRecordToCallBackAudioRouteDestinationTypeContactsCallCapability.
// Creates a start call intent object with the specified parameters.
func ExampleNewINStartCallIntentWithCallRecordFilterCallRecordToCallBackAudioRouteDestinationTypeContactsCallCapability() {
	_ = intents.NewINStartCallIntentWithCallRecordFilterCallRecordToCallBackAudioRouteDestinationTypeContactsCallCapability(
		intents.INCallRecordFilter{},    // callRecordFilter INCallRecordFilter
		intents.INCallRecord{},          // callRecordToCallBack INCallRecord
		intents.INCallAudioRoute{},      // audioRoute INCallAudioRoute
		intents.INCallDestinationType{}, // destinationType INCallDestinationType
		[]intents.INPerson{},            // contacts []INPerson
		intents.INCallCapability{},      // callCapability INCallCapability
	)
	// Output:
}
