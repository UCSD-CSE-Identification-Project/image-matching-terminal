// main.go

package main

//import "os"

func main() {
	a := App{}

	/*a.Initialize(
	os.Getenv("APP_DB_USERNAME"),
	os.Getenv("APP_DB_PASSWORD"),
	os.Getenv("APP_DB_NAME"))
	*/

	a.Initialize("test", "", "test_one")

	a.Run(":8080")





	defer a.DB.Close()

}
