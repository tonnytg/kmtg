# kmtg
Kubernetes Management


### Helper

Create a user `api-explorer`

Set token for user `api-explorer`

Export Token

    export TOKEN=$(kubectl describe secret $(kubectl get secrets | grep api-explorer | awk '{print $1}') | grep -E '^token' | awk '{print $2}')

Create CA Certificate

    kubectl get secret/api-explorer-secret -o json | jq -Mr '.data["ca.crt"]' | base64 -d > /tmp/ca.crt

Call Go

    APISERVER=https://$(kubectl -n default get endpoints kubernetes --no-headers | awk '{ print $2 }')

    go run main.go --url "$APISERVER/api/v1/namespaces/default/pods" --method "GET"