package protection

import (
	"go-post-api-projects/logger"
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

// >>>  Get Secret Key From .env File
var secretKey = []byte(getSecret())

func getSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if(secret == ""){
		logger.AppLogger.Error.Println("JWT_SECRET not set in environment")
	}
	return secret
}


/// >> Generated Token & Every Day 12:00PM expire its
func GenerateToken(username string) (string, error) {
    // 👉 BD timezone set করা
    loc, _ := time.LoadLocation("Asia/Dhaka")

    presentTime := time.Now().In(loc)

    //  আজকের রাত 12:00 (midnight)
    expireTime := time.Date(presentTime.Year(),presentTime.Month(),presentTime.Day(), 0, 0, 0, 0,loc,)

    // যদি এখন সময় already রাত 12টা পার হয়ে যায়
    if !presentTime.Before(expireTime) {
        expireTime = expireTime.Add(24 * time.Hour)
    }

    claims := jwt.MapClaims{
        "username": username,
        "exp":      expireTime.Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secretKey)
}