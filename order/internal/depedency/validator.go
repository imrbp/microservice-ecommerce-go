package dependency

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CustomValidator struct {
  Validate *validator.Validate
  Log *logrus.Logger
}

func NewCustomValidator(logger *logrus.Logger) *CustomValidator {
  return &CustomValidator{
    Validate : validator.New(),
    Log : logger,
  }
}

func (cV *CustomValidator) CreateErrorMessage(fieldError validator.FieldError) string {
  var sb strings.Builder
  
  sb.WriteString("Validation failed on field '"  + fieldError.Field() + "'" )
  sb.WriteString(", Condition: " + fieldError.ActualTag())


  if fieldError.Param() != "" {
		sb.WriteString(" { " + fieldError.Param() + " }")
	}

	if fieldError.Value() != nil && fieldError.Value() != "" {
		sb.WriteString(fmt.Sprintf(", actual: %v", fieldError.Value()))
	}

  return sb.String()
}

func (cV *CustomValidator) ValidateStruct(i interface{}) error {
  err := cV.Validate.Struct(i)
  
  errMessage := ""

  if _, ok := err.(*validator.InvalidValidationError) ; ok {
    cV.Log.WithFields(logrus.Fields{
      "Error" : err,
    }).Warn("Unkown Validation Struct Product")
  }

  if err != nil {
    for _,err := range err.(validator.ValidationErrors) {
      errMessage += cV.CreateErrorMessage(err) + "\n"
    }
    return err
  } else {
    return nil
  }
 }


func (cV *CustomValidator) ParseBody(ctx *fiber.Ctx, payload interface{}) error {
  if err := ctx.BodyParser(&payload) ; err != nil {
    return err
  }
  return nil
}
