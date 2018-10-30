package main

import (
    "flag"
    "k8s.io/client-go/rest"
    "os"
    "os/signal"
    "syscall"

    log "github.com/Sirupsen/logrus"
    meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
    "k8s.io/client-go/util/workqueue"

    cache "k8s.io/client-go/tools/cache"
    manifestclientset "github.com/joostvdg/k8s-cat-resource-controller/pkg/client/clientset/versioned"
    manifestinformer_v1 "github.com/joostvdg/k8s-cat-resource-controller/pkg/client/informers/externalversions/manifest/v1"
    "github.com/joostvdg/k8s-cat-resource-controller/webserver"
)

// GetClient returns a k8s clientset to the request from inside of cluster
func GetClient() (kubernetes.Interface, manifestclientset.Interface) {
    config, err := rest.InClusterConfig()
    if err != nil {
        log.Fatalf("Can not get kubernetes config: %v", err)
    }

    client, err := kubernetes.NewForConfig(config)
    if err != nil {
        log.Fatalf("Can not create kubernetes client: %v", err)
    }

    manifestClient, err := manifestclientset.NewForConfig(config)
    if err != nil {
        log.Fatalf("getClusterConfig: %v", err)
    }

    log.Info("Successfully constructed k8s client")
    return client, manifestClient
}

func buildOutOfClusterConfig() (*rest.Config, error) {
    kubeconfigPath := os.Getenv("KUBECONFIG")
    if kubeconfigPath == "" {
        kubeconfigPath = os.Getenv("HOME") + "/.kube/config"
    }
    return clientcmd.BuildConfigFromFlags("", kubeconfigPath)
}

// GetClientOutOfCluster returns a k8s clientset to the request from outside of cluster
func GetClientOutOfCluster() (kubernetes.Interface, manifestclientset.Interface) {
    config, err := buildOutOfClusterConfig()
    if err != nil {
        log.Fatalf("Can not get kubernetes config: %v", err)
    }

    client, err := kubernetes.NewForConfig(config)

    manifestClient, err := manifestclientset.NewForConfig(config)
    if err != nil {
        log.Fatalf("getClusterConfig: %v", err)
    }

    log.Info("Successfully constructed k8s client")
    return client, manifestClient
}

// retrieve the Kubernetes cluster client from outside of the cluster
func getKubernetesClient(outOfClusterConfigRequested bool) (kubernetes.Interface, manifestclientset.Interface) {
    if (outOfClusterConfigRequested) {
        return GetClientOutOfCluster()
    }
    return GetClient()
}

// main code path
func main() {
    outOfClusterConfig := flag.Bool("e", false, "boolean flag, for external or 'out of cluster' kubernetes client configuration")
    flag.Parse()

    // get the Kubernetes client for connectivity
    client, manifestClient := getKubernetesClient(*outOfClusterConfig)

    // retrieve our custom resource informer which was generated from
    // the code generator and pass it the custom resource client, specifying
    // we should be looking through all namespaces for listing and watching
    informer := manifestinformer_v1.NewManifestInformer(
        manifestClient,
        meta_v1.NamespaceAll,
        0,
        cache.Indexers{},
    )

    // create a new queue so that when the informer gets a resource that is either
    // a result of listing or watching, we can add an idenfitying key to the queue
    // so that it can be handled in the handler
    queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

    // add event handlers to handle the three types of events for resources:
    //  - adding new resources
    //  - updating existing resources
    //  - deleting resources
    informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
        AddFunc: func(obj interface{}) {
            // convert the resource object into a key (in this case
            // we are just doing it in the format of 'namespace/name')
            key, err := cache.MetaNamespaceKeyFunc(obj)
            log.Infof("Add Manifest: %s", key)
            if err == nil {
                // add the key to the queue for the handler to get
                queue.Add(key)
            }
        },
        UpdateFunc: func(oldObj, newObj interface{}) {
            key, err := cache.MetaNamespaceKeyFunc(newObj)
            log.Infof("Update Manifest: %s", key)
            if err == nil {
                queue.Add(key)
            }
        },
        DeleteFunc: func(obj interface{}) {
            // DeletionHandlingMetaNamsespaceKeyFunc is a helper function that allows
            // us to check the DeletedFinalStateUnknown existence in the event that
            // a resource was deleted but it is still contained in the index
            //
            // this then in turn calls MetaNamespaceKeyFunc
            key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
            log.Infof("Delete Manifest: %s", key)
            if err == nil {
                queue.Add(key)
            }
        },
    })

    // construct the Controller object which has all of the necessary components to
    // handle logging, connections, informing (listing and watching), the queue,
    // and the handler
    controller := Controller{
        logger:    log.NewEntry(log.New()),
        clientset: client,
        informer:  informer,
        queue:     queue,
        handler:   &TestHandler{},
    }

    // use a channel to synchronize the finalization for a graceful shutdown
    stopCh := make(chan struct{})
    defer close(stopCh)

    // run the controller loop to process items
    go controller.Run(stopCh)

    serverPort := "80"
    if len(os.Getenv("SERVER_PORT")) > 0 {
        serverPort = os.Getenv("SERVER_PORT")
    }
    log.Info("Starting WebServer on port ", serverPort)
    go webserver.StartServer(serverPort)

    // use a channel to handle OS signals to terminate and gracefully shut
    // down processing
    sigTerm := make(chan os.Signal, 1)
    signal.Notify(sigTerm, syscall.SIGTERM)
    signal.Notify(sigTerm, syscall.SIGINT)
    <-sigTerm
}
