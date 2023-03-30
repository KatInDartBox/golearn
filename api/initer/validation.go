package initer

import (
	"log"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

func AddBookableDate() {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := v.RegisterValidation("bookabledate", bookableDate)
		if err != nil {
			log.Fatal("can not add bookable validation")
		}
	}
	if !ok {
		log.Fatal("can not add bookable validation")
	}
}
