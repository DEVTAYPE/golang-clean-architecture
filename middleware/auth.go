package middleware

import (
	"api-basico-dev/config"
	"api-basico-dev/handlers"
	"api-basico-dev/server"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next server.HandlerFunc) server.HandlerFunc {
	return func(ctx *server.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")

		if strings.TrimSpace(authHeader) == "" {
			handlers.RespondError(ctx, handlers.NewAppError("Token de autenticação ausente", http.StatusUnauthorized))
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			handlers.RespondError(ctx, handlers.NewAppError("Formato de token inválido", http.StatusUnauthorized))
			return
		}

		// validar token
		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			// El token string se ha parseado correctamente, pero debemos verificar que el método de firma sea el esperado
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, handlers.NewAppError("Método de assinatura inesperado", http.StatusUnauthorized)
			}

			return []byte(config.AppConfig.JWT_SECRET), nil
		})

		if err != nil || !token.Valid {
			handlers.RespondError(ctx, handlers.NewAppError("Token inválido", http.StatusUnauthorized))
			return
		}

		// recuperamos claims del token (los claims son los datos que se guardan dentro del token, como el user_id y la fecha de expiración)
		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			handlers.RespondError(ctx, handlers.NewAppError("Claims inválidos", http.StatusUnauthorized))
			return
		}

		userID, ok := claims["user_id"].(float64) // los números en JSON se parsean como float64

		if !ok {
			handlers.RespondError(ctx, handlers.NewAppError("User ID inválido en claims", http.StatusUnauthorized))
			return
		}

		ctx.SetUserUID(uint(userID))

		// si todo es correcto, llamamos al siguiente handler en la cadena de middlewares o al handler final
		next(ctx)
	}
}
