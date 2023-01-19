# gocx

Dialogflow CX agent handy commands.

## Commands

```sh
# Example shell commands
gocx help
gocx get intents $agent-id -dir ${dir} -file ${filename} -verbose
gocx get flows $agent-id  -dir ${dir} -file ${filename} ${filename} -verbose
gocx get pages $agent-id -dir ${dir} -file ${filename} ${filename} -verbose 
gocx server -port=9000
```

## Docker setup

Run docker:

```sh
# Run the docker on alpine
docker build -t gocx:v1 .
#To save on docker
docker login -u username
docker tag gocx:v1 uopendocker/gocx:v1
docker push uopendocker/gocx:v1

# Run anywhere 
docker run uopendocker/gocx:v1 /gocx get intents --agentID="hello"
```

## Testing

```shell
# Run specific tests
go test -v -run TestExecute_RootCmd ./...
```
