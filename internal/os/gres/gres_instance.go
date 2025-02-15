package gres

import "github.com/ilylx/gconv/container/gmap"

const (
	// Default group name for instance usage.
	DEFAULT_NAME = "default"
)

var (
	// Instances map.
	instances = gmap.NewStrAnyMap(true)
)

// Instance returns an instance of Resource.
// The parameter <name> is the name for the instance.
func Instance(name ...string) *Resource {
	key := DEFAULT_NAME
	if len(name) > 0 && name[0] != "" {
		key = name[0]
	}
	return instances.GetOrSetFuncLock(key, func() interface{} {
		return New()
	}).(*Resource)
}
