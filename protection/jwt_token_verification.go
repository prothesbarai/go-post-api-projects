package protection

import (
	"go-post-api-projects/logger"
	"os"
	"time"
	"github.com/golang-jwt/jwt"
)

var secretKey = []byte(getSecret())
func getSecret() string{
	secret := os.Getenv("JWT_SECRET")
	if(secret == ""){logger.AppLogger.Error.Println("Not Found SECRET KEY in env file")}
	return secret
}


func GeneratedToken(username string) (string,error){
	presentTime := time.Now()
	
	expireTime := time.Date(presentTime.Year(),presentTime.Month(),presentTime.Day(),12,0,0,0,presentTime.Location())

	if presentTime.After(expireTime){expireTime = expireTime.Add(24 * time.Hour)}


}