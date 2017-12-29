# YoGo - golang code generator 
![Logo](https://github.com/andriylesch/generator-yogo/blob/master/logo_small.png)

[Yeoman](http://yeoman.io/) generator for [GO](https://golang.org/) language allows you to create projects like :
- Console (basic "Hello world")
- Rest API microservice

So right now as developer you don't need to create each time the same skeleton of project. 
Just generate it and focus for implementing logic in new microservice.

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

Create new directory in folder **GOPATH/src** folder

```bash
cd $GOPATH/src
mkdir [ProjectName]
cd [ProjectName]
yo yogo
```

And follow list of options

## Structure of projects

```
  ATTENTION:  YoGo generator will only generate directory hierarchy in *$GOPATH/src* folder.
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

**FYI:** In future will be provided list of new projects
- GO-KIT microservice
- KAFKA producer
- KAFKA consumer
- etc

# Generate project via docker 

Follow options bellow
- install docker on your local system 
- clone this repository (git clone)
- open terminal and go to `generator-yogo` folder
- type command bellow

for windows users
``` bash
docker build -t generator-yogo:latest .
docker run -it -v {YOUR_LOCAL_PATH}:/home/yeoman -e LOCAL_PATH=${LOCAL_PATH} --name generator-yogo-container generator-yogo
```

for Linux and MacOS users

```bash
make docker LOCAL_PATH=YOUR_LOCAL_PATH
```
where **YOUR_LOCAL_PATH** - local folder where will be generated your project


# How to create new package in existed microservice.

Open in terminal your project.

```bash
yo yogo:pkgg
```

List of packages what is possible to create and use :
- config
- REST-API endpoint

### Config package

Developer has possibility to add **config** package in his application. 
In this case it will be generated standart config package with structure 

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

###### NOTE: As default, **YOGO** generate config.toml file with three ENV valiables ( *APP_PORT*, *KAFKA_BROKERS*, *SVC_TRACING_ZIPKIN* )

#### Extend config package 

Add config.toml file in your root folder of project with list of ENV variables.

###### As example
```
TEST=12345
TEST1=4567
TEST2=qwerty
DB_PORT=5432
```

Right now generate config package. All ENVs will be added in code plus default ( *APP_PORT*, *KAFKA_BROKERS*, *SVC_TRACING_ZIPKIN* ).