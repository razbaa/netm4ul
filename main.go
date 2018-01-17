package main

//import des modules
import (
	"fmt"
	//"github.com/netm4ul/netm4ul/modules/recon"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

//main function
func main() {
	//print from main package
	fmt.Println("- NetM4ul -")

	//call recon module function
	//recon.Scan_output()
	fmt.Println("Database dev:")

	db, err := gorm.Open("postgres", "host=172.23.74.101 user=netm4ul dbname=netm4ul sslmode=disable password=netm4ul")

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect database.")
	}
	defer db.Close()

	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "Zgeg", Price: 1220})

	var product Product

	db.First(&product, 1)

	db.First(&product, "code = ?", "Zgeg")

	db.Model(&product).Update("Price", 1230)

}
