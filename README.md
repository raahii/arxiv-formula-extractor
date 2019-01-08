# Arxiv Equations



[Arxiv Equations](https://arxiv-equations.netlify.com/) provides latex format equations from arxiv paper.

This app is made with `Golang`, `Vue.js` and `MySQL`.



![](https://user-images.githubusercontent.com/13511520/50848768-70ddb900-13b8-11e9-9c17-d18f5791ac5f.png)



## Requirement

- Golang 

  - [dep](https://github.com/golang/dep)

    ```shell
    go get -u github.com/golang/dep/cmd/dep
    ```

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

- backend

  ```shell
  dep ensure
  ```

- frontend

  ```shell
  cd frontend
  npm install --save
  ```

### Start development

- backend server

  ```shell
  go run server.go
  ```

- frontend 

- ```shell
  cd frontend
  npm run dev
  ```

