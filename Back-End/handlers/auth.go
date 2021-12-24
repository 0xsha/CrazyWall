package handlers

import (
	"back-end/config"
	"back-end/database"
	"back-end/models"
	"github.com/form3tech-oss/jwt-go"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"time"
	"unicode"
)

//  handlers return error accept context

// checkHashedPassword check bcrypt hash against clear text password
func checkHashedPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}


// hashPassword generate bcrypt hash for clear text password
// output example $2a$14$P1TBmZBsF8siKWiGnariP.hctvsysaX8KhhK9NAXLxdIKD3BWJ4ka == secret
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password),14)
	return string(bytes), err
}


//// UpdateUser update user
//func UpdateUser(c *fiber.Ctx) error {
//	type UpdateUserInput struct {
//		Names string `json:"names"`
//	}
//	var uui UpdateUserInput
//	if err := c.BodyParser(&uui); err != nil {
//		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
//	}
//	id := c.Params("id")
//	token := c.Locals("user").(*jwt.Token)
//
//	if !validToken(token, id) {
//		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid token id", "data": nil})
//	}
//}


//// validateToken token with user id
//func validateToken(t *jwt.Token, id string) bool {
//	n, err := strconv.Atoi(id)
//	if err != nil {
//		return false
//	}
//
//	claims := t.Claims.(jwt.MapClaims)
//
//	// type assertion because claims is interface
//	uid := int(claims["user_id"].(float64))
//
//	if uid != n {
//		return false
//	}
//
//	return true
//}

func Login(c *fiber.Ctx) error  {

	// memory will be freed after function ;)
	type Authentication struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}



	auth := new(Authentication)
	c.Accepts("application/json") // "application/json"

	//log.Println(string(c.Body()))

	err := c.BodyParser(auth )
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}


	// actual credential authentication

	// get user from database
	user , err := models.GetUserByEmail(auth.Email , database.DBInstance)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}


	// only report invalid combination so no user enumeration
	if !checkHashedPassword(auth.Password,user.Password){

		//log.Println("huh")
		return c.SendStatus(fiber.StatusUnauthorized)

	}

	//if auth.Email != "me@0xsha.io" || auth.Password != "$2a$14$P1TBmZBsF8siKWiGnariP.hctvsysaX8KhhK9NAXLxdIKD3BWJ4ka"{
	//	return c.SendStatus(fiber.StatusUnauthorized)
	//}

	// generate a new token and set up claims
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix()

	// try to sign the token using app secret
	t, err := token.SignedString([]byte(config.GetKey("APP_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func Logout(c *fiber.Ctx) error  {
	c.ClearCookie()
	return c.JSON(fiber.Map{"msg" : "Stay safe :)"})

}

func User(c *fiber.Ctx) error  {

	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	return c.JSON(fiber.Map{"user": uid})
}

func Register(c *fiber.Ctx) error  {

	type Register struct {
	//	Data struct  {
			Username        string `json:"username" validate:"required,min=3,max=32,alphanum"`
			Email           string `json:"email" validate:"required,email"`
			Password        string `json:"password"`
			PasswordConfirm string `json:"password_confirm"`
			InviteCode string `json:"invite_code"`


			// captcha ?
			//} `json:"data" ,validate:"dive"`

	}

	//log.Println(string(c.Body()))
	reg := new(Register)
	c.Accepts("application/json") // "application/json"


	err := c.BodyParser(reg)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}



	if reg.InviteCode != "UDLx7TZld7LBvV9XFeu7QF4Sqsjrv0cz"{

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error" : "Invalid license key"})

	}

	// server side validation
	validate := validator.New()
	err = validate.Struct(reg)
	if err != nil{

		for _, err := range err.(validator.ValidationErrors) {

			if err.ActualTag() == "alphanum" && err.Field() == "Username" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error" : "Username should be alphanumeric"})
			}

			if err.ActualTag() == "min" && err.Field() == "Username" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error" : "Username should be at least 3 chars"})
			}

			if err.ActualTag() == "email" && err.Field() == "Email" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error" : "Invalid e-mail address"})
			}

		}

	}

	// custom validators
	if models.EmailExists(reg.Email , database.DBInstance){
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"Error" : "E-mail already exist"})
	}
	//
	// match passwords
	// also can validate.VarWithValue(password, confirmpassword, "eqfield")
	if reg.Password != reg.PasswordConfirm {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error" : "Password confirmation doesn't match password"})
	}
	//
	// check if password is strong enough
	if !checkPasswordStrength(reg.Password , 8) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error":"Password is weak min 8 chars one digit, one CAPITAL case , one special symbol"})
	}

	// hash the password using bcrypt
	hashedPassword , err := hashPassword(reg.Password)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var user models.User

	user.Email = reg.Email
	user.Name = reg.Username
	user.Password = hashedPassword

	err = models.CreateUser(user , database.DBInstance)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}


	return c.SendStatus(fiber.StatusCreated)


}



//checkPasswordStrength a very simple function to check password strength
func checkPasswordStrength(password string , length int)  bool {

	var lower,upper,symbol,digit bool

	// couldn't find this one on unique space
	isSpecial := func(c rune) bool {

		special :=  `!"#$%&'()*+,-./:;<=>?@[\]^_{|}~\`
		for _ , s := range special {
			if s == c {
				return true
			}
		}
		return false

	}

	// any length
	if len(password) > length {

		for _, c := range password {

			if unicode.IsUpper(c) {
				upper = true
			}

			if unicode.IsDigit(c) {
				digit = true
			}

			if unicode.IsLower(c) {
				lower = true
			}
			if isSpecial(c) {
				symbol = true
			}
		}

		// got all lower case , upper case and symbols
		if upper && lower && symbol && digit {
			return true
		}
	}
	return false
}
