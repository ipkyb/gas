package api

import "time"

// Temps is a module for temporary (auto-expiry) database.
// The approach is generally known as key-value store.
var Temps struct {
	TempsInterface
}

// The api interface for Temps module.
type TempsInterface interface {
	// Creates a new Temps object.
	New() TempsObject
}

// The object of temp store, container of the managed database.
type TempsObject interface {
	// Delete object by key from Temps.
	Delete(key string)
	// Load object by key from Temps.
	Load(key string) (TempsStored, bool)
	// Load and delete the object by key from Temps.
	LoadAndDelete(key string) (TempsStored, bool)
	// Store object to Temps using key for retrieval later. The value expires
	// is the specific unix timestamp of when to removed the data.
	StoreAT(key string, data any, expires int64) TempsStored
	// Store object to Temps using key for retrieval later. The value duration
	// is the duration (since now) of when to removed the data (in seconds).
	StoreDI(key string, data any, duration int) TempsStored
	// Store object to Temps using key for retrieval later. The value duration
	// is the duration (since now) of when to removed the data.
	StoreDT(key string, data any, duration time.Duration) TempsStored
}

// The information of stored object inside Temps.
type TempsStored interface {
	// Retrieve your original data.
	Data() any
	// Does this object has been successfully removed from Temps.
	Removed() bool
	// Does this object should have been expired.
	Expired() bool
	// Unix timestamp (in seconds) of planned configured time to be removed
	// from Temps.
	Expires() int64
	// Unix timestamp (in seconds) of when you store this object in Temps.
	Timestamp() int64
	// Time Since Creation (TSC) indicates the lifetime of this object in the
	// Temps (in seconds). Expression: Now - Timestamp.
	TSC() int64
	// Time To Live (TTL) indicates the remaining time to live for this object
	// in the Temps (in seconds). Expression: Expires - Now.
	TTL() int64
}
