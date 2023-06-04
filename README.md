# Tiktok Tech Immersion 2023

### Overview
Implementing HTTP and RPC servers using Kitex (Golang). The assignment requires two endpoints `/send` and `/pull`. 

- `send/` - POST request with the following payload:
```json
{
    "chat": "str",
    "text": "str",
    "sender": "str"
}
```

- `pull/` - GET request with the following parameters:
```json
{
    "chat": "str",
    "cursor": "int64",
    "limit": "int32",
    "reverse": "boolean"
}
```

