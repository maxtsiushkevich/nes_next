# NES Next

NES Next is an educational project for building a small 6502 CPU emulator in Go. The goal of the project is to understand how a CPU core works internally: instruction decoding, addressing modes, registers, flags, stack operations, memory access, and interrupt handling.

Although the project is inspired by the architecture used in classic Nintendo systems, it is not yet a full NES emulator. At the moment, it focuses on the CPU and memory subsystem.

## What I am building

This repository contains a lightweight emulator core that aims to:

- emulate the behavior of a 6502-style processor,
- execute opcodes step by step,
- model CPU registers such as accumulator, index registers, program counter, and stack pointer,
- support flag updates for arithmetic and branching logic,
- provide a simple memory abstraction for program and data storage.

## Current capabilities

The current implementation includes:

- a CPU core with register state and flag handling,
- support for many 6502 instructions, including:
  - arithmetic operations such as ADC and SBC,
  - bitwise operations such as AND, OR, XOR,
  - shifts and rotates,
  - comparisons and branches,
  - stack operations like PHA, PLA, PHP, and PLP,
  - jumps and subroutine calls,
  - transfers between registers,
  - flag manipulation instructions,
- several addressing modes, including:
  - immediate,
  - zero page,
  - absolute,
  - indexed and indirect modes,
  - relative branching,
- basic RAM implementation with read/write access,
- interrupt support for NMI, IRQ, BRK, and RESET,
- a set of unit tests covering core CPU behavior.

## Project structure

- [main.go](main.go) — entry point that initializes the CPU and memory and runs a sample program.
- [processor/cpu](processor/cpu) — CPU implementation, opcode table, addressing modes, instruction handlers, and tests.
- [processor/ram](processor/ram) — simple RAM implementation.
- [test_roms](test_roms) — ROM and result data used for testing and experimentation.
- [utils](utils) — helper utilities.

## Getting started

### Prerequisites

- Go 1.24 or newer

### Run the example

From the project root, run:

```bash
go run .
```

### Run tests

```bash
go test ./...
```
