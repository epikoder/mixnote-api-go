package totem

import (
	"crypto/ecdsa"
	"os"
	"sync"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/mixnote/mixnote-api-go/configs"
	"github.com/mixnote/mixnote-api-go/src/core/models"
	"github.com/mixnote/mixnote-api-go/src/core/service/totem/guard"
	"github.com/mixnote/mixnote-api-go/src/framework/cache"
	"github.com/mixnote/mixnote-api-go/src/framework/utilities"
)

type (
	Totem struct{}

	TotemAttributes struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		AccessUuid   string `json:"-"`
		RefreshUuid  string `json:"-"`
		AtExpires    int64  `json:"-"`
		RtExpires    int64  `json:"-"`
	}
)

var (
	JWT_SECRET     string
	privKey        *ecdsa.PrivateKey
	useRedis       bool
	once           sync.Once
	err            error
	accessExpires  int64 = 20
	refreshExpires int64 = 10080
)

func NewTotem() (tt *Totem) {
	once.Do(func() {
		if JWT_SECRET = os.Getenv("JWT_SECRET"); JWT_SECRET == "" {
			utilities.Console.Fatal("JWT_SECRET not found")
		}
		privKey, err = jwt.ParseECPrivateKeyFromPEM([]byte(JWT_SECRET))
		if ok := utilities.Logger.HandleError(err); !ok {
			utilities.Console.Fatal(err)
		}
	})
	return
}

func (tt *Totem) CreateToken(user *models.User) (tA *TotemAttributes, err error) {
	tA = &TotemAttributes{}
	tA.AccessUuid = uuid.New().String()
	expires := time.Now().Add(time.Minute * time.Duration(accessExpires)).Unix()
	uAccessToken := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"aud": map[string]string{
			"id":         user.ID.String(),
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"photo":      user.Photo,
		},
		"sub": "access_token",
		"id":  tA.AccessUuid,
		"iss": configs.App.Name,
		"exp": expires,
		"iat": time.Now().Unix(),
	})
	if tA.AccessToken, err = uAccessToken.SignedString(privKey); err != nil {
		return nil, err
	}
	tA.AtExpires = expires

	tA.RefreshUuid = uuid.New().String()
	expires = time.Now().Add(time.Minute * time.Duration(refreshExpires)).Unix()
	uRefreshToken := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"id":  tA.RefreshUuid,
		"aid": tA.AccessUuid,
		"iss": configs.App.Name,
		"sub": "refresh_token",
		"exp": expires,
		"iat": time.Now().Unix(),
	})
	if tA.RefreshToken, err = uRefreshToken.SignedString(privKey); err != nil {
		return nil, err
	}
	tA.RtExpires = expires
	if useRedis {
		cache.UseDB(1)
		if err := cache.Set(tA.AccessUuid, user.ID.String(), tt.Duration(tA.AtExpires)).Err(); err != nil {
			return nil, err
		}
		if err := cache.Set(tA.RefreshUuid, user.ID.String(), tt.Duration(tA.RtExpires)).Err(); err != nil {
			return nil, err
		}
	}

	return tA, nil
}

func User() *models.User {
	return guard.User()
}

func (*Totem) Duration(unix int64) time.Duration {
	return time.Until(time.Unix(unix, 0))
}

func UseRedis() {
	useRedis = true
}

func AccessExpiresIn(minutes int64) {
	accessExpires = minutes
}

func RefreshExpiresIn(minutes int64) {
	refreshExpires = minutes
}

func DisbleRedis() {
	useRedis = false
}
