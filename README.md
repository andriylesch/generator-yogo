# generator-yogo

![Logo](https://github.com/ricardo-ch/generator-golang/blob/master/logo_small.png)

this generator will allow you to create projects like
- Console (basic "Hello world")
- GO-KIT microservice
- Rest API microservice
- KAFKA producer
- KAFKA consumer
- Worcker

So right now as developer you don't need to create each time the same skeleton of project. Just generate and focus for implementing your logic.

## Install 

Clone this repository (git clone)

If yeoman is not installed please

```bash
npm install -g yo
```

Open the solution
```bash
npm install
sudo npm link
```

now that we are ready to test.

```bash
mkdir [ProjectName]
cd [ProjectName]
yo golang
```

You can test it it will work only for (go-kit and restapi applications) :

```bash
go test ./... -tags=unit -v
```

In order to run it, you have to create the endpoints in your main.go
To help you a _main.todelete.txt is generated.
Copy/paste and adpat maybe a little bit.

```bash
go run main.go
```

Then you'll be able to test the endpoints :

```bash
GET /[YOUR_PACKAGE_NAME]/{ID}

// Response
{
    "ID" : {ID}
}
```

# Generate project via docker 

- install docker on your local system 
- clone this repository (git clone)
- open terminal and open `generator-golang` folder
- type command bellow

for windows users
``` bash
docker build -t generator-golang:latest .
docker run -it -v {YOUR_LOCAL_PATH}:/home/yeoman --name generator-golang-container generator-golang
```

for linux and macOS users

```bash
make docker LOCAL_PATH=YOUR_LOCAL_PATH
```
where **YOUR_LOCAL_PATH** - local folder where will be generated your project

create your folder
```bash
mkdir [ProjectName]
cd [ProjectName]
yo golang
```

# How to create new package in existed microservice.

Developer needs to open in terminal his project folder.

```bash
yo golang:pkgg
```

List of packages what is possible to create and use :
- config
- GO-KIT endpoint

**FYI: Soon will exist possibility to create :**
- `kafkapkg (consumer/producer)`
- `rest endpoint`
-  etc  

### Config package

Developer has possibility to add **config** package in his application. 
In this case it will be generated standart config package with structure 

```
- config
    |- config.toml
    |- config.toml.example
    |- config.go
    |- config_test.go

```

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
  



