# knative-eventing-websocket-source
This project explores Knative Eventing. 

In order to run this demo, you will need:

- A Kubernetes cluster (a single node cluster running on Docker desktop is fine)
- kubectl
- Istio installed (instructions below)

## Installing Istio

To install Istio, I am currently following the instructions below:

https://istio.io/docs/setup/getting-started/#download

Then run the following command in order to enable Istio in the default (or alternative) namespace:

```kubectl label namespace <namespace> istio-injection=enabled```


## Installing Knative 

Run the install_knative.sh script to install the Knative components. 
