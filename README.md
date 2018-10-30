# Kubernetes controller for CAT Manifest CRD

Inspired by the article [create kubernetes controllers for core and custom resources](https://medium.com/@trstringer/create-kubernetes-controllers-for-core-and-custom-resources-62fc35ad64a3) 
by Thomas Stringer.

## Steps to take

* create CRD building blocks
    * version
    * api
    * types
    * hooks for generator
* use [Kubernetes Code Generator](https://github.com/kubernetes/code-generator)
    * `./generate-groups.sh all github.com/joostvdg/k8s-cat-resource-controller/pkg/client github.com/joostvdg/k8s-cat-resource-controller/pkg/apis "manifest:v1"`
* create controller
    * main go command (in a `main.go`)
    * handler for the controller events
    * controller to hook into Kubernetes API
* create CRD resource definition
    * see [crd/manifest.yml](crd/manifest.yml)
    * yaml file describing the basics of the resource, the details are in your Go code
    * example resource to demo with
* install CRD resource definition into cluster
    * `kubectl apply -f crd/manifest.yml`
* install controller into cluster
    * create binary
    * docker image with binary in registry
    * create helm chart (deployment, service account, cluster role, clusterrole binding)
    * helm install in cluster
* test with example manifest
* any subsequent update to the CRD will have to rerun the generator again
    * according to the OpenShift article, you should use the scripts in `./hack/` of the code generator
    * `hack/update-codegen.sh`
    * `hack/verify-codegen.sh`

## Resources

* https://engineering.bitnami.com/articles/kubewatch-an-example-of-kubernetes-custom-controller.html
