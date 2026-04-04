package protection

import (
	"go-post-api-projects/logger"
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

// >>>  Get Secret Key From .env File >> Convert Byte For Mutable but human rot readable only machine readable
var secretKeys = []byte(getSecretKey())
func getSecretKey() string{
    secretKey := os.Getenv("JWT_SECRET")
    if(secretKey == ""){
        logger.AppLogger.Error.Println("Secret Key Not Found in evn file")
    }
    return secretKey
}


/// >> Generated Token & Every Day 12:00PM expire its
func GenerateToken(username string) (string,error){
    // >> Get BD Location
    loc,_:= time.LoadLocation("Asia/Dhaka")
    currentTime := time.Now().In(loc)
    expairTime := time.Date(currentTime.Year(),currentTime.Month(),currentTime.Day(),0,0,0,0,currentTime.Location())

    // >> Check: Is it already past 12 ? means currentTime ≥ expairTime
    if(!currentTime.Before(expairTime)){
        expairTime = expairTime.Add(24 * time.Hour)
    }

    claims := jwt.MapClaims{"username" : username,"exp" : expairTime.Unix()}
    createToken := jwt.NewWithClaims(jwt.SigningMethodES512,claims)
    return createToken.SignedString(secretKeys)
}


/// >>> Login (token create) 
