# Blockchain Transactions Event Streaming
This project contains the files to build an application that connects to a websocket and streams transaction messages as CloudEvents to a sink. The message stream data is from the blockchain.info WebSocket API and sends information about new bitcoin transactions in real-time. You can find more information about this service at the following [link](https://www.blockchain.com/api/api_websocket).

This application is used as part of another [tutorial](https://github.com/josiemundi/knative-web-event-display) on Knative eventing.

## Building the Docker image

There is already a Docker image available for this tutorial, however if you want to make your own then you could make your own image for your own go application.

To build a Docker image you will need to ensure you have a Dockerfile (there is one in this repo, which I used for building the image we deploy) and then from the directory where it is located, you can run the following commands (ensure you are also logged into your image repo account e.g Dockerhub):

example to build:

```docker build -t josiemundi/wseventsourceimage .```

example to push: 

```docker push josiemundi/wseventsourceimage```



 
