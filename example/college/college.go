package college

import "fmt"

// ✅ College Struct
type College struct {
	Name      string
	Location  string
	Students  int
	CollegeID string
}

// ✅ Function to Create College Data + Append New College
func CreateCollegeData() []College {

	colleges := []College{
		{Name: "IIT Delhi", Location: "Delhi", Students: 12000, CollegeID: "C101"},
		{Name: "NIT Patna", Location: "Bihar", Students: 8000, CollegeID: "C102"},
		{Name: "DU University", Location: "Delhi", Students: 15000, CollegeID: "C103"},
	}

	// ✅ Append New College Data
	colleges = append(colleges, College{
		Name:      "BHU Varanasi",
		Location:  "Uttar Pradesh",
		Students:  20000,
		CollegeID: "C104",
	})

	return colleges
}

// ✅ Function to Display Colleges
func DisplayColleges(colleges []College) {

	for i := 0; i < len(colleges); i++ {

		fmt.Println("College", i+1)
		fmt.Println("Name      :", colleges[i].Name)
		fmt.Println("Location  :", colleges[i].Location)
		fmt.Println("Students  :", colleges[i].Students)
		fmt.Println("CollegeID :", colleges[i].CollegeID)
		fmt.Println("--------------------------")
	}
}



