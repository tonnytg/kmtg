kubectl create serviceaccount api-explorer
kubectl create rolebinding api-explorer:log-reader --clusterrole log-reader --serviceaccount default:api-explorer
kubectl apply -f role.yaml
SERVICE_ACCOUNT=api-explorer
SECRET=$(kubectl create token api-explorer )
TOKEN=$(kubectl get secret $SECRET -o jsonpath='{.data.token}' | base64 -d)
kubectl config set-credentials $SERVICE_ACCOUNT --token=$TOKEN
kubectl get secret/api-explorer-secret -o json | jq -Mr '.data["ca.crt"]' | base64 -d > /tmp/ca.crt

// path to consum pods information from API
URL = "https://kubernetes.default.svc/api/v1/namespaces/default/pods"