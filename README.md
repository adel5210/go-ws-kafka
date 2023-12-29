# Go Websocket kafka

### Deployment
- docker build -t go-ws-kafka .
- docker run -p 8090:8090 go-ws-kafka #foreground
- docker run -d -p 8090:8090 go-ws-kafka #background
### IDE config
- Enable "Go mod integration"