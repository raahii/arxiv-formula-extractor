# Arxiv Equations

![](https://user-images.githubusercontent.com/13511520/50848768-70ddb900-13b8-11e9-9c17-d18f5791ac5f.png)

[Arxiv Equations](https://arxiv-equations.netlify.com/) provides latex format equations from arxiv paper.

This app is made with `Golang`, `Vue.js` and `MySQL`.

## Requirement

- Golang 
  - [dep](https://github.com/golang/dep), `go get -u github.com/golang/dep/cmd/dep`
- npm
- mysql 

## Getting Started

```
go get github.com/raahii/arxiv-equations
```

### Environment variables

```
export DB_USER=<user>
export DB_PASS=<password>
export DB_NAME=<dbname, ex)arxiv_equations>
export BACKEND_BASEURL=<backend url, ex)http://localhost:1323">
```

### Install dependencies

```shell
# backend 
dep ensure

# frontend 
cd frontend; npm install --save
```

### Start development


```shell
# backend 
go run server.go

# frontend 
cd frontend; npm run dev
```

## TODO:

- [ ] remove vue router

- [ ] updating paper function

- [ ] exclusive processing for pasing latex source

- [ ] accurate tex parsing (by using pandoc or something)

- [ ] obtain number of equation

- [ ] obtain number of equation
