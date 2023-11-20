#!/bin/bash

# Starting postgres pod:
kubectl apply -f deployment/database/postgres-secret.yaml 
kubectl apply -f deployment/database/postgres-deployment.yaml 
kubectl apply -f deployment/database/postgres-service.yaml 
echo

# Starting Bookservice pod:
kubectl apply -f deployment/bookservice/bookservice-secret.yaml 
kubectl apply -f deployment/bookservice/bookservice-deployment.yaml 
kubectl apply -f deployment/bookservice/bookservice-service.yaml 
echo

# Starting Authservice pod:
kubectl apply -f deployment/authservice/authservice-secret.yaml 
kubectl apply -f deployment/authservice/authservice-deployment.yaml 
kubectl apply -f deployment/authservice/authservice-service.yaml    
echo

# Checking if the services available or not: 
echo
minikube_ip=$(minikube ip)
desired_code=200
max_retries=50
retry_counter=0
bookservice_Nodeport=$(kubectl get services booksservice -o jsonpath='{.spec.ports[0].nodePort}')
authservice_Nodeport=$(kubectl get services authservice -o jsonpath='{.spec.ports[0].nodePort}')
while true; do
	response_Code_Bookservice=$(curl -s -o /dev/null -w "%{http_code}" http://${minikube_ip}:${bookservice_Nodeport}/queryCount)
	response_Code_Authservice=$(curl -s -o /dev/null -w "%{http_code}" http://${minikube_ip}:${authservice_Nodeport}/queryCount)

	echo "Status Code for bookservice:"
	echo "$response_Code_Bookservice"
	echo
	echo "Status Code for authservice"
	echo "$response_Code_Authservice"
	echo

	if [ "$response_Code_Bookservice" -eq "$desired_code" ] && [ "$response_Code_Authservice" -eq "$desired_code" ]; then
		echo "Services is running. Starting nginx pod..."
        kubectl create configmap nginx-config --from-file=nginx/default.conf 
		kubectl apply -f deployment/nginx/nginx-deployment.yaml 
		kubectl apply -f deployment/nginx/nginx-service.yaml 
		break
	fi

	retry_counter=$((retry_counter + 1))
	echo "Try number: $retry_counter"

    if [ "$retry_counter" -ge "$max_retries" ]; then
        echo "Max retries reached. Exiting script."
        exit 1
    fi

	echo "Waiting for desired code. Retrying in 5 secs...."
	sleep 5
done
