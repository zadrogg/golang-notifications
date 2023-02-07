package notifications

import "github.com/joho/godotenv"

func init() {
	//load values from .env
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}

// @title           Документация сервиса уведомлений
// @version         1.0
// @description     Сервис уведомлений

// @contact.name   API Support

// @host      localhost
// @BasePath  /
// @schemes http
func main() {

}
