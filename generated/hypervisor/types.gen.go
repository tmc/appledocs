// Code generated from Apple documentation for Hypervisor. DO NOT EDIT.

package hypervisor
import (
	"unsafe"
)


// C struct types
// hv_apic_state
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_apic_state
type hv_apic_state struct {
	Aeoi uint32
	Apic_controls uint64
	Apic_gpa uint64
	Apic_id uint32
	Apr uint32
	Boot_state hv_boot_state
	Ccr_timer uint32
	Dcr_timer uint32
	Dfr uint32
	Esr uint32
	Esr_pending uint32
	Icr uint32
	Icr_timer uint32
	Irr uint32
	Isr uint32
	Ldr uint32
	Lvt uint32
	Svr uint32
	Tmr uint32
	Tpr uint32
	Tsc_deadline uint64
	Ver uint32
}/* debug [types.gen.go/struct]: hv_apic_state */

// hv_apic_state_ext_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_apic_state_ext_t
type hv_apic_state_ext_t struct {
	State Hv_apic_state
	Version uint32
}/* debug [types.gen.go/struct]: hv_apic_state_ext_t */

// hv_atpic_state
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_atpic_state
type hv_atpic_state struct {
	Aeoi bool
	Elc uint8
	Icw_num uint8
	Intr_raised bool
	Irq_base uint8
	Last_request uint8
	Lowprio uint8
	Mask uint8
	Poll bool
	Rd_cmd_reg uint8
	Ready bool
	Request uint8
	Rotate bool
	Service uint8
	Sfn bool
	Smm bool
}/* debug [types.gen.go/struct]: hv_atpic_state */

// hv_atpic_state_ext_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_atpic_state_ext_t
type hv_atpic_state_ext_t struct {
	State Hv_atpic_state
	Version uint32
}/* debug [types.gen.go/struct]: hv_atpic_state_ext_t */

// hv_ioapic_state
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_ioapic_state
type hv_ioapic_state struct {
	Ioa_id uint32
	Ioregsel uint32
	Irr uint32
	Rtbl uint64
}/* debug [types.gen.go/struct]: hv_ioapic_state */

// hv_ioapic_state_ext_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_ioapic_state_ext_t
type hv_ioapic_state_ext_t struct {
	State Hv_ioapic_state
	Version uint32
}/* debug [types.gen.go/struct]: hv_ioapic_state_ext_t */

// hv_ion_message_t - The structure that describes the Mach message that the Hypervisor sends when an I/O notifier delivers the notifications you request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_ion_message_t
type hv_ion_message_t struct {
	Addr uint64 // The address of the I/O write.
	Header unsafe.Pointer // The Mach message header.
	Size uint64 // The size of the value written by the notifier.
	Trailer unsafe.Pointer // The Mach message trailer.
	Value uint64 // An unsigned 64-bit integer that represents the contents of an I/O notifier message.
}/* debug [types.gen.go/struct]: hv_ion_message_t */

// hv_vcpu_exit_exception_t - The structure that describes information about an exit from the virtual CPU (vCPU) to the host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_exit_exception_t
type hv_vcpu_exit_exception_t struct {
	Physical_address Hv_ipa_t // The intermediate physical address of the exception in the client.
	Syndrome Hv_exception_syndrome_t // The vCPU exception syndrome causing the exception.
	Virtual_address Hv_exception_address_t // The vCPU virtual address of the exception.
}/* debug [types.gen.go/struct]: hv_vcpu_exit_exception_t */

// hv_vcpu_exit_t - Information about an exit from the vCPU to the host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_exit_t
type hv_vcpu_exit_t struct {
	Exception Hv_vcpu_exit_exception_t // Information about an exit exception from the vcpu to the host.
	Reason unsafe.Pointer // Information about an exit from the vcpu to the host.
}/* debug [types.gen.go/struct]: hv_vcpu_exit_t */

// hv_vcpu_sme_state_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_sme_state_t
type hv_vcpu_sme_state_t struct {
	Streaming_sve_mode_enabled bool
	Za_storage_enabled bool
}/* debug [types.gen.go/struct]: hv_vcpu_sme_state_t */





