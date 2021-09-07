#! /bin/bash
case $1 in
    setup)
    printf "Starting minikube \n"
    eval minikube start
    eval $(minikube docker-env)

    printf "\nBuilding sender application image:\n"
    docker build -t sender-app ./sender
    printf "\nBuilding receiver application image:\n"
    docker build -t receiver-app ./receiver
    printf "\nCreating and starting containers and service in K8:\n"
    kubectl apply -f receiver-deployment.yaml,receiver-service.yaml,sender-deployment.yaml
    ;;
    logs)
    printf "\nK8 Pods:\n"
    kubectl get pods
    printf "\nK8 Services:\n"
    kubectl get service

    printf "\nSender application logs:\n"
    kubectl logs -l name=sender-app

    printf "\nReceiver application logs:\n"
    kubectl logs -l name=receiver-app
    printf "\n"
    ;;
    clean)
    printf "\nDeleting containers and images\n"
    kubectl delete -f receiver-deployment.yaml,receiver-service.yaml,sender-deployment.yaml
    eval $(minikube docker-env)
    docker image rm sender-app 
    docker image rm receiver-app 
    ;;
    *)
    echo "Choose script functions - "
    echo "k8_helper setup - set up containers with kubernets"
    echo "k8_helper logs - view container output"
    echo "k8_helper clean - stop containers and delete images"
    echo "k8_helper help any other key to exit"
    ;;
esac