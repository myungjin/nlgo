// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs syscall_ignore.go

package nlgo

type Ndmsg struct {
	Family  uint8
	Pad1    uint8
	Pad2    uint16
	Ifindex int32
	State   uint16
	Flags   uint8
	Type    uint8
}

type Tcmsg struct {
	Family  uint8
	X_pad1  uint8
	X_pad2  uint16
	Ifindex int32
	Handle  uint32
	Parent  uint32
	Info    uint32
}

const (
	SizeofNdmsg = 0xc
	SizeofTcmsg = 0x14
)
