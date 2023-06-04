# Tiktok Tech Immersion 2023

### Overview
Implementing HTTP and RPC servers using Kitex (Golang). The assignment requires two endpoints `/send` and `/pull`. 

- `api/send/` - POST request with the following payload:
```json
{
    "chat": "str",
    "text": "str",
    "sender": "str"
}
```

- `api/pull/` - GET request with the following parameters:
```json
{
    "chat": "str",
    "cursor": "int64",
    "limit": "int32",
    "reverse": "boolean"
}
```

### Initializing

#### 1. Clone Repository
```shell
git clone https://github.com/joshjms/tiktok-tech-immersion.git
```

#### 2. Run docker-compose
```shell
cd tiktok-tech-immersion
docker-compose up --build
```

#### 3. Test Endpoint
You can test the endpoint at port `8080`. 
