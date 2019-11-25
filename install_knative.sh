#!/bin/bash 

set -e


install_knative_k8s() {
   
   kubectl apply --selector knative.dev/crd-install=true \
   --filename https://github.com/knative/serving/releases/download/v0.10.0/serving.yaml \
   --filename https://github.com/knative/eventing/releases/download/v0.10.0/release.yaml \
   --filename https://github.com/knative/serving/releases/download/v0.10.0/monitoring.yaml

}

install_knative_eventing_k8s() {
   kubectl apply --filename https://github.com/knative/serving/releases/download/v0.10.0/serving.yaml \
   --filename https://github.com/knative/eventing/releases/download/v0.10.0/release.yaml \
   --filename https://github.com/knative/serving/releases/download/v0.10.0/monitoring.yaml

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
<<<<<<< HEAD
show_knative_status
=======
show_knative_status
>>>>>>> 4bccb644647a2765f7c39e41d25122fc8db71727
