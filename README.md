## Table of contents
* [General info](#general-info)
* [Technologies](#technologies)
* [Setup](#setup)
* [Example](#example)

## General info
REST API for getting values of Fibonacci numbers by their ordinal number.

## Technologies
Project is created with:
* GO
* Memcached
* REST
* gin-gonic
* Docker
* gRPC

## Setup

Сheck that you have free ports:

* 8080  
* 11211
* 50051

if not, replace with yours!


* docker-compose.yml

```
    ports:
      - "8080:8080"
      - "50051:50051"
    environment:
      - PORT=8080
      - MC_PORT=11211
      - GRPC_PORT=50051
    ports:
      - "11211:11211"
```


To run this project :

Expand docker-compose, go to localhost with port:8080, where:

* x - the beginning of the interval

* y - end of interval

```
 http://localhost:8080/x/y
```
When you deployed the docker-compose make sure, that containers are lifted:
1. api
2. memcached

## Example
```
http://localhost:8080/5/12
```
The screen displays the fibonacci values in the specified interval.
```
	
0	
id	5
value	3
1	
id	6
value	5
2	
id	7
value	8
3	
id	8
value	13
4	
id	9
value	21
5	
id	10
value	34
6	
id	11
value	55
7	
id	12
value	89
```
