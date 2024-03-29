package doc

import (
	"encoding/json"
	"github.com/CloverOS/gin-swagger-bootstrap/bootstrap"
	"github.com/CloverOS/goctl-swag/gen"
	"net/http"
	"path/filepath"
	"strings"
)

func WrapHandler() http.HandlerFunc {
	return WrapHandlerWithGroup(Group)
}

func WrapHandlerWithGroupObject(group []gen.SwaggerGroupObject) http.HandlerFunc {
	bytes, err := json.MarshalIndent(group, "", "    ")
	if err != nil {
		return nil
	}
	return WrapHandlerWithGroup(string(bytes))
}

func WrapHandlerWithGroup(group string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			_, _ = rw.Write([]byte("Method Not Allowed"))
			return
		}
		i := strings.LastIndex(r.URL.Path, "/")
		if i == -1 {
			return
		}
		path := r.URL.Path[i+1:]
		if path == "" {
			path = "index.html"
		}
		switch filepath.Ext(path) {
		case ".html":
			rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		case ".css":
			rw.Header().Set("Content-Type", "text/css; charset=utf-8")
		case ".js":
			rw.Header().Set("Content-Type", "application/javascript")
		case ".png":
			rw.Header().Set("Content-Type", "image/png")
		case ".json":
			rw.Header().Set("Content-Type", "application/json; charset=utf-8")
		}
		switch path {
		case "spider-api.json":
			_, _ = rw.Write([]byte(Doc))
		case "group.json":
			_, _ = rw.Write([]byte(group))
		case "index.html":
			readFile, err := bootstrap.ReadFile(path)
			if err != nil {
				bootstrap.Handler.ServeHTTP(rw, r)
				return
			}
			_, _ = rw.Write(readFile)
		default:
			tmp := strings.Replace(r.URL.Path, "/swagger/", "", 1)
			readFile, err := bootstrap.ReadFile(tmp)
			if err != nil {
				bootstrap.Handler.ServeHTTP(rw, r)
				return
			}
			_, _ = rw.Write(readFile)
		}
	}
}

const Doc = `{
    "consumes": [
        ""
    ],
    "produces": [
        ""
    ],
    "schemes": [
        ""
    ],
    "swagger": "2.0",
    "info": {
        "contact": {},
        "license": {}
    },
    "host": "localhost:8888",
    "basePath": "/",
    "paths": {
        "/driver/aliyuncallback": {
            "get": {
                "produces": [
                    ""
                ],
                "tags": [
                    ""
                ],
                "operationId": "AliyunCallback",
                "parameters": [
                    {
                        "type": "string",
                        "name": "code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/AliYunCallbackResp"
                        }
                    }
                }
            }
        },
        "/spider/create": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    ""
                ],
                "tags": [
                    ""
                ],
                "operationId": "SpiderHandler",
                "parameters": [
                    {
                        "description": "  ",
                        "name": "SpiderRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/SpiderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/SpiderResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "AliYunCallbackResp": {
            "type": "object",
            "title": "AliYunCallbackResp",
            "properties": {
                "accessToken": {
                    "description": " accesstoken",
                    "type": "string",
                    "required": [
                        "true"
                    ]
                },
                "code": {
                    "description": " 状态码",
                    "type": "uint32",
                    "required": [
                        "true"
                    ]
                },
                "message": {
                    "description": " 信息",
                    "type": "string",
                    "required": [
                        "true"
                    ]
                }
            }
        },
        "SpiderRequest": {
            "type": "object",
            "title": "SpiderRequest",
            "properties": {
                "uid": {
                    "type": "array",
                    "required": [
                        "true"
                    ],
                    "items": {
                        "type": "uint32"
                    }
                }
            }
        },
        "SpiderResponse": {
            "type": "object",
            "title": "SpiderResponse",
            "properties": {
                "message": {
                    "type": "string",
                    "required": [
                        "true"
                    ]
                }
            }
        }
    }
}`

const Group = `[
    {
        "name": "spider-api",
        "url": "../swagger/spider-api.json",
        "swaggerVersion": "2.0",
        "location": ""
    }
]`
