#!/bin/bash

function CreateAccount {
    echo "Creating account $1"
    kubectl create serviceaccount api-explorer
    export TOKEN=`kubectl create token api-explorer`
    echo $TOKEN
    kubectl create secret generic api-explorer-secret --from-literal=token=$TOKEN
    kubectl apply -f ./config/role.yaml
}

if [ $1 = "init" ]; then
    CreateAccount
fi

export TOKEN=$(kubectl describe secret $(kubectl get secrets | grep api-explorer | awk '{print $1}') | grep -E '^token' | awk '{print $2}')

kubectl get secret/api-explorer-secret -o json | jq -Mr '.data["ca.crt"]' | base64 -d > /tmp/ca.crt

export APISERVER=https://$(kubectl -n default get endpoints kubernetes --no-headers | awk '{ print $2 }')

go run main.go --url "$APISERVER/api/v1/namespaces/default/pods" --method "GET"