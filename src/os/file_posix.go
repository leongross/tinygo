package os

import (
	"time"
)

// Chtimes is a stub, not yet implemented
func Chtimes(name string, atime time.Time, mtime time.Time) error {
	return ErrNotImplemented
}

func (f *File) checkValid(op string) error {
	if f == nil {
		return ErrInvalid
	}
	return nil
}

// setReadDeadline sets the read deadline.
func (f *File) setReadDeadline(t time.Time) error {
	if err := f.checkValid("SetReadDeadline"); err != nil {
		return err
	}
	return ErrNotImplemented
}
