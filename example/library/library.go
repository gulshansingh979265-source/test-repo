package library

import "fmt"

// ✅ Library Struct
type Library struct {
	Name       string
	City       string
	Books      int
	LibraryID  string
}

// ✅ Function to Create Library Data + Append New Library
func CreateLibraryData() []Library {

	libraries := []Library{
		{Name: "Central Library", City: "Delhi", Books: 50000, LibraryID: "L101"},
		{Name: "State Library", City: "Patna", Books: 30000, LibraryID: "L102"},
		{Name: "City Library", City: "Mumbai", Books: 45000, LibraryID: "L103"},
	}

	// ✅ Append New Library Data
	libraries = append(libraries, Library{
		Name:      "University Library",
		City:      "Varanasi",
		Books:     60000,
		LibraryID: "L104",
	})

	return libraries
}

// ✅ Function to Display Libraries
func DisplayLibraries(libraries []Library) {

	for i := 0; i < len(libraries); i++ {

		fmt.Println("Library", i+1)
		fmt.Println("Name      :", libraries[i].Name)
		fmt.Println("City      :", libraries[i].City)
		fmt.Println("Books     :", libraries[i].Books)
		fmt.Println("LibraryID :", libraries[i].LibraryID)
		fmt.Println("--------------------------")
	}
}
