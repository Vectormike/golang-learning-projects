package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter email address: ")

	for scanner.Scan() {
		email := scanner.Text()
		if checkEmail(email) {
			fmt.Println("Email address is valid")
		} else {
			fmt.Println("Email address is invalid")
		}
		fmt.Println("Enter email address: ")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading input: ", err)
	}
}

func checkEmail(email string) bool {
	if len(email) < 3 && len(email) > 254 {
		return false
	}

	if strings.Count(email, "@") != 1 {
		return false
	}

	parts := strings.Split(email, "@")
	username := parts[0]
	domain := parts[1]

	if len(username) < 1 {
		return false
	}

	if len(domain) < 3 && len(domain) > 253 {
		return false
	}

	if strings.HasPrefix(domain, ".") || strings.HasSuffix(domain, ".") {
		return false
	}

	if strings.Contains(domain, "..") {
		return false
	}

	if !strings.Contains(domain, ".") {
		return false
	}

	if strings.Count(domain, ".") > 1 {
		return false
	}

	if strings.Count(domain, ".") == 1 {
		parts := strings.Split(domain, ".")
		topLevelDomain := parts[1]

		if len(topLevelDomain) < 2 || len(topLevelDomain) > 6 {
			return false
		}
	}

	if net.ParseIP(domain) != nil {
		return false
	}

	if strings.Contains(username, "..") {
		return false
	}

	mxs, err := net.LookupMX(domain)
	if err != nil || len(mxs) == 0 {
		return false
	}

	return true
}
