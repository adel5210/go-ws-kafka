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