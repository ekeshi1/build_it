[![Work in Repl.it](https://classroom.github.com/assets/work-in-replit-14baed9a392b3a25080506f3b7b6d57f295ec2978f6f33ec97e36a161684cbe9.svg)](https://classroom.github.com/online_ide?assignment_repo_id=413595&assignment_repo_type=GroupAssignmentRepo)

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
* Run service with docker-compose <br/>
Start <br/>
```
docker-compose up -d
```
Rebuild  <br/>
```
docker-compose build
```
Stop  <br/>
```
docker-compose down -v
```
* Testing
[Aceptance Test](https://github.com/ekeshi1/build_it/blob/CC5-CC8-Orkhan/documentation/Acceptance%20test.yml)
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


