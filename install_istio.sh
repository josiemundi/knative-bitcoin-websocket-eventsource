#!/bin/sh -e

# Make sure we're in the script directory so that Istio is downloaded next to it
cd "$(dirname "$0")"

OS="$(uname)"
if [ "$OS" = "Darwin" ]; then
    OSEXT="osx"
else
    OSEXT="linux"
fi

ISTIO_VERSION="1.4.0"
ISTIO_URL="https://github.com/istio/istio/releases/download/${ISTIO_VERSION}/istio-${ISTIO_VERSION}-${OSEXT}.tar.gz"


download_istio() {
    local name="istio-$ISTIO_VERSION"
    if [ ! -d "$name" ]; then
        echo "Downloading $name from $ISTIO_URL ..." >&2
        curl -L "$ISTIO_URL" | tar xz
    fi

    if [ ! -L "istio" ]; then
        ln -s "$name" "istio"
    fi
}

install_istio_to_k8s() {
    kubectl apply -f istio/install/kubernetes/istio.yaml
}

install_sidecar_injector_to_k8s() {
    # Signed cert as a k8s secret
    ./istio/install/kubernetes/webhook-create-signed-cert.sh \
        --service istio-sidecar-injector \
        --namespace istio-system \
        --secret sidecar-injector-certs

    # Config map
    kubectl apply -f istio/install/kubernetes/istio-sidecar-injector-configmap-release.yaml

    # CA bundle
    cat istio/install/kubernetes/istio-sidecar-injector.yaml | \
     ./istio/install/kubernetes/webhook-patch-ca-bundle.sh > \
     istio/install/kubernetes/istio-sidecar-injector-with-ca-bundle.yaml

    # Install sidecar injector w/ the CA
    kubectl apply -f istio/install/kubernetes/istio-sidecar-injector-with-ca-bundle.yaml

    # Label the default namespace so that it uses the sidecar injection by default
    kubectl label namespace default istio-injection=enabled
}

setup_istio() {
    install_istio_to_k8s 

    # Sidecar injector for attaching Envoy to Kubernetes pods
    install_sidecar_injector_to_k8s
}


download_istio
setup_istio
