## Overview
A single page application skeleton using Vue.js with a go-chi API backend.

## Requirements
- [Node.js](https://nodejs.org/en/)
- [Vue CLI](https://github.com/vuejs/vue-cli)
- [Go v1.11+ installed and configured](https://golang.org/dl/).
- Configure GOPATH environmental variable

```
git clone https://github.com/zooraze/chi-vue-spa.git
cd chi-vue-spa
```

### Local Testing

```
cd server
go build
./server
```

Browse to http://localhost:3333.

Watch and compile :

```
npm run serve
```

Open http://localhost:8080 in a browser.


### Frontend

Setup:

```
cd web
npm install
```

Compile and minify:

```
npm run build
```

## Deploy

### Google Cloud
```gcloud app deploy```


## 3rd Party Libraries

[go-chi](https://github.com/go-chi/chi): API server (backend)

[Vue.js](https://github.com/vuejs/vue): single-page application (frontend)
