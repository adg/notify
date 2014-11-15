// +build darwin dragonfly freebsd netbsd openbsd
// +build !fsnotify

package notify

import "syscall"

const (
	Create Event = 0x0100 << iota
	Delete
	Write
	Move
	Recursive // An internal event, in final version won't be exported.
)

const (
	// NOTE_DELETE is an even reported when the unlink() system call was called
	// on the file referenced by the descriptor.
	NOTE_DELETE = Event(syscall.NOTE_DELETE)
	// NOTE_WRITE is an event reported when a write occurred on the file
	// referenced by the descriptor.
	NOTE_WRITE = Event(syscall.NOTE_WRITE)
	// NOTE_EXTEND is an event reported when the file referenced by the
	// descriptor was extended.
	NOTE_EXTEND = Event(syscall.NOTE_EXTEND)
	// NOTE_ATTRIB is an event reported when the file referenced
	// by the descriptor had its attributes changed.
	NOTE_ATTRIB = Event(syscall.NOTE_ATTRIB)
	// NOTE_LINK is an event reported when the link count on the file changed.
	NOTE_LINK = Event(syscall.NOTE_LINK)
	// NOTE_RENAME is an event reported when the file referenced
	// by the descriptor was renamed.
	NOTE_RENAME = Event(syscall.NOTE_RENAME)
	// NOTE_REVOKE is an event reported when access to the file was revoked via
	// revoke(2) or	the underlying file system was unmounted.
	NOTE_REVOKE = Event(syscall.NOTE_REVOKE)
)

var osestr = map[Event]string{
	NOTE_DELETE: "notify.NOTE_DELETE",
	NOTE_WRITE:  "notify.NOTE_WRITE",
	NOTE_EXTEND: "notify.NOTE_EXTEND",
	NOTE_ATTRIB: "notify.NOTE_ATTRIB",
	NOTE_LINK:   "notify.NOTE_LINK",
	NOTE_RENAME: "notify.NOTE_RENAME",
	NOTE_REVOKE: "notify.NOTE_REVOKE",
}

var ekind = map[Event]Event{
	syscall.NOTE_WRITE:  Write,
	syscall.NOTE_RENAME: Move,
	syscall.NOTE_DELETE: Delete,
}
