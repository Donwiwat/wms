package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"wms-backend/internal/config"
	"wms-backend/internal/database"
	"wms-backend/internal/models"
	"wms-backend/internal/repositories"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	// Parse command line flags
	godotenv.Load(".env")
	simple := flag.Bool("simple", false, "Use simple mode (password visible)")
	flag.Parse()

	if *simple {
		createUserSimple()
	} else {
		log.Fatal("Please select simple")
	}
}

func createUserSimple() {
	fmt.Println("WMS Superuser Creation Tool (Simple Mode)")
	fmt.Println("=========================================")

	cfg := config.Load()

	db, err := database.Initialize(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	userRepo := repositories.NewUserRepository(db)
	reader := bufio.NewReader(os.Stdin)

	username := readInput(reader, "Enter username: ")
	if username == "" {
		log.Fatal("Username cannot be empty")
	}

	// Check duplicate username
	if user, err := userRepo.GetByUsername(username); err == nil && user != nil {
		log.Fatalf("Username '%s' already exists", username)
	}

	email := readInput(reader, "Enter email: ")
	if email == "" {
		log.Fatal("Email cannot be empty")
	}

	if user, err := userRepo.GetByEmail(email); err == nil && user != nil {
		log.Fatalf("Email '%s' already exists", email)
	}

	password := readPassword("Enter password: ")
	if len(password) < 6 {
		log.Fatal("Password must be at least 6 characters long")
	}

	createUser(userRepo, username, email, password)
}

func readInput(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	value, _ := reader.ReadString('\n')
	return strings.TrimSpace(value)
}

func readPassword(prompt string) string {
	fmt.Print(prompt)
	pwdBytes, _ := terminal.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	return strings.TrimSpace(string(pwdBytes))
}

func createUser(userRepo repositories.UserRepository, username, email, password string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	user := &models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
		Role:     "admin",
		IsActive: true,
	}

	if err := userRepo.Create(user); err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}

	fmt.Printf("\nSuperuser '%s' created successfully!\n", username)
	fmt.Printf("User ID: %d\n", user.ID)
	fmt.Printf("Email: %s\n", user.Email)
	fmt.Printf("Role: %s\n", user.Role)
	fmt.Printf("Created at: %s\n", user.CreatedAt.Format("2006-01-02 15:04:05"))
}
