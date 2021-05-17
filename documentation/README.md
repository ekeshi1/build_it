# Project BuildIT
In this project we have build a service to support the operations of BuildIT. BuildIT is a construction company specialized in public works (roads, bridges, pipelines, tunnels,  railroads, etc.). The goal of this project is to implement a service-oriented system in order to support the business processes.
We have created a backend service in Golang, while following the  the Golang hexagonal architecture guidelines :
* Packaging
* Struct-oriented
* Dependency injection
* Interfaces

### Built With
The technologies used in this project are :
* Go
* PostgreSQL
* Docker
* Hetzner cloud

## Getting Started
In this sections we will describe the requirements and the steps to set up the environment to run the service
### Installation
* Fork the project to your account and clone it to your local system.
```
git clone https://github.com/cs-ut-ee/project-group-8.git 
```
* Running the service
```
go run main.go
```
* Testing
[Aceptance Test](https://github.com/cs-ut-ee/project-group-8/blob/CC5-CC8-Orkhan/documentation/Acceptance%20test.yml)
### Structure
1. main.go 
2. pkg
  * handlers
  * internald
      * domain
      * ports
  * repositories
  * services
3. tests
4. Documentation

### Project management
1. [Project management](https://github.com/cs-ut-ee/project-group-8/tree/CC5-CC8-Orkhan/documentation/Project%20management%20-%20team%20log)
