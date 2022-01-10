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

## Setup

Сheck that you have two free ports:

* 8080  
* 11211 

if not, replace with yours!

* main.go
```
if err := srv.Run("YOUR_PORT", internal.GetInterval())

```

* docker-compose.yml

```
  ports:
  - "YOUR_PORT:11211"
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
http://localhost:8080/1/20
```
The screen displays the fibonacci values in the specified interval.
```
1 0
2 1
3 1
4 2
5 3
6 5
7 8
8 13
9 21
10 34
11 55
12 89
13 144
14 233
15 377
16 610
17 987
18 1597
19 2584
20 4181
```
