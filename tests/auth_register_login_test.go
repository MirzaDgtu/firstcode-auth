package tests

import (
	"firstcode-auth/tests/suite"
	"testing"
	"time"

	ssov1 "github.com/MirzaDgtu/firstcode_protos/gen/go/sso"
	"github.com/golang-jwt/jwt"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	appID          = 1
	appSecret      = "test-secret"
	passDefaultLen = 10
)

func TestRegisterLogin_Login_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	email := gofakeit.Email()
	pass := randomFakePassword()
	firstname := gofakeit.FirstName()
	name := gofakeit.Name()
	lastname := gofakeit.LastName()
	phone := gofakeit.Phone()
	sex := ""

	respReg, err := st.AuthClient.Register(ctx, &ssov1.RegisterRequest{
		Email:     email,
		Password:  pass,
		FirstName: firstname,
		Name:      name,
		LastName:  lastname,
		Phone:     phone,
		Sex:       sex,
	})

	require.NoError(t, err)
	assert.NotEmpty(t, respReg.GetUserId())

	respLogin, err := st.AuthClient.Login(ctx, &ssov1.LoginRequest{
		Email:    email,
		Password: pass,
		AppId:    appID,
	})
	require.NoError(t, err)

	token := respLogin.GetToken()
	require.NotEmpty(t, token) // Проверяем, что он не пустой

	// Отмечаем время, в которое бы выполнен логин.
	// Это понадобится для проверки TTL токена
	loginTime := time.Now()

	// Парсим и валидируем токен
	tokenParsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(appSecret), nil
	})
	// Если ключ окажется невалидным, мы получим соответствующую ошибку
	require.NoError(t, err)

	// Преобразуем к типу jwt.MapClaims, в котором мы сохраняли данные
	claims, ok := tokenParsed.Claims.(jwt.MapClaims)
	require.True(t, ok)

	// Проверяем содержимое токена
	assert.Equal(t, respReg.GetUserId(), int64(claims["uid"].(float64)))
	assert.Equal(t, email, claims["email"].(string))
	assert.Equal(t, appID, int(claims["app_id"].(float64)))

	const deltaSeconds = 1

	// Проверяем, что TTL токена примерно соответствует нашим ожиданиям.
	assert.InDelta(t, loginTime.Add(st.Cfg.TokenTTL).Unix(), claims["exp"].(float64), deltaSeconds)
}

func randomFakePassword() string {
	return gofakeit.Password(true, true, true, true, false, passDefaultLen)
}
