package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		// TODO: answer here
		

		var claims model.Claims
		cookie, err := ctx.Cookie("session_token")
		
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) && ctx.Request.URL.Path != "/" {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "Unauthorized",
				})
			}
			ctx.AbortWithStatusJSON(http.StatusSeeOther, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		// Parsing JWT token pada cookie
		token, err := jwt.ParseWithClaims(cookie, &claims, func(token *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil // Ganti dengan secret key yang sesuai
		})
		
		if err != nil {
			// Parsing token gagal, mengembalikan respon HTTP 401 atau 400 tergantung dari jenis error yang terjadi
			if err == jwt.ErrSignatureInvalid {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			} else {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
			}
			return
		}

		// Mendapatkan claims dari token
		claimsData, ok := token.Claims.(*model.Claims)
		if !ok || !token.Valid {
			// Token tidak valid, mengembalikan respon HTTP 401
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Menyimpan nilai UserID dari claims ke dalam context dengan key "id"
		ctx.Set("email", claimsData.Email)

		 // Melanjutkan request ke handler atau endpoint selanjutnya
		ctx.Next()
	})
}
