package algolia

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/joho/godotenv"
)

func init() {
	var err error
	if getEnv() == "development" {
		err = godotenv.Load()
	}

	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}
}

func getEnv() string {
	return os.Getenv("APP_ENV")
}

//InitAlgolia creates algolia connection
func InitAlgolia() {
	client := search.NewClient(os.Getenv("ALGOLIA_APP_ID"), os.Getenv("ALGOLIA_ADMIN_KEY"))
	// client := search.NewClient(Sprintf(("%s","%s"),os.Getenv("ALGOLIA_APP_ID"),os.Getenv("ALGOLIA_ADMIN_KEY")))
	index := client.InitIndex("demo_ecommerce")

	getRes, _ := http.Get("https://alg.li/doc-ecommerce.json")
	defer getRes.Body.Close()

	var products []map[string]interface{}
	err := json.NewDecoder(getRes.Body).Decode(&products)
	// error handling
	if err != nil {
		log.Fatalf("Error setting up algolia, %v", err)
	} else {
		log.Println("Setting up algolia")
	}
	_, err = index.SaveObjects(products, opt.AutoGenerateObjectIDIfNotExist(true))
	if err != nil {
		log.Fatalf("Error setting up algolia, %v", err)
	} else {
		log.Println("Setting up algolia")
	}

	_, err = index.SetSettings(search.Settings{
		// Select the attributes you want to search in
		SearchableAttributes: opt.SearchableAttributes(
			"brand", "name", "categories", "description",
		),
		// Define business metrics for ranking and sorting
		CustomRanking: opt.CustomRanking(
			"desc(popularity)",
		),
		// Set up some attributes to filter results on
		AttributesForFaceting: opt.AttributesForFaceting(
			"categories", "searchable(brand)", "price",
		),
	})
	if err != nil {
		log.Fatalf("Error setting up algolia, %v", err)
	} else {
		log.Println("Setting up algolia")
	}

} // error handling
// type Contact struct {
// 	Firstname string
// 	Lastname  string
// 	Company   string
// 	// Other fields omitted
//   }

// var contacts []Contact
// content, err := ioutil.ReadFile("contacts.json")
// err = json.Unmarshal(content, &contacts)

// index := client.InitIndex("contacts")
// res, err := index.SaveObjects(contacts, opt.AutoGenerateObjectIDIfNotExist(true))

// settings := search.Settings{
// 	CustomRanking: opt.CustomRanking("desc(followers)"),
//   }

//   res, err := index.SetSettings(settings)

//   settings := search.Settings{
//     SearchableAttributes: opt.SearchableAttributes(
//         "firstname",
//         "lastname",
//         "company",
//         "email",
//         "city",
//         "address",
//     ),
// }

// res, err := index.SetSettings(settings)

// // Search for a first name
// res, err := index.Search("jimmie")

// // Search for a first name with typo
// res, err := index.Search("jimie")

// // Search for a company
// res, err := index.Search("california paint")

// // Search for a first name and a company
// res, err := index.Search("jimmie paint")
