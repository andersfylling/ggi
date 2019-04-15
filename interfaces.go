package ggi

// Reseter Resets the struct to their default values. Except fields that implements the Reseter interface.
//
// Fields/values that implements the Reseter interface will simply invoke .Reset(), instead of being set to
// nil/0/etc. to reduce GC.
//
// TODO: allow tweaking this behavior with flags.
type Reseter interface {
	Reset()
}