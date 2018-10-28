package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/joostvdg/cat/application"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MyResource describes a MyResource resource
type CatManifest struct {
	// TypeMeta is the metadata for the resource, like kind and apiversion
	meta_v1.TypeMeta `json:",inline"`
	// ObjectMeta contains the metadata for the particular object, including
	// things like...
	//  - name
	//  - namespace
	//  - self link
	//  - labels
	//  - ... etc ...
	meta_v1.ObjectMeta `json:"metadata,omitempty"`

	// Spec is the custom resource spec
	Spec CatManifestSpec `json:"spec"`
}

// CatManifestSpec is the spec for a MyResource resource
type CatManifestSpec struct {
	// Message and SomeValue are example custom spec fields
	//
	// this is where you would put your custom resource data
    application.Application
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MyResourceList is a list of MyResource resources
type CatManifestList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata"`

	Items []CatManifest `json:"items"`
}
