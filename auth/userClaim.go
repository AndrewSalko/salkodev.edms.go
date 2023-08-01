package auth

import jwt "github.com/golang-jwt/jwt/v5"

// UserClaim.Flags - флаг означає що це токен для підтвердження реєстрації користувача
// цей токен не можна використовувати ніде в інших функціях
const FlagRegistrationConfirmation = 1

type UserClaim struct {
	jwt.RegisteredClaims
	Email string
	Flags uint //позначки-флаги для окремих типів jwt
}