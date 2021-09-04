# Geopath

An API and web app written in golang :

The API endpoints help to create and get geojson data and get distances and durations.

The web app allows to represent paths(linestrings) in a map


## requirements:

## Install 


first clone this repository

the app will run in a docker container and the
docker commands are shortcuted in a makefile.



to build the image run:

```
$ make build

```

you can list more commands shortcuts by running the help command:

```
$ make help
```

## Run

start docker

```

$ make start

```
  and launch the web server 

```
$ make up
```
the server will be running at : 

http://localhost:10000

## Usage

### API

The backend contains a REST API built with flask and flak_restfull.

The API has 4 endpoints:


 #### /getPath

  return the data in a geojson representation:

 `GET /getPath`

 curl &nbsp; -i &nbsp; -H &nbsp; 'Accept: application/json'&nbsp;  http://localhost:10000/getPath

 will return :

 add a new feature to the feature collection:

 `POST /getPath`

 where data to post is in this json format:
```
{
    "pathname": pathname
    "coordinates":coordinates
}
```
 where   ```pathname ```  is a string and  ```coordinates ```  is  an array of 4 dimensions arrays of floats

example:


```
[

        [
    
          longitude, 

          latitude,

          altitude,

          timestamp
        ],
        [
    
          longitude1, 

          latitude1,

          altitude1,

          timestamp1

],

...
]
```

 
 #### /getDistance

 
  list of all saved paths distances:

 `GET /getDistance`

 curl &nbsp; -i &nbsp; -H &nbsp; 'Accept: application/json' &nbsp; http://localhost:10000/getDistance

 will return:
 ```
 HTTP/1.1 200 OK
Content-Type: application/json
Date: Fri, 03 Sep 2021 15:28:01 GMT
Content-Length: 80

{
"path1":75.05008730255076,
"path2":163.4858058443857,
"path3":30.485916554839854
}
 ```
 you can add a parameter to the request to get the distance of a specific path

 example:

 curl &nbsp; -i &nbsp; -H &nbsp; 'Accept: application/json' &nbsp; http://localhost:10000/getDistance?path=path1
 
 will return :
 ```
 HTTP/1.1 200 OK
Content-Type: application/json
Date: Fri, 03 Sep 2021 15:32:11 GMT
Content-Length: 17

75.05008730255076

 ```

#### /getDuration

list all saved paths durations:

 `GET /getDuration`

 curl &nbsp; -i &nbsp; -H &nbsp; 'Accept: application/json' &nbsp;  http://localhost:10000/getDuration

 will return :

  ```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Fri, 03 Sep 2021 15:35:25 GMT
Content-Length: 40

{"path1":9000,
"path2":9000,
"path3":3000
}

 ```

 you can add a parameter to the request to get the duration of a specific path

 example:

 curl &nbsp; -i &nbsp;  -H &nbsp; 'Accept: application/json' &nbsp;  http://localhost:10000/getDuration?path=path1

will return:

```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Fri, 03 Sep 2021 15:40:25 GMT
Content-Length: 4

9000

```
#### /getPathNames

list of all saved paths names:

 `GET /getDuration`

 example:
 curl &nbsp; -i &nbsp; -H &nbsp; 'Accept: application/json'&nbsp; http://localhost:10000/getPathNames
 
 will return:

 ```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Fri, 03 Sep 2021 15:52:05 GMT
Content-Length: 25

["path1","path2","path3"]

 ```


### web app

The  web app is UI to represent a linestring in a map using leaflet.js including 3 functionalities:
* select the path(linestring) to visualize on the map 
   and show distance and duration:

   use the select input to choose the path(s) you want to visualize
*  save a new path (add new feature to the linetrings featurecollection):

   click the "add new path" button a modal will show up and allow you to create and save a new path




## Next todos

In the future we will :

* add ArangoDB database

* improve the ui