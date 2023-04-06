package password

import "golang.org/x/crypto/bcrypt"

func getPass(pass string) string {
	return "}#$0{" + pass + "&@01"
}

func HashPassword(secrete string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(getPass(secrete)), bcrypt.DefaultCost)
	return string(bytes), err
}

func IsMatch(secrete string, hash string) bool {
	pass := getPass(secrete)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))

	return err == nil
}
