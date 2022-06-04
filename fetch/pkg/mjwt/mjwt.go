package mjwt

import (
	"errors"
	"log"

	"github.com/golang-jwt/jwt/v4"
)

const (
	CLAIMS       = "claims"
	jtiKey       = "jti"
	subKey       = "sub"
	nameKey      = "name"
	roleKey      = "role"
	phoneKey     = "phone"
	timestampKey = "timestamp"
	tokenTypeKey = "type"
	expKey       = "exp"
	freshKey     = "fresh"
)

var (
	ErrCastingClaims = errors.New("fail to type casting")
	ErrInvalidToken  = errors.New("token not valid")
)

type TokenReader interface {
	ValidateToken(tokenString string) (*jwt.Token, error)
	ReadToken(token *jwt.Token) (CustomClaim, error)
}

func New(secretKey string) TokenReader {
	if secretKey == "" {
		log.Fatal("secret key cannot be empty")
	}
	newCore := &core{
		secretKey: []byte(secretKey),
	}
	return newCore
}

type core struct {
	secretKey []byte
}

// ReadToken membaca inputan token dan menghasilkan pointer struct CustomClaim
// struct CustomClaim digunakan untuk nilai passing antar middleware
func (j *core) ReadToken(token *jwt.Token) (CustomClaim, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return CustomClaim{}, ErrCastingClaims
	}

	uniqueID, ok := claims[jtiKey].(string)
	if !ok {
		return CustomClaim{}, ErrCastingClaims
	}
	subject, ok := claims[subKey].(string)
	if !ok {
		return CustomClaim{}, ErrCastingClaims
	}
	name, ok := claims[nameKey].(string)
	if !ok {
		return CustomClaim{}, ErrCastingClaims
	}
	role, ok := claims[roleKey].(string)
	if !ok {
		return CustomClaim{}, ErrCastingClaims
	}
	phone, ok := claims[phoneKey].(string)
	if !ok {
		return CustomClaim{}, ErrCastingClaims
	}
	exp, ok := claims[expKey].(float64)
	if !ok {
		return CustomClaim{}, ErrCastingClaims
	}
	tokenType, ok := claims[tokenTypeKey].(string)
	if !ok {
		return CustomClaim{}, ErrCastingClaims
	}
	fresh, ok := claims[freshKey].(bool)
	if !ok {
		return CustomClaim{}, ErrCastingClaims
	}
	timestamp, ok := claims[timestampKey].(string)
	if !ok {
		return CustomClaim{}, ErrCastingClaims
	}

	customClaim := CustomClaim{
		UniqueID:  uniqueID,
		Sub:       subject,
		Phone:     phone,
		Name:      name,
		Role:      role,
		Timestamp: timestamp,
		Exp:       int64(exp),
		Type:      tokenType,
		Fresh:     fresh,
	}

	return customClaim, nil
}

// ValidateToken memvalidasi apakah token string masukan valid, termasuk memvalidasi apabila field exp nya kadaluarsa
func (j *core) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return j.secretKey, nil
	})

	// Jika expired akan muncul disini asalkan ada claims exp
	if err != nil {
		return nil, ErrInvalidToken
	}

	return token, nil
}
