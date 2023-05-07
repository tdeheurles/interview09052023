package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

type User struct {
	Name  string
	Email string
	Age   int
}

var users = map[string]User{
	"alice": {"Alice", "alice@example.com", 30},
	"bob":   {"Bob", "bob@example.com", 35},
	"carol": {"Carol", "carol@example.com", 28},
}

func getUserInfo(username string) (User, bool) {
	user, exists := users[username]
	return user, exists
}

func main() {
	println("Configuring server ...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bonjour, ceci est un serveur web en Golang!")
	})

	http.HandleFunc("/env", func(w http.ResponseWriter, r *http.Request) {
		envVars := os.Environ()
		envVarsString := strings.Join(envVars, "\n")
		fmt.Fprintf(w, "Liste des variables d'environnement:\n%s", envVarsString)
	})

	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		username := strings.TrimPrefix(r.URL.Path, "/user/")
		user, exists := getUserInfo(username)
		if !exists {
			http.Error(w, "Utilisateur non trouvé", http.StatusNotFound)
			return
		}

		fmt.Fprintf(w, "Nom: %s\nEmail: %s\nAge: %d\n", user.Name, user.Email, user.Age)
	})

	println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}
