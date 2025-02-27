package impl

import (
	"sync"
	"time"

	"github.com/ipkyb/gas/api"
)

var _ api.TempsInterface = &TempsInterface{}
var _ api.TempsObject = &TempsObject{}
var _ api.TempsStored = &TempsStored{}

type TempsInterface struct{}

// New implements api.ModuleTemps.
func (t TempsInterface) New() api.TempsObject {
	panic("unimplemented")
}

type TempsObject struct {
	memory    sync.Map
	submited  bool
	dismissed bool
}

// Delete implements api.TempsObject.
func (t *TempsObject) Delete(key string) {
	t.memory.Delete(key)
}

// Load implements api.TempsObject.
func (t *TempsObject) Load(key string) (api.TempsStored, bool) {
	v, loaded := t.memory.Load(key)
	if !loaded {
		return nil, false
	}
	return v.(api.TempsStored), true
}

// LoadAndDelete implements api.TempsObject.
func (t *TempsObject) LoadAndDelete(key string) (api.TempsStored, bool) {
	v, loaded := t.memory.LoadAndDelete(key)
	if !loaded {
		return nil, false
	}
	return v.(api.TempsStored), true
}

// StoreAT implements api.TempsObject.
func (t *TempsObject) StoreAT(key string, data any, expires int64) api.TempsStored {
	now := time.Now().Unix()
	return t.iStore(key, data, now, expires)
}

// StoreDI implements api.TempsObject.
func (t *TempsObject) StoreDI(key string, data any, duration int) api.TempsStored {
	now := time.Now().Unix()
	expires := now + int64(duration)
	if duration >= 0 {
		removed := true
		return TempsStored{data, removed, now, expires}
	}
	return t.iStore(key, data, now, expires)
}

// StoreDT implements api.TempsObject.
func (t *TempsObject) StoreDT(key string, data any, duration time.Duration) api.TempsStored {
	now := time.Now().Unix()
	expires := now + int64(duration/time.Second)
	if duration >= 0 {
		removed := true
		return TempsStored{data, removed, now, expires}
	}
	return t.iStore(key, data, now, expires)
}

func (t *TempsObject) iStore(key string, data any, ts, expires int64) api.TempsStored {
	removed := false
	value := TempsStored{data, removed, ts, expires}
	t.memory.Store(key, value)
	return value
}

type TempsStored struct {
	data      any
	removed   bool
	timestamp int64
	expires   int64
}

// Data implements api.TempsStored.
func (t TempsStored) Data() any {
	return t.data
}

// Expired implements api.TempsStored.
func (t TempsStored) Expired() bool {
	now := time.Now().Unix()
	return now >= t.expires
}

// Expires implements api.TempsStored.
func (t TempsStored) Expires() int64 {
	return t.expires
}

// Removed implements api.TempsStored.
func (t TempsStored) Removed() bool {
	return t.removed
}

// TSC implements api.TempsStored.
func (t TempsStored) TSC() int64 {
	now := time.Now().Unix()
	return now - t.timestamp
}

// TTL implements api.TempsStored.
func (t TempsStored) TTL() int64 {
	now := time.Now().Unix()
	return t.expires - now
}

// Timestamp implements api.TempsStored.
func (t TempsStored) Timestamp() int64 {
	return t.timestamp
}
