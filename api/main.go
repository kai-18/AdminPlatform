package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// Lists of names and domains
var firstNames = []string{"Jaxon", "Lena", "Milo", "Aurora", "Zane", "Ivy", "Niko", "Daphne", "Theo", "Elara", "Finn", "Willow", "Asher", "Sienna", "Kai", "Luna", "Ronan", "Eloise", "Xavier", "Nova", "Levi", "Hazel", "Julian", "Celeste", "Ryder", "Sophie", "Knox", "Isla", "Beckett", "Emilia", "Archer", "Zara", "Caleb", "Freya", "Ezra", "Violet", "Gideon", "Maeve", "Silas", "Phoebe", "Cassian", "Dahlia", "Lucian", "Serena", "Orion", "Brielle", "Magnus", "Camille", "Atlas", "Amara"}
var lastNames = []string{"Fernandez", "Covington", "Serrano", "McAllister", "Holloway", "Calloway", "Harrington", "Vasquez", "Lancaster", "Fitzgerald", "Donovan", "Prescott", "Monroe", "Belmont", "Sterling", "Whitaker", "Delgado", "Chambers", "Hastings", "Kensington", "Morrison", "Sinclair", "Wellington", "Bennett", "Westwood", "Carmichael", "Davenport", "Ellington", "Fleming", "Goldman", "Hampton", "Iverson", "Jensen", "Kingston", "Langley", "Montgomery", "Nash", "Oakley", "Pembroke", "Quincy", "Remington", "Stratford", "Underwood", "Vanderbilt", "Winslow", "Xander", "York", "Ziegler"}
var emailDomains = []string{"gmail.com", "yahoo.com", "outlook.com", "icloud.com", "zoho.com", "protonmail.com", "aol.com", "mail.com", "gmx.com", "fastmail.com"}
var jobPositions = []string{ "CEO", "CTO", "CFO", "COO", "CIO", "CMO", "CHRO", "Director of Engineering", "Director of Marketing", "Director of Sales", "Director of Product", "Director of Operations", "IT Manager", "HR Manager", "Sales Manager", "Product Manager", "Project Manager", "Software Engineer", "Frontend Developer", "Backend Developer", "Full-Stack Developer", "Data Scientist", "DevOps Engineer", "Cloud Architect", "Cybersecurity Analyst", "Database Administrator", "IT Support Specialist", "UI/UX Designer", "Graphic Designer", "Content Writer", "SEO Specialist", "Social Media Manager", "Marketing Analyst", "Account Executive", "Sales Representative", "Business Development Manager", "Customer Support Specialist", "Accountant", "Financial Analyst", "HR Specialist", "Recruiter",
}
// Employee structure
type Employee struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Position string `json:"position"`
}

// Generate random employees
func generateUsers(count int) []Employee {
	rand.Seed(time.Now().UnixNano()) // Randomize seed

	users := []Employee{}
	for i := 1; i <= count; i++ {
		first := firstNames[rand.Intn(len(firstNames))]
		last := lastNames[rand.Intn(len(lastNames))]
		username := first + "." + last
		email := username + strconv.Itoa(i) + "@" + emailDomains[rand.Intn(len(emailDomains))]
		position:= jobPositions[rand.Intn(len(jobPositions))]

		users = append(users, Employee{
			Name:     first,
			Lastname: last,
			Username: username,
			Email:    email,
			Position: position,
		})
	}
	return users
}

// Middleware for CORS
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle OPTIONS request (Preflight)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// API handler
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(generateUsers(200))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", getUsers)

	// Wrap the mux with the CORS middleware
	handler := corsMiddleware(mux)

	println("🚀 Server running on port 8080")
	http.ListenAndServe(":8080", handler)
}
