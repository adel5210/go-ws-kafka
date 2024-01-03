# Go Websocket kafka

### Deployment on docker
- docker build -t go-ws-kafka .
- docker run -p 8090:8090 go-ws-kafka #Foreground
- docker run -d -p 8090:8090 go-ws-kafka #Background
### Deployment with docker-compose
- docker-compose up -d #Start
- docker-compose down #Stop and remove
### IDE config
- Enable "Go mod integration"
### Push to docker hub
- docker images
- docker tag go-ws-kafka adel5210/go-websocket-kafka-backend:1.0.2
- docker push adel5210/go-websocket-kafka-backend:1.0.2
### Kube setup attempt 1
- kubectl create -f kube-go-ws-kafka-pod.yml
- kubectl get pods
- kubectl get svc
### Kube setup attempt 2
- kubectl run go-ws-kafka --image=adel5210/go-websocket-kafka-backend:1.0.1 --port=8090
- kubectl port-forward go-ws-kafka 8090:8090
### Kube setup attempt 3
- kubectl create deployment go-ws-kafka --image=adel5210/go-websocket-kafka-backend:1.0.1
- kubectl expose deployment go-ws-kafka --type=NodePort --port=8090
- kubectl get svc
- minikube service go-ws-kafka --url
### Kube setup attempt 4
- kubectl apply -f k8s/deployment.yaml 
### Kube deploy as service
- kubectl apply -f k8s/service.yaml