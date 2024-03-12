package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const (
	baseURL         = "https://api.lovense.com/api/lan/"
	getToysEndpoint = "getToys"
	controlEndpoint = "command"
	accessCode      = "your_access_code_here"
)

type Toy struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ControlRequest struct {
	ToyID     string `json:"toy_id"`
	Command   string `json:"command"`
	Strength  int    `json:"strength"`
	Duration  int    `json:"duration"`
	Loop      bool   `json:"loop"`
	Rotation  int    `json:"rotation"`
	Pump      int    `json:"pump"`
	AirLevel  int    `json:"air_level"`
	Vibration int    `json:"vibration"`
}

func getConnectedToys() ([]Toy, error) {
	url := baseURL + getToysEndpoint

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var toys []Toy
	err = json.NewDecoder(resp.Body).Decode(&toys)
	if err != nil {
		return nil, err
	}

	return toys, nil
}

func controlToy(toyID, command string, strength, duration, rotation, pump, airLevel, vibration int, loop bool) error {
	url := baseURL + controlEndpoint

	data := map[string]interface{}{
		"token":   accessCode,
		"uid":     toyID,
		"command": command,
		"timeSec": duration,
		"apiVer":  1,
	}

	switch command {
	case "Vibrate":
		data["strength"] = strength
		data["vibration"] = vibration
	case "Rotate":
		data["strength"] = strength
	case "Pump":
		data["status"] = pump
	case "AirIn", "AirOut":
		data["level"] = airLevel
	case "RotateChange":
		data["rotation"] = rotation
	}

	if loop {
		data["loopRun"] = 1
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", strings.NewReader(string(jsonData)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send command")
	}

	return nil
}

func getToysHandler(w http.ResponseWriter, r *http.Request) {
	toys, err := getConnectedToys()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(toys)
}

func controlHandler(w http.ResponseWriter, r *http.Request) {
	var req ControlRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = controlToy(req.ToyID, req.Command, req.Strength, req.Duration, req.Rotation, req.Pump, req.AirLevel, req.Vibration, req.Loop)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Command sent successfully")
}

func main() {
	localServer := flag.Bool("local", true, "Enable local server mode")
	flag.Parse()

	if *localServer {
		http.HandleFunc("/toys", getToysHandler)
		http.HandleFunc("/control", controlHandler)

		log.Println("Local server running on http://localhost:8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	} else {
		toys, err := getConnectedToys()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Connected Toys:")
		for _, toy := range toys {
			fmt.Printf("ID: %s, Name: %s\n", toy.ID, toy.Name)
		}

		var toyID, command string
		var strength, duration, rotation, pump, airLevel, vibration int
		var loop bool

		fmt.Print("Enter Toy ID: ")
		fmt.Scanln(&toyID)

		fmt.Print("Enter Command (Vibrate, Rotate, RotateChange, Pump, AirIn, AirOut, Stop): ")
		fmt.Scanln(&command)

		switch command {
		case "Vibrate":
			fmt.Print("Enter Strength (0-20): ")
			fmt.Scanln(&strength)
			fmt.Print("Enter Vibration Pattern (0-3): ")
			fmt.Scanln(&vibration)
		case "Rotate":
			fmt.Print("Enter Strength (0-20): ")
			fmt.Scanln(&strength)
		case "RotateChange":
			fmt.Print("Enter Rotation (0-3): ")
			fmt.Scanln(&rotation)
		case "Pump":
			fmt.Print("Enter Pump Status (0-3): ")
			fmt.Scanln(&pump)
		case "AirIn", "AirOut":
			fmt.Print("Enter Air Level (0-3): ")
			fmt.Scanln(&airLevel)
		}

		fmt.Print("Enter Duration (in seconds): ")
		fmt.Scanln(&duration)

		fmt.Print("Loop Command? (true/false): ")
		fmt.Scanln(&loop)

		err = controlToy(toyID, command, strength, duration, rotation, pump, airLevel, vibration, loop)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Command sent successfully")
	}
}
