package main

import (
	log "github.com/sirupsen/logrus"
	manifest_v1 "github.com/joostvdg/k8s-cat-resource-controller/pkg/apis/manifest/v1"
	client "github.com/joostvdg/cat-grpc-client/cmd"
    "github.com/joostvdg/cat/pkg/api/v1"
)

// Handler interface contains the methods that are required
type Handler interface {
	Init() error
	ObjectCreated(obj interface{})
	ObjectDeleted(obj interface{})
	ObjectUpdated(objOld, objNew interface{})
}

// TestHandler is a sample implementation of Handler
type TestHandler struct{}

// Init handles any handler initialization
func (t *TestHandler) Init() error {
	log.Info("TestHandler.Init")
	return nil
}

// ObjectCreated is called when an object is created
func (t *TestHandler) ObjectCreated(obj interface{}) {
	log.Info("TestHandler.ObjectCreated")
	// assert the type to a Pod object to pull out relevant data
	manifest := obj.(*manifest_v1.Manifest)
	log.Infof("    ResourceVersion: %s", manifest.ObjectMeta.ResourceVersion)
	log.Infof("    Name: %s", manifest.Spec.Name)
    application := v1.Application {
        Name:        manifest.Spec.Name,
        Description: manifest.Spec.Description,
        Namespace:   manifest.Spec.Namespace,
        Sources:     manifest.Spec.Sources,
        ArtifactIDs: manifest.Spec.ArtifactIDs,
    }
	client.CreateApplication("cat:9090", application)
    log.WithFields(log.Fields{
        "Name": manifest.Spec.Name,
    }).Info("Create new application")
}

// ObjectDeleted is called when an object is deleted
func (t *TestHandler) ObjectDeleted(obj interface{}) {
	log.Info("TestHandler.ObjectDeleted")
}

// ObjectUpdated is called when an object is updated
func (t *TestHandler) ObjectUpdated(objOld, objNew interface{}) {
	log.Info("TestHandler.ObjectUpdated")
}
