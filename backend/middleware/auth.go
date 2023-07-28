package middleware

import (
	"AELS/ahttp"
	"AELS/persistence"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

/*
Checks 'Auth' Cookie.

No Cookie or Expired Token --> 401

Other Error --> 500

Valid Signature --> Checks User Still Exists

Signature Valid && User Exists --> Sets userid in request Context and
refreshes Auth cookie with new JWT to expire in 24 hours
*/
func Auth(next ahttp.Handler) ahttp.Handler {
	return ahttp.Handler(func(w http.ResponseWriter, r *http.Request) (int, error) {
		tokenCookie, err := r.Cookie("Auth")
		if errors.Is(err, http.ErrNoCookie) {
			return 401, errors.New("you are not logged in")
		} else if err != nil {
			return 500, err
		}

		token, err := jwt.ParseWithClaims(tokenCookie.Value, &JwtClaims{}, jwtKeyFunc)
		if errors.Is(err, jwt.ErrTokenExpired) {
			return 401, errors.New("your session has expired")
		} else if err != nil {
			fmt.Printf("JWT Validation Failed Failed: %s", err.Error())
			return 500, err
		}

		userid := token.Claims.(*JwtClaims).UserId

		var user persistence.User
		if err := persistence.DB.First(&user, userid).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			err := fmt.Sprintf("User id(%d) attempted access that does not exist", userid)
			fmt.Println(err)
			return 500, errors.New(err)
		}

		if err := AssignNewCookie(userid, w); err != nil {
			return 500, err
		}

		ctx := context.WithValue(r.Context(), UserID{}, userid)

		next.ServeHTTP(w, r.WithContext(ctx))

		return 0, nil
	})
}

/*
JWT Body Definition
*/
type JwtClaims struct {
	UserId uint64 `json:"userid"`
	jwt.RegisteredClaims
}

/*
Type for retrieved UserId from context
*/
type UserID struct{}

/*
Returns Signing Key if signing method is HMAC
*/
func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return []byte(os.Getenv("JWT_SECRET")), nil
}

/*
Creates new JWT to expire in 24hrs and assigns it to the 'Auth' cookie
*/
func AssignNewCookie(userid uint64, w http.ResponseWriter) error {
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS512, JwtClaims{
		userid,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	})

	newSignedToken, err := newToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return err
	}

	newAuthCookie := http.Cookie{
		Name:     "Auth",
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24),
		Value:    newSignedToken,
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
	}
	http.SetCookie(w, &newAuthCookie)

	return nil
}
