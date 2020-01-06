#!/bin/bash 

set -e


install_knative_k8s() {
   
   kubectl apply --selector knative.dev/crd-install=true \
   --filename https://github.com/knative/serving/releases/download/v0.11.1/serving.yaml \
   --filename https://github.com/knative/eventing/releases/download/v0.11.0/release.yaml \
   --filename https://github.com/knative/serving/releases/download/v0.11.1/monitoring.yaml

}

install_knative_eventing_k8s() {
   kubectl apply --filename https://github.com/knative/serving/releases/download/v0.11.1/serving.yaml \
   --filename https://github.com/knative/eventing/releases/download/v0.11.0/release.yaml \
   --filename https://github.com/knative/serving/releases/download/v0.11.1/monitoring.yaml

}


setup_knative() {
    # Install Knative
    install_knative_k8s 

    # Install eventing
    install_knative_eventing_k8s
}

show_knative_status(){
   kubectl get pods --namespace knative-serving
   kubectl get pods --namespace knative-eventing
   kubectl get pods --namespace knative-monitoring
}

setup_knative
show_knative_status

