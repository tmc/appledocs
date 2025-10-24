// Code generated from Apple documentation for Hypervisor. DO NOT EDIT.

package hypervisor

// Type aliases and typedefs
// Hv_allocate_flags_t type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_allocate_flags_t
// hv_allocate_flags_t has base type: uint64_t
type Hv_allocate_flags_t uintptr
// Hv_capability_t - The type of system capabilities.
//
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_capability_t
// hv_capability_t has base type: uint64_t
type Hv_capability_t uintptr
// Hv_gic_config_t - An alias for this value type’s equivalent Hypervisor generic interrupt controller (GIC) configuration’s reference type.
//
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_config_t
// hv_gic_config_t has base type: NSObject<OS_hv_gic_config> *
type Hv_gic_config_t uintptr
// Hv_gic_state_t - An alias for this value type’s equivalent Hypervisor generic interrupt controller (GIC) state’s reference type.
//
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gic_state_t
// hv_gic_state_t has base type: NSObject<OS_hv_gic_state> *
type Hv_gic_state_t uintptr
// Hv_gpaddr_t - The type of a guest physical address (GPA).
//
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_gpaddr_t
// hv_gpaddr_t has base type: uint64_t
type Hv_gpaddr_t uintptr
// Hv_ion_flags_t - The bitfield that you use to set the options flags for the I/O notifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_ion_flags_t
// hv_ion_flags_t has base type: uint32_t
type Hv_ion_flags_t uintptr
// Hv_ipa_t - The type of an intermediate physical address, which is a guest physical address space of the VM.
//
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_ipa_t
// hv_ipa_t has base type: uint64_t
type Hv_ipa_t uintptr
// Hv_memory_flags_t - The permissions for guest physical memory regions.
//
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_memory_flags_t
// hv_memory_flags_t has base type: uint64_t
type Hv_memory_flags_t uintptr
// Hv_return_t - The return type of framework functions.
//
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_return_t
// hv_return_t has base type: mach_error_t
type Hv_return_t uintptr
// Hv_sme_zt0_uchar64_t type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_sme_zt0_uchar64_t
// hv_sme_zt0_uchar64_t has base type: unsigned char __attribute__((ext_vector_type(64)))
type Hv_sme_zt0_uchar64_t uintptr
// Hv_uvaddr_t - The type of a user virtual address.
//
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_uvaddr_t
// hv_uvaddr_t has base type: const void *
type Hv_uvaddr_t uintptr
// Hv_vcpu_options_t - Options for creating a new vCPU instance.
//
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_options_t
// hv_vcpu_options_t has base type: uint64_t
type Hv_vcpu_options_t uintptr
// Hv_vcpu_t - An opaque value that represents a vCPU instance.
//
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_t
// hv_vcpu_t has base type: uint64_t
type Hv_vcpu_t uintptr
// Hv_vcpuid_t - The type that describes a vCPU ID.
//
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vcpuid_t
type Hv_vcpuid_t uint32
// Hv_vm_config_t - The type that defines a virtual-machine configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_config_t
// hv_vm_config_t has base type: NSObject<OS_hv_vm_config> *
type Hv_vm_config_t uintptr
// Hv_vm_options_t - Options you use when creating a virtual machine.
//
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_options_t
// hv_vm_options_t has base type: uint64_t
type Hv_vm_options_t uintptr
// Hv_vm_space_t - The type of a guest-address space.
//
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor/hv_vm_space_t
type Hv_vm_space_t uint32

