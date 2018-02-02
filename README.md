# YoGo - golang code generator 
![Logo](https://github.com/andriylesch/generator-yogo/blob/master/logo_small.png)

YoGo is a [Yeoman](http://yeoman.io/) generator for [GO](https://golang.org/) language, which allows you to create projects in the form of:
- Console (basic "Hello world") apps
- REST API microservice
- GO-KIT API microservice

This tool is meant for developers with the aim of removing the need to recreate each time the same skeleton when starting a new project. 
YoGo generates it for him, so he can focus on implementing logic in new app.

### The Go Programming Language

Detail info you can find [here](https://golang.org/doc/).

# Preparation for Usage

- Install yeoman on your system

```bash
npm install -g yo
```
- Install generator-yogo

```bash
npm install generator-yogo
```

- Launch the generator
```bash
yo yogo
```

# Create new project

Create a new directory in **GOPATH/src**

```bash
cd $GOPATH/src
mkdir [ProjectName]
cd [ProjectName]
yo yogo
```

...and follow the interactive option selection

## Structure of projects

```
  NOTE: the YoGo generator will only generate the directory hierarchy in *$GOPATH/src*.
```
**Example console app**
<pre>
  $GOPATH folder
  └── src
      └── <YOUR_APP>
          ├── main.go                   # entrypoint
          ├── .gitignore                 
          ├── Dockerfile                 
          ├── Makefile                  
          └── Readme.md                  
</pre>
  
**Example REST API microservice**
<pre>
$GOPATH folder
└── src
    └── <YOUR_APP>
        ├── main.go                     # entrypoint
        ├── config
        │   ├── config.go               # source file
        │   ├── config_test.go          # test file
        │   ├── config.toml             # config file
        │   └── config.toml.example     # config example file
        ├── <NAME_YOUR_PACKAGE>
        │   ├── endpoint.go             
        │   ├── interface.go            
        │   ├── model.go                
        │   ├── repository.go           
        │   ├── service.go              
        │   └── service_tracing.go      
        ├── .gitignore                   
        ├── Dockerfile                   
        ├── Makefile                    
        ├── Gopkg.toml          
        └── Readme.md                   
</pre>

**FYI:** Future development will add support for:
- KAFKA producer
- KAFKA consumer
- etc.

# Generate project via docker 

If you want to generate your project in a Docker container (i.e. you don't want to install the YoGo dependencies on your machine):
- install docker on your machine 
- clone this repository (git clone)
- open a terminal and go to the `generator-yogo` folder
- enter the following command:

for windows users
``` bash
docker build -t generator-yogo:latest .
docker run -it -v {YOUR_LOCAL_PATH}:/home/yeoman -e LOCAL_PATH=${LOCAL_PATH} --name generator-yogo-container generator-yogo
```

for Linux and MacOS users

```bash
make docker LOCAL_PATH=<YOUR_LOCAL_PATH>
```
where **<YOUR_LOCAL_PATH>** is the local folder where you want your project to be generated.


# How to create a new package in an existing microservice.

Open a terminal, and from your project's folder run:

```bash
yo yogo:pkgg
```

You can choose the kind of package you want to create and use:
- config
- REST-API endpoint
- GO-KIT endpoint

### Config package

One possible type of package is **config**. 
In this case, the package will be generated with the following structure 

<pre>
$GOPATH folder
└── src
    └── <YOUR_APP>
          ├── config
          ├── config.go               # source file
          ├── config_test.go          # test file
          ├── config.toml             # config file
          └── config.toml.example     # config example file
</pre>

###### NOTE: By default, **YOGO** generates a config.toml file with three ENV variables ( *APP_PORT*, *KAFKA_BROKERS*, *SVC_TRACING_ZIPKIN* )

#### Extend config package 

You can add a config.toml file in your project's root folder with the list of ENV variables you need.

###### For example:
```
TEST=12345
TEST1=4567
TEST2=qwerty
DB_PORT=5432
```

After doing this, when generating the config package, all variables will be added on top of the default ones (*APP_PORT*, *KAFKA_BROKERS*, *SVC_TRACING_ZIPKIN*).
