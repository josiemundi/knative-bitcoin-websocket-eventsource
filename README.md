# knative-eventing-websocket-source
This project explores Knative Eventing. The aim of the project is to deploy a go application that streams messages to a Knative broker, which then delivers the messages to an event display service. The message stream data is from the blockchain.info WebSocket API and sends information about new bitcoin transactions in real-time. You can find more information about this service at the link below:

https://www.blockchain.com/api/api_websocket


In order to run this demo, you will need:

- A Kubernetes cluster (a single node cluster running on Docker desktop is fine and is used to build this example)
- kubectl
- Istio installed (instructions below)
- Knative installed (inctructions below)

## Installing Istio

To install Istio, I am currently following the instructions below:

https://knative.dev/docs/install/installing-istio/

Then run the following command in order to enable Istio in the default (or alternative) namespace:

```kubectl label namespace <namespace> istio-injection=enabled```

I had a lot of issues getting this demo up and running. Finally I realised that I needed to add the cluster local gateway to my Istio installation. This is not installed as standard in the Knative instructions but is mentioned further down the page. 


## Installing Knative 

Run the install_knative.sh script to install the Knative components. This will do a complete install of Knative Serving, Eventing and Monitoring. For a lighter install, you can follow the below instructions for Docker Desktop:

https://knative.dev/docs/install/knative-with-docker-for-mac/


To confirm your install is complete, you can run the following command:

```kubectl get pods --all-namespaces```

You should have namespaces for ```istio-system```, ```knative-eventing```, ```knative-serving``` (and ```knative-monitoring``` if you have installed using the install script).

## Run the yaml scripts

There are 4 main yaml scripts that need to be run to get this tutorial working. 

First ```kubectl apply -f 001-namespace.yaml``` This will deploy the knative-eventing-websocket-source namespace and enable the knative-eventing injection. 

The run the following:

```kubectl apply -f 010-deployment.yaml```

```kubectl apply - 040-trigger.yaml```

```kubectl apply - 030-service.yaml```

