kubectl exec -n apps -it $(kubectl get pods -n apps | grep "client" | awk 'NR==1{print $1}') sh
