package services

import (
	"fmt"

	"github.com/atlas-devel/Go-Journey/challenges/LMS/models"
)

func GetDetails(b models.Book) string {
	return fmt.Sprintf("Title: %s Author: %s Price: %f Copies: %d", b.Title,b.Author,b.Price,b.Copies)

}
func GetTotalValue(b models.Book)float64 {
	totalBooks:=b.Borrowed+b.Copies
	return float64(totalBooks)*b.Price
}