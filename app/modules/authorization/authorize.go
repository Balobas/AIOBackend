package authorization

/*
Авторизация пользователя
 */

//import (
//	"../../database"
//	"../users"
//	"encoding/json"
//	"github.com/dgrijalva/jwt-go"
//	"github.com/pkg/errors"
//	"../../../data"
//	"golang.org/x/crypto/bcrypt"
//	"os"
//	"../../models"
//)


//func Login(phone, password string) (string, error) {
//	_, err := users.ValidatePhoneByLocation(phone)
//	if err != nil {
//		return "", err
//	}
//	result, err := database.DATABASE.QueryAllFieldsWithSelector(`phone="` + phone + `"`)
//	if err != nil {
//		return "", err
//	}
//	if len(result) != 1 {
//		return "", errors.New("account is not registered")
//	}
//	var user models.User
//
//	b, err := json.Marshal(result)
//	if err != nil {
//		return "", err
//	}
//	err = json.Unmarshal(b, &user)
//	if err != nil {
//		return "", err
//	}
//	var pass models.UserPassword
//	res, err := database.DATABASE.Get("UserPass" + string(user.UID))
//	b, err = json.Marshal(res)
//	if err != nil {
//		return "", err
//	}
//	err = json.Unmarshal(b, &pass)
//	if err != nil {
//		return "", err
//	}
//	err = bcrypt.CompareHashAndPassword([]byte(pass.Password), []byte(password))
//	if err != nil {
//		return "", err
//	}
//	tk := &data.AccessToken{UserUid:pass.UserUID}
//	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
//	return token.SignedString([]byte(os.Getenv("token_password")))
//}