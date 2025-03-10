package v2

import (
	"errors"
	"net/url"
	"path"
)

const (
	// ExtensionsResource is the name of this resource type
	ExtensionsResource = "extensions"
)

// StorePrefix returns the path prefix to this resource in the store
func (e *Extension) StorePrefix() string {
	return ExtensionsResource
}

// URIPath returns the path component of an extension URI.
func (e *Extension) URIPath() string {
	return path.Join(URLPrefix, "namespaces", url.PathEscape(e.Namespace), ExtensionsResource, url.PathEscape(e.Name))
}

// Validate validates the extension.
func (e *Extension) Validate() error {
	if err := ValidateName(e.Name); err != nil {
		return err
	}
	if e.URL == "" {
		return errors.New("empty URL")
	}
	if e.Namespace == "" {
		return errors.New("empty namespace")
	}
	return nil
}

// FixtureExtension given a name returns a valid extension for use in tests
func FixtureExtension(name string) *Extension {
	return &Extension{
		URL:        "https://localhost:8080",
		ObjectMeta: NewObjectMeta(name, "default"),
	}
}

// NewExtension intializes an extension with the given object meta
func NewExtension(meta ObjectMeta) *Extension {
	return &Extension{ObjectMeta: meta}
}

// ExtensionFields returns a set of fields that represent that resource
func ExtensionFields(r Resource) map[string]string {
	resource := r.(*Extension)
	return map[string]string{
		"extension.name":      resource.ObjectMeta.Name,
		"extension.namespace": resource.ObjectMeta.Namespace,
	}
}

// SetNamespace sets the namespace of the resource.
func (e *Extension) SetNamespace(namespace string) {
	e.Namespace = namespace
}
