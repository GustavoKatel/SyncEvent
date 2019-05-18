package syncevent

// SyncEvent synchronizes goroutines with a set-reset flag style
type SyncEvent interface {
	// IsSet returns true if set has been called
	IsSet() bool

	// Set sets the flag to true and awake pending goroutines
	Set()

	// Wait waits this flag to be set
	Wait()

	// Reset resets this flag
	Reset()
}
