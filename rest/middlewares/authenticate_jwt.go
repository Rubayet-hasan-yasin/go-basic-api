package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecommerce/config"
	"ecommerce/util"
	"net/http"
	"strings"
)

func AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			util.SendError(w, http.StatusUnauthorized, "Missing Auth Token")
			return
		}
		headerArr := strings.Split(header, " ")

		if len(headerArr) != 2 || headerArr[0] != "Bearer" {
			util.SendError(w, http.StatusUnauthorized, "Invalid Auth Token")
			return
		}

		accessToken := headerArr[1]

		tokenParts := strings.Split(accessToken, ".")
		if len(tokenParts) != 3 {
			util.SendError(w, http.StatusUnauthorized, "Invalid Auth Token")
			return
		}

		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		jwtSignature := tokenParts[2]

		cnf := config.GetConfig()

		message := jwtHeader + "." + jwtPayload

		byteArrSecret := []byte(cnf.JwtSecretKey)
		byteArrMessage := []byte(message)

		h := hmac.New(sha256.New, byteArrSecret)
		h.Write(byteArrMessage)

		hash := h.Sum(nil)
		newSignatureB64 := util.Base64UrlEncode(hash)

		if newSignatureB64 != jwtSignature {
			util.SendError(w, http.StatusUnauthorized, "Invalid Auth Token")
			return
		}

		next.ServeHTTP(w, r)
	})
}
