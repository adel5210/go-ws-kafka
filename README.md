# Go Websocket kafka

### Deployment on docker
- docker build -t go-ws-kafka .
- docker run -p 8090:8090 go-ws-kafka #foreground
- docker run -d -p 8090:8090 go-ws-kafka #background
### Deployment with docker-compose
- docker-compose up -d
### IDE config
- Enable "Go mod integration"