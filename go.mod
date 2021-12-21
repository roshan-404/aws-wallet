module aws-wallet

// +heroku goVersion go1.15
go 1.15

require github.com/gin-gonic/gin v1.7.7

require (
	github.com/go-playground/validator/v10 v10.9.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/joho/godotenv v1.4.0
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	golang.org/x/crypto v0.0.0-20211209193657-4570a0811e8b
	golang.org/x/sys v0.0.0-20211214234402-4825e8c3871d // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

require (
	github.com/aws/aws-sdk-go-v2 v1.11.2
	github.com/aws/aws-sdk-go-v2/config v1.11.0
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.10.0
)

require (
	github.com/aws/aws-sdk-go-v2/credentials v1.6.4
	github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue v1.4.4
)

require (
	github.com/aws/aws-sdk-go v1.42.23
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gofrs/uuid v4.2.0+incompatible
)

require (
	github.com/go-playground/locales v0.14.0
	github.com/go-playground/universal-translator v0.18.0
	github.com/ugorji/go v1.2.6 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0
)
