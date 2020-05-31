# pinkbubbleapi

The application uses following libraries:

go get -u golang.org/x/crypto/bcrypt <br/>
go get -u github.com/go-sql-driver/mysql <br/>
go get github.com/aws/aws-sdk-go <br/>
go get github.com/gin-gonic/gin <br/>
go get -u github.com/gin-contrib/cors <br/>
export GO111MODULE=on <br/>
go get -u github.com/appleboy/gin-jwt/v2 <br/>
export GO111MODULE=off <br/>
go get -u github.com/jinzhu/gorm <br/>
go get -u github.com/dgrijalva/jwt-go <br/>
go build api \n <br/>

environment vars:
export DBUSER=<DATABASE_USER_NAME>
export DBPASSWORD=<DATABASE_PASSWORD>
export DBNAME=<DATABASE_NAME>
export BUCKET_NAME=<S3_BUCKET_NAME>
