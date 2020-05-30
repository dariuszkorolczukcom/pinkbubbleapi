# pinkbubbleapi

The application uses following libraries:

go get -u golang.org/x/crypto/bcrypt
go get -u github.com/go-sql-driver/mysql
go get github.com/aws/aws-sdk-go
go get github.com/gin-gonic/gin
go get -u github.com/gin-contrib/cors
export GO111MODULE=on
go get -u github.com/appleboy/gin-jwt/v2
export GO111MODULE=off
go get -u github.com/jinzhu/gorm
go get -u github.com/dgrijalva/jwt-go
go build api
