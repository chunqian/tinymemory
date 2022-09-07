// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package tinymemory

import "unsafe"

var PX_hexToDexTable [103]int32 = [103]int32{
	0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 1, 2, 3, 4, 5, 6,
	7, 8, 9, 0, 0, 0, 0, 0, 0,
	0, 10, 11, 12, 13, 14, 15, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 10,
	11, 12, 13, 14, 15}

// ******************** Typedef ********************
type PX_bool = int32
type PX_dword = uint32
type PX_short = int16
type PX_int16 = int16
type PX_uint16 = uint16
type PX_word = uint16
type PX_ushort = uint16
type PX_uint = uint32
type PX_uint32 = uint32
type PX_int = int32
type PX_int32 = int32
type PX_char = int8
type PX_byte = uint8
type PX_uchar = uint8
type PX_ulong = uint64
type PX_long = int64
type PX_float = float32
type PX_float32 = float32
type PX_double = float64
type PX_double64 = float64
type PX_qword = uint64
type PX_uint64 = uint64
type PX_int64 = int64

type PX_MEMORYPOOL_ERROR = int32
type PX_MP_ErrorCall = func(PX_MEMORYPOOL_ERROR)
type PX_fifobuffer = PX_memory
type PX_stack = PX_memory

// ******************** Enum ********************
const (
	PX_MEMORYPOOL_ERROR_OUTOFMEMORY     PX_MEMORYPOOL_ERROR = 0
	PX_MEMORYPOOL_ERROR_INVALID_ACCESS  PX_MEMORYPOOL_ERROR = 1
	PX_MEMORYPOOL_ERROR_INVALID_ADDRESS PX_MEMORYPOOL_ERROR = 2
)

// ******************** Struct ********************
type PX_str struct {
	data [64]PX_char
}

type PX_memorypool struct {
	allocAddr         unsafe.Pointer
	startAddr         unsafe.Pointer
	endAddr           unsafe.Pointer
	size              PX_uint32
	freeSize          PX_uint32
	nodeCount         PX_uint32
	freeTableCount    PX_uint32
	maxMemoryfragSize PX_uint32
	errorCallPtr      func(PX_MEMORYPOOL_ERROR)
}

type PX_memory struct {
	buffer    *PX_byte
	mp        *PX_memorypool
	usedsize  PX_int
	allocsize PX_int
}

type PX_circularBuffer struct {
	mp      *PX_memorypool
	buffer  *PX_double
	size    PX_int
	pointer PX_int
}
