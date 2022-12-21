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
    HTTP/1.1
    HTTP/2 - Coming soon...
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
    Web documentation - Coming soon...
    Authorization - Coming soon...
    TLS - Coming soon...
    Source file validator - Coming soon...
    Api dublicator - Coming soon...
    Swagger integration - Coming soon...

## Example usage

### Execute
    api-mocker --port=8080 --src=source.xml
    
### Example source.xml

```xml
<xml protocol="http11">
    
    <name>API v1</name>
    <description>API Mocker Protocol - HTTP</description>

    <api>
        <group pattern="api/v1" description="API v1">
            
            <route pattern="file" description="Simple file show" method="GET" status-code="200">
                <header name="Content-Type" value="application/xml" description="my-first-header"/>
                <header name="X-Referral-ID" value="999" description="my-second-header"/>
                
                <cookie name="cookie" description="my-cookie" value="hello world" max-age="3600" path="/" domain="example.com" secure="true" http-only="true"/>
                
                <delay value="10" description="my-delay"/>
                
                <return type="file" description="XML file">source.xml</return>
            </route>

            <group pattern="healthcheck" description="Healthcheck services">
                
                <route pattern="mysql" description="Check health status MySQL" method="ANY" status-code="200">
                    <return type="json" description="JSON text">{"status":"success"}</return>
                </route>
                
                <route pattern="redis" description="Check health status Redis" method="ANY" status-code="200">
                    <return type="json" description="JSON text">{"status":"success"}</return>
                </route>   
                
            </group>
            
        </group>

        <route pattern="json" description="Simple JSON show" method="GET" status-code="200">
            <header name="X-Referral-ID" value="999" description="my-second-header"/>
            <return type="json" description="JSON text">{"name":"JSON"}</return>
        </route>

    </api>
    
</xml>
```
    