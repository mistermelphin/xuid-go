# XUID

XUIDs are URL-friendly compressed UUIDs [www.xuid.org](http://www.xuid.org)

## Why XUIDs?

A UUID v4 is a great choice as a primary key for your database tables.

But, they are quite long to use in URLs and databases : 32 alphanumeric characters and four hyphens (36 characters total)

A XUID is a UUID, converted into a 128-bit value, converted into a base64 string (stripped from padding characters), 
then converted into url-safe base64 (replacing `+` and `/` into `-` and `_` respectively).

This gives you a 22 character string, safe to use in URLs.

It also safely decodes into a full UUID string again.

## Benefits

* Takes less space in database fields
* Shorter URLs
* Un-guessable primary keys in URLs

## Install

```shell
go get github.com/mistermelphin/xuid-go
```

## Usage

Ways to create XUID
```go
id := xuid.New() // generate new XUID

fmt.Println(id) // prints XUID as string
fmt.Println(id.String()) // same result as previous

id = xuid.MustParse("1Od2S4tZTQyMkM3NCCTP9Q") // parses XUID string
id = xuid.MustParse("b5838375-c297-4ecc-a002-ad42e8c4f515") // parses UUID string and converts it to XUID
```

Converts XUID to UUID
```go
id := xuid.New()

fmt.Println(id.ToUUIDString()) // converts to UUID string

uid := id.ToUUID() // converts to google.uuid UUID
```

## Licenses

All source code is licensed under the [MIT License](LICENSE).

