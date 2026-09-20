package main
import (
	"github.com/atlas-devel/Go-Journey/users"
	"fmt"
)

func main (){
var newUser string= users.CreateUser("Atlas")
	fmt.Println(newUser)
}