module github.com/Rohanrevanth/chat-demo-go/chat-demo

go 1.23.1

replace github.com/Rohanrevanth/chat-demo-go/database => ../database

require (
	github.com/Rohanrevanth/chat-demo-go/database v0.0.0-00010101000000-000000000000
	github.com/Rohanrevanth/chat-demo-go/http v0.0.0-00010101000000-000000000000
	github.com/Rohanrevanth/chat-demo-go/websockets v0.0.0-00010101000000-000000000000
)

require (
	github.com/Rohanrevanth/chat-demo-go/auth v0.0.0-00010101000000-000000000000 // indirect
	github.com/Rohanrevanth/chat-demo-go/controllers v0.0.0-00010101000000-000000000000 // indirect
	github.com/Rohanrevanth/chat-demo-go/models v0.0.0-00010101000000-000000000000 // indirect
	github.com/Rohanrevanth/chat-demo-go/routes v0.0.0-00010101000000-000000000000 // indirect
	github.com/bytedance/sonic v1.11.6 // indirect
	github.com/bytedance/sonic/loader v0.1.1 // indirect
	github.com/cloudwego/base64x v0.1.4 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/gin-contrib/cors v1.7.2 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.10.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.20.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.7 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	golang.org/x/arch v0.8.0 // indirect
	golang.org/x/crypto v0.27.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/text v0.18.0 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/driver/sqlite v1.5.6 // indirect
	gorm.io/gorm v1.25.12 // indirect
)

replace github.com/Rohanrevanth/chat-demo-go/http => ../http

replace github.com/Rohanrevanth/chat-demo-go/controllers => ../controllers

replace github.com/Rohanrevanth/chat-demo-go/routes => ../routes

replace github.com/Rohanrevanth/chat-demo-go/auth => ../auth

replace github.com/Rohanrevanth/chat-demo-go/models => ../models

replace github.com/Rohanrevanth/chat-demo-go/websockets => ../websockets