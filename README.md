# golang-quick-start
### Specifications
Give a GitHub user name and repository name as a request parameter, return the number of stars in the repository and the list of followers of the user.

### How to run

#### Pre-tasks
Store the github access token in environment variable GITHUB_TOKEN\
Clone the project from github


#### Run the service
Run the following Makefile script command under the project folder
```
make run
```
or you can run the following go command under the project folder directly
```
go run cmd/hrbrain-go-assignment/main.go
```
Validate the service by the command
```
curl 'http://localhost:8080/repos?user=codecov&repo=example-go'
```

#### Run the test
Run the following Makefile script command under the project folder
```
make test
```
or you can run the following go command under the project folder directly
```
go test -v ./internal/services/...
```

#### Configuration
You are able to change the following configurations in config/config.yaml file
```
#The service will be launched on this port
port : 8080
#The timeout spec for the API call in milliseconds
timeoutMS : 3000
#Retry failed API call until reaching this number
maxAttempt : 3
```
