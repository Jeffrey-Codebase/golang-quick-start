# hrbrain-go-assignment
### Completed specifications
Give a GitHub user name and repository name as a request parameter\
Return the number of stars in the repository and the list of followers of the user as a
response\
Use go-github for the implementation\
Implement the test code\
Implement a command line tool to start the server\
Implement the code in clean and organized package structure\
Throttle the API call if it reaches the rate limit\
Retry the API call properly if 500 error is responded\
Set timeout on API call\
Implement the code which works as fast as possible in multi-core environment\
Allocate memory as little as possible\
Implement the reproducible\
Write simple and readable code

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
go test ./internal/services/...
```

#### Configuration
You are able to change the following configurations in config/config.yaml file
```
#The service will be launched on this port
port : 8080
#The timeout spec for the API call in milliseconds
timeoutMS : 1000
#Retry failed API call until reaching this number
maxAttempt : 3
```