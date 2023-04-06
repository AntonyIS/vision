<!-- <img src="./assets/imgs/images.jpeg" alt="Logo of the project" align="right"> -->

# Vision Backend


## Installing / Getting started
1. Ensure you have Git installed on your computer and clone this project using the below command.
```
git clone https://github.com/AntonyIS/vision.git
```
The project will be cloned into you machine

## Technology stack
*   Golang - Application runtime
*   Go Gin - Go web application framework
*   JWT - For authentication and resource authorization
*   AWS DynamoDB - Store user and projects
*   Docker - Application packaging | Containerization
*   AWS Elasticbeanstalk for deployment to cloud

# Application Architecture
This application is made up of the Hexagonal(Ports and Adapter) architecture

```
├── config
│   └── config.go
├── go.mod
├── go.sum
├── internal
│   ├── adapters
│   │   ├── database
│   │   │   ├── dynamodb
│   │   │   │   └── dynamodb.go
│   │   │   ├── s3
│   │   │   │   └── s3.go
│   │   │   └── timestreamdb
│   │   │       └── timestream.go
│   │   └── http
│   │       ├── gin.go
│   │       ├── handlers.go
│   │       └── middleware.go
│   └── core
│       ├── domain
│       │   └── domain.go
│       ├── ports
│       │   └── ports.go
│       └── services
│           └── services.go
├── LICENSE
├── main.go
├── Makefile
└── README.md
```
<!-- ## Developing 

<!-- ### Built With
List main libraries, frameworks used including versions (React, Angular etc...)

### Prerequisites
What is needed to set up the dev environment. For instance, global dependencies or any other tools. include download links.

-->
### Setting up Dev
* Clone the application
```
git clone https://github.com/AntonyIS/vision
```
* Navigate into the working directory
```
cd vision
```
* Tidy up the application: install dependancies
```
go mod tidy
```
* Run the application
```
make serve
```
* Run the tests
```
make test
```
<!-- 
```shell
git clone https://github.com/your/your-project.git
cd your-project/
packagemanager install
```

And state what happens step-by-step. If there is any virtual environment, local server or database feeder needed, explain here.

### Building

If your project needs some additional steps for the developer to build the
project after some code changes, state them here. for example:

```shell
./configure
make
make install
```

Here again you should state what actually happens when the code above gets
executed.

### Deploying / Publishing
give instructions on how to build and release a new version
In case there's some step you have to take that publishes this project to a
server, this is the right time to state it.

```shell
packagemanager deploy your-project -s server.com -u username -p password
```

And again you'd need to tell what the previous code actually does.

## Versioning

We can maybe use [SemVer](http://semver.org/) for versioning. For the versions available, see the [link to tags on this repository](/tags).


## Configuration

Here you should write what are all of the configurations a user can enter when using the project.

## Tests

Describe and show how to run the tests with code examples.
Explain what these tests test and why.

```shell
Give an example
```

## Style guide

Explain your code style and show how to check it.

## Api Reference

If the api is external, link to api documentation. If not describe your api including authentication methods as well as explaining all the endpoints with their required parameters.


## Database

Explaining what database (and version) has been used. Provide download links.
Documents your database design and schemas, relations etc... 

## Licensing

State what the license is and how to find the text version of the license. --> 