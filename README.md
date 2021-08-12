# hrbrain-go-assignment
### Completed Specifications
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
Clone the project from github\
Store the github access token in environment variable GITHUB_TOKEN

#### Run the unit test
Run following command under project folder
```
make test
```

#### Run the service
Run following command under project folder
```
make run
```
Test the service by the command
```
curl 'http://localhost:8080/repos?user=codecov&repo=example-go'
```