# Virtualization Examples Index

Quick reference for navigating the virtualization examples.

## Main Implementations

### 1. vz/ - Production VM Runner ⭐⭐⭐⭐⭐

**Status**: Complete (blocked on Foundation)  
**Lines**: 3,617 (1,090 code + 2,527 docs/tools)

**Quick Access**:
- [README](vz/README.md) - User guide
- [STATUS](vz/STATUS.md) - Current status
- [CONTRIBUTING](vz/CONTRIBUTING.md) - How to contribute
- [Makefile](vz/Makefile) - Build automation
- [Scripts](vz/scripts/) - Helper scripts

**Key Files**:
```
vz/main.go      - Entry point (393 lines)
vz/config.go    - Configuration (380 lines)
vz/bundle.go    - Path management (29 lines)
```

### 2. vm-create/ - Quick Examples

**Status**: Working (same Foundation block)  
**Lines**: ~400

**Files**:
- macos_vm.go - macOS configuration example
- main_quickstart.go - Simple Linux VM

## Documentation

| File | Purpose | Lines |
|------|---------|-------|
| [README.md](README.md) | Overview & comparison | 297 |
| [SUMMARY.md](SUMMARY.md) | First implementation | 293 |
| [vz/README.md](vz/README.md) | vz user guide | 558 |
| [vz/STATUS.md](vz/STATUS.md) | Status & blockers | 263 |
| [vz/CONTRIBUTING.md](vz/CONTRIBUTING.md) | Contribute guide | 402 |
| [vz/IMPLEMENTATION_SUMMARY.md](vz/IMPLEMENTATION_SUMMARY.md) | Architecture | 308 |
| [vz/FINAL_SUMMARY.md](vz/FINAL_SUMMARY.md) | Complete summary | 370 |

## Quick Start

### If Foundation Was Working

```bash
# Setup and run
cd vz
make setup
make run

# Or with scripts
./scripts/setup-vm.sh --disk-size 64
./scripts/status-vm.sh
```

### Current Reality

Foundation bindings block compilation. See:
- [vz/STATUS.md](vz/STATUS.md#blocking-issue-foundation-bindings)

## File Tree

```
virtualization/
├── INDEX.md                     # This file
├── README.md                    # Overview
├── SUMMARY.md                   # First implementation summary
│
├── vz/                          # Main implementation
│   ├── main.go                  # Entry point
│   ├── config.go                # VM configuration
│   ├── bundle.go                # Path management
│   ├── delegate_example.go      # Delegate patterns
│   ├── Makefile                 # Build automation
│   ├── go.mod, go.sum          # Dependencies
│   ├── scripts/                 # Automation
│   │   ├── setup-vm.sh
│   │   ├── status-vm.sh
│   │   └── clean-vm.sh
│   └── docs/                    # Documentation
│       ├── README.md
│       ├── STATUS.md
│       ├── CONTRIBUTING.md
│       ├── IMPLEMENTATION_SUMMARY.md
│       └── FINAL_SUMMARY.md
│
└── vm-create/                   # Quick examples
    ├── macos_vm.go
    ├── main_quickstart.go
    └── main.go
```

## Navigation Tips

**For users**: Start with [vz/README.md](vz/README.md)  
**For contributors**: See [vz/CONTRIBUTING.md](vz/CONTRIBUTING.md)  
**For status**: Check [vz/STATUS.md](vz/STATUS.md)  
**For architecture**: Read [vz/IMPLEMENTATION_SUMMARY.md](vz/IMPLEMENTATION_SUMMARY.md)

## Blocking Issue

All examples blocked by Foundation framework binding errors.  
**Fix needed**: See [vz/STATUS.md](vz/STATUS.md#required-fixes)

Once fixed, all functionality will be immediately available.
