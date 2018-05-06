# pack
This project is a sample Go app that implements a Graphql API which is a basic CRUD in Mongo

## Using service ##

Find below the usage of this service

### Queries ###

* Query a pack by code. It returns id, packcode, productid and name.

```sh
curl -g 'http://localhost:8287/graphql?query={byCode(packcode:"wh1000"){id,packcode,ownerid,resources{name,isfree},prodid,name}}'
```

* Query pack by its id. It returns id, packcode, productid and name.

```sh
curl -g 'http://localhost:8287/graphql?query={byID(id:"PACK_ID"){id,packcode,prodid,name}}'
```

* Query id pack by its code. It returns id, packcode, productid and name.

```sh
curl -g 'http://localhost:8287/graphql?query={idByCode(packcode:"wh13")}'
```

* Query pack by its product id. It returns id, packcode, productid and name.

```sh
curl -g 'http://localhost:8287/graphql?query={byProductID(productid:"13"){id,packcode,prodid,name}}'
```

* Query pack by its main keys: mno id, product id and pack code. It returns true if exists a pack with those keys.

```sh
curl -g 'http://localhost:8287/graphql?query={byKeys(mnoid:2,packcode:"wh13",productid:"13")}'
```

### Mutations ###

* Create a pack. returns boolean success, any code for reference and a message in an error case.

```sh
curl -XPOST -H 'Content-Type:application/graphql' -d 'mutation PackMutation { create(prodid:"1000",packcode:"wh1000",name:"saladino internet",desc:"navega como loco pana",imgurl:"/img/caasi/ss.png",kwds:"internet saladino",price:3400,ownerid:1,type:{id:1,name:"App"},mno:{id:2,name:"Claro"},term:{unit_id:4,unit:"dia",amount:2},currency:{id:3,name:"cop"}) { success, code, msg} }' http://localhost:8287/graphql
```

* Change the state of a pack. returns boolean success, any code for reference and a message in an error case.

```sh
curl -XPOST -H 'Content-Type:application/graphql' -d 'mutation PackMutation { changeState(id:"59dce5b6ea68afcfe60ae8cb",state:1){ success, code, msg} }' http://localhost:8287/graphql
```

* Change the product id of a pack. returns boolean success, any code for reference and a message in an error case.

```sh
curl -XPOST -H 'Content-Type:application/graphql' -d 'mutation PackMutation { changeProductID(id:"59dce5b6ea68afcfe60ae8cb",mnoid:2,productid:"13"){ success, code, msg} }' http://localhost:8287/graphql
```

* Change the code of a pack. returns boolean success, any code for reference and a message in an error case.

```sh
curl -XPOST -H 'Content-Type:application/graphql' -d 'mutation PackMutation { changePackCode(id:"59dce5b6ea68afcfe60ae8cb",mnoid:2,packcode:"wh13"){ success, code, msg} }' http://localhost:8287/graphql
```

* Change the name of a pack. returns boolean success, any code for reference and a message in an error case.

```sh
curl -XPOST -H 'Content-Type:application/graphql' -d 'mutation PackMutation { changeName(id:"59dce5b6ea68afcfe60ae8cb",newname:"Whatsapp weekend 13_1"){ success, code, msg} }' http://localhost:8287/graphql
```

* Change the description of a pack. returns boolean success, any code for reference and a message in an error case.

```sh
curl -XPOST -H 'Content-Type:application/graphql' -d 'mutation PackMutation { changeDescription(id:"59dce5b6ea68afcfe60ae8cb",newdesc:"Whatsapp para el fin de semana, para que hables con tus amigos todo el dia ois."){ success, code, msg} }' http://localhost:8287/graphql
```

* Change the image url of a pack. returns boolean success, any code for reference and a message in an error case.

```sh
curl -XPOST -H 'Content-Type:application/graphql' -d 'mutation PackMutation { changeImageUrl(id:"59dce5b6ea68afcfe60ae8cb",newimgurl:"/appdata/img/whatsappwknd2.png"){ success, code, msg} }' http://localhost:8287/graphql
```

* Change the keywords of a pack. returns boolean success, any code for reference and a message in an error case.

```sh
curl -XPOST -H 'Content-Type:application/graphql' -d 'mutation PackMutation { changeKeywords(id:"59dce5b6ea68afcfe60ae8cb",newkeywords:"internet whatsapp dia fin semana calidad"){ success, code, msg} }' http://localhost:8287/graphql
```

* Change the price of a pack. returns boolean success, any code for reference and a message in an error case.

```sh
curl -XPOST -H 'Content-Type:application/graphql' -d 'mutation PackMutation { changePrice(id:"59dce5b6ea68afcfe60ae8cb",newprice:5601){ success, code, msg} }' http://localhost:8287/graphql
```

* Change the type of a pack. returns boolean success, any code for reference and a message in an error case.

```sh
curl -XPOST -H 'Content-Type:application/graphql' -d 'mutation PackMutation { changeType(id:"59dce5b6ea68afcfe60ae8cb",newtype:{id:2,name:"App2"}){ success, code, msg} }' http://localhost:8287/graphql
```

* Change the mno owner of a pack. returns boolean success, any code for reference and a message in an error case.

```sh
curl -XPOST -H 'Content-Type:application/graphql' -d 'mutation PackMutation { changeMno(id:"59dce5b6ea68afcfe60ae8cb",productid:"13",packcode:"wh13",newmno:{id:5,name:"Virgin"}){ success, code, msg} }' http://localhost:8287/graphql
```

* Change the validity of a pack. returns boolean success, any code for reference and a message in an error case.

```sh
curl -XPOST -H 'Content-Type:application/graphql' -d 'mutation PackMutation { changeValidity(id:"59dce5b6ea68afcfe60ae8cb",newterm:{unit_id:5,unit:"week",amount:2}){ success, code, msg} }' http://localhost:8287/graphql
```

* Change the currency of a pack. returns boolean success, any code for reference and a message in an error case.

```sh
curl -XPOST -H 'Content-Type:application/graphql' -d 'mutation PackMutation { changeCurrency(id:"59dce5b6ea68afcfe60ae8cb",newcurrency:{id:2,name:"pe"}){ success, code, msg} }' http://localhost:8287/graphql
```

* Move the stock of a pack. returns boolean success, any code for reference and a message in an error case.

```sh
curl -XPOST -H 'Content-Type:application/graphql' -d 'mutation PackMutation { moveStock(id:"59ec341d18c4f5ec3732b165",amount:3){ success, code, msg} }' http://localhost:8287/graphql
```

* Delete a pack physically. returns boolean success, any code for reference and a message in an error case.

```sh
curl -XPOST -H 'Content-Type:application/graphql' -d 'mutation PackMutation { delete(id:"59dcf455ea68afcfe60aebd4"){ success, code, msg} }' http://localhost:8287/graphql
```

* replace resources of a pack. returns boolean success, any code for reference and a message in an error case.

```sh
curl -XPOST -H 'Content-Type:application/graphql' -d 'mutation PackMutation { replaceResources(id:"5a12211dcc7c76da03df50f7",newresources:[{id:1,name:"voz",units:"min",amount:200,isfree:false},{id:2,name:"datos",units:"mb",amount:200,isfree:false},{id:3,name:"sms",units:"sms",amount:200,isfree:true}]){ success, code, msg} }' http://localhost:8287/graphql
```

* delete resources of a pack. returns boolean success, any code for reference and a message in an error case.

```sh
curl -XPOST -H 'Content-Type:application/graphql' -d 'mutation PackMutation { deletePackResources (id:"5a07bbc9e82fd55107594491"){ success, code, msg} }' http://localhost:8287/graphql
```

* Check health of the service.

```sh
curl -XGET -H 'Content-Type:application/json' http://localhost:8287/health
```

## What is this repository for? ##

* Contains source code that implements pack management service.
* 0.1.0

## How do I get set up? ##

* This a go application.
* We use go dep.
* Dependencies.
  * [mgo - mongo driver](https://gopkg.in/mgo.v2)
  * [go dep](https://github.com/golang/dep)
* Database configuration
* How to run tests
  * you can run this command: go test ./...
* Deployment instructions

## Docker ##

To build in the OS that we want to run. This produces an executable binary named using the -o flag. I usually brew my own little shell script that automates this for me.

Since both OS X and our Linux-based container runs on the AMD64 CPU architecture we don�t need to set (and reset) the GOARCH env var. But if you are building something for an 32-bit OS or perhaps an ARM processor you would need to set GOARCH appropriately before building.

Below the steps to build our docker image for pack-search service.

* Export GOOS variable.

```sh
export GOOS=linux
```

* Build the project

```sh
CGO_ENABLED=0 go build -o pack-service-linux-amd64 -installsuffix cgo github.com/fernandoocampo/pack
```

* Create the docker image. When building a Docker container image, we usually tag it with a name usually using a [prefix]/[name] naming convention.

```sh
cd WORKSPACE_FOLDER/go/src/bitbucket.org/team-anfora
docker build -t team-anfora/pack-service pack/
```

* Run the docker container (-p format: ip:hostPort:containerPort)

```sh
docker run --rm --network=bridge -v /home/luisfer/appdata/pack/conf:/etc/pack/conf -p 8297:8287 --name pack team-anfora/pack-service
```

* For development purpose remember execute mongo container in this way

```sh
docker run --network=bridge --name mongo-database-2 -p 27017:27017 -d mongo:3.5
```

## Contribution guidelines ##

* Writing tests
* Code review
* Other guidelines

## Who do I talk to? ##

* Repo owner or admin
* Other community or team contact