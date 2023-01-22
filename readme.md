# API Mocker
[![Go Reference](https://pkg.go.dev/badge/github.com/plin2k/api-mocker.svg)](https://pkg.go.dev/github.com/plin2k/api-mocker)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/plin2k/api-mocker)

![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/plin2k/api-mocker)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/plin2k/api-mocker)
![GitHub last commit](https://img.shields.io/github/last-commit/plin2k/api-mocker)


## Installation
    go install github.com/plin2k/api-mocker@latest

### If you don't have Go
    brew install go

### If there is no brew
    Follow instructions - https://go.dev/doc/install

## Support Protocols
    HTTP
    WebSocket
    gRPC - Coming soon...
    
## Features
    Route grouping
    HTTP:
        Headers
        Cookies
        Methods:
            - GET
            - POST
            - DELETE
            - HEAD
            - OPTIONS
            - PATCH
            - PUT
            - ANY
        Various responses:
            - File
            - Download file
            - HTML file
            - JSON
            - XML
            - TOML
            - YAML
            - Simple String
    Delay before response
    Websockets
    Fake and random data
    Set the maximum number of CPUs
    Web documentation - Coming soon...
    Authorization - Coming soon...
    TLS - Coming soon...
    Source file validator - Coming soon...
    Api dublicator - Coming soon...
    Swagger integration - Coming soon...

## Example usage

### Execute
    api-mocker --help
    api-mocker run --host="127.0.0.1" --port=8080 --src=api-mocker.xml --gomaxprocs=4
    
### Example HTTP api-mocker.xml

```xml
<xml protocol="http">
    
    <name>API v1</name>
    <description>API Mocker Protocol - HTTP</description>

    <api>
        <group path="api/v1" description="API v1">
            
            <route path="file" description="Simple file show" method="GET" status-code="200">
                <header key="Content-Type" value="application/xml" description="my-first-header"/>
                <header key="X-Referral-ID" value="999" description="my-second-header"/>
                
                <cookie name="cookie" description="my-cookie" value="hello world" max-age="3600" path="/" domain="example.com" secure="true" http-only="true"/>
                
                <response delay="10" type="file" description="XML file">source.xml</response>
            </route>

            <group path="healthcheck" description="Healthcheck services">
                
                <route path="mysql" description="Check health status MySQL" method="ANY" status-code="200">
                    <response type="json" description="JSON text">{"status":"success"}</response>
                </route>
                
                <route path="redis" description="Check health status Redis" method="ANY" status-code="200">
                    <response type="json" description="JSON text">{"status":"success"}</response>
                </route>   
                
            </group>
        </group>

        <group path="api/v2" description="API v2">
            
            <header key="Authorization" value="{{ randJWT }}"/>
            
            <route path="json" description="Simple JSON show" method="GET" status-code="200">
                
                <header key="Authorization" value="{{ randJWT }}" description="Redeclared"/>
                <header key="X-Referral-ID" value="999" description="my-second-header"/>
                
                <response type="json" description="JSON text">{"name":"JSON"}</response>
            </route>
        </group>

    </api>
    
</xml>
```

### Example WebSocket api-mocker.xml

```xml
<xml protocol="websocket">
    <name>API V1</name>
    <description>API Mocker Protocol - WebSocket</description>

    <onopen description="OnOpen message">
        <response delay="10" type="text" description="Can respond with message type Text and Binary">I'm Opened a Connection</response>
    </onopen>

    <onclose description="OnClose message">
        <response type="text" description="Can respond with message type Text and Binary">I'm Closed a Connection</response>
    </onclose>

    <onmessage description="Can accept message type Text and Binary">
        <response type="binary" description="Can respond with message type Text and Binary">My Reply Message</response>
    </onmessage>

    <onerror description="Some error has occurred">
        <response type="text" description="Can respond with message type Text and Binary">My Error Message</response>
    </onerror>

    <ping description="Reply to Ping message">
        <response type="text" description="Can respond with message type Text and Binary">Ping</response>
    </ping>

    <pong description="Reply to Pong message">
        <response type="text" description="Can respond with message type Text and Binary">Pong</response>
    </pong>

    <while>
        <message delay="10" count="10">
            <response type="text" description="Sends a message every 10 sec. Can respond with message type Text and Binary">
                {"msg":"every 10 sec","event":"new_chat"}
            </response>
        </message>

        <message delay="3" count="10" >
            <response type="text" description="Sends a message every 3 sec. Can respond with message type Text and Binary">
                {"msg":"every 3 sec","event":"new_message"}
            </response>
        </message>
    </while>
</xml>
```


### Fakers

```
Random between dates {{ randDate "01-01-1970" "01-01-2022" }} -> "04-11-1997"
Random Bool {{ randBool }} -> true
Random String {{ randString 4 }} -> "rHdD"
Random Int {{ randInt 10 20 }} -> 15
Random Float {{ randFloat 10 1000 }} -> 43.24
        
Random Array Ints {{ randArrayInt 6 10 100 }} -> [15,35,22,10,39,99]
Random Array Strings {{ randArrayString 6 3 }} -> ["deZ","Zjq","sdg","pDp","dWW","WMN"]

Random Firstname (0 - male, 1 - female) {{ randFirstName 1 }} -> "James"
Random Lastname {{ randLastName }} -> "Jones"
Random Sex {{ toTitle randSex }} -> "female"
        
Random Email {{ randEmail }} -> "aleksandr@kalink.in" 
Random Domain {{ randDomain }} -> "google.com" 
        
Random Lang Code {{ randLangCode }} -> "en"
Random Country {{ randCountry }} -> "Japan" 
Random City {{ randCity }} -> "Lefkosia" 
        
Random JWT {{ randJWT }} -> "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCIsInR5MjIyMjIyMjIyMjIyMzIzMjNwIjoiSldkc2FkYXNkYXNzYVQifQ.eyJzdWIiOiIxMjM0NWRzYWRhc2RhczY3ODkwIiwibmFtZSI6IkpvZHNhZGFzZGFzZGFzZGFzZGFzZGFzZGFzZGRzYWRhc2RobiBEb2UiLCJpYXQiOjE1MTYyMzkwMjJ9.iVsrj7dQjLvA1YUUhYUGzlnfPNS2zfUcNBi-n8rK_ls"
Random Cookie {{ randCookie }} -> "id=a3fWa; Expires=Thu, 31 Oct 2021 07:28:00 GMT;"
Random UUID {{ randUUID }} -> "f7ec47d0-9a7e-11ed-a8fc-0242ac120002"
Random Hash (md5, sha1, sha256, sha512) {{ randHash "sha256" }} -> "56f3403d88c2423f711b946be61d15a7eec895635fd0e90243b6cd8526571fec"

To Title {{ toTitle ("hello") }} -> "Hello"
To Lower {{ toLower ("HELLO") }} -> "hello"
To Upper {{ toUpper ("hello") }} -> "HELLO"
Format date {{ formatDate "01-01-2014" }} ->  Mon, 02 Jan 2006
```

### Usage Faker

```xml
<xml protocol="http">
    
    <name>API v1</name>
    <description>
        API Mocker Protocol - HTTP. 
        Author {{ randFirstName 0 }} {{ randLastName }}
    </description>

    <api>
        <group path="api/v{{ randInt 1 5 }}" description="Root group">
            <route path="json" description="Simple JSON show" method="GET">
                <response type="json" description="JSON text">
                    {
                    "name":"{{ randFirstName 1 }} {{ randLastName }}",
                    "number":{{ randInt 10 20 }},
                    "float": {{ randFloat 10 1000}},
                    "arrayInt": {{ randArrayInt 6 10 100 }},
                    "arrayString": {{ randArrayString 6 10 }},
                    "randSex": "{{ toTitle randSex }}",
                    "randDate": "{{ randDate "01-01-2014" "01-01-2022" }}",
                    "randUUID": "{{ randUUID }}",
                    "randHash": "{{ randHash "sha256" }}",
                    "randCity": "{{ randCity }}",
                    "randCountry": "{{ randCountry }}",
                    "randEmail": "{{ randEmail }}",
                    "randDomain": "{{ randDomain }}",
                    }
                </response>
            </route>
        </group>

    </api>
    
</xml>
```