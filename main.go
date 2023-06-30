package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

const (
	DEFAULT_VOICE = "en-us-libritts-high.onnx"
	DEFAULT_PORT  = "8080"
	MODELS_DIR    = "models"
	PIPER_PATH    = "./bin/piper/piper"

	// DEFAULT_TEXT = "Welcome to the world of speech synthesis!"  // unused
)

func escapeString(input string) string {
	// escape single quotes, pipe, and backslash, and double quotes
	input = strings.Replace(input, "'", "\\'", -1)
	input = strings.Replace(input, "|", "\\|", -1)
	input = strings.Replace(input, "\\", "\\\\", -1)
	input = strings.Replace(input, "\"", "\\\"", -1)
	return input
}

func getListOfVoices() ([]string, error) {
	// read the list of files in the models directory
	cmd := exec.Command("ls", MODELS_DIR)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	stdoutBytes, err := io.ReadAll(stdoutPipe)
	if err != nil {
		return nil, err
	}

	err = cmd.Wait()
	if err != nil {
		return nil, err
	}

	voices := strings.Split(string(stdoutBytes), "\n")
	// filter out non .onnx files
	var filteredVoices []string
	for _, voice := range voices {
		if strings.HasSuffix(voice, ".onnx") {
			filteredVoices = append(filteredVoices, voice)
		}
	}

	return filteredVoices, nil
}

func logToTextFile(text string, voice string) {
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer f.Close()

	// timestamp, text, voice " CSV Style"
	timestamp := time.Now()
	_, err = f.WriteString(timestamp.Format("2006-01-02 15:04:05") + ", " + text + ", " + voice + "\n")
	if err != nil {
		fmt.Println("Error writing to file")
		return
	}
}

func runExecutable(input string, voice string) (io.Reader, error) {
	// fileName := hashString(input) + ".wav"
	esc := escapeString(input)
	logToTextFile(esc, voice)
	voice = MODELS_DIR + "/" + voice
	cmd := exec.Command(PIPER_PATH, "--model", voice, "--output_file", "-")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	_, err = io.WriteString(stdin, esc)
	if err != nil {
		return nil, err
	}
	stdin.Close()
	return stdoutPipe, nil
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	var jsonBody struct {
		Text  string `json:"text"`
		Voice string `json:"voice"`
	}

	err := json.NewDecoder(r.Body).Decode(&jsonBody)
	if err != nil {
		http.Error(w, "Error parsing json", http.StatusBadRequest)
		return
	}

	// trim whitespace
	jsonBody.Text = strings.TrimSpace(jsonBody.Text)
	if jsonBody.Text == "" {
		http.Error(w, "Error parsing json - text", http.StatusBadRequest)
		return
	}

	voice := DEFAULT_VOICE
	if jsonBody.Voice != "" {
		voice = jsonBody.Voice
	}

	defer r.Body.Close()
	w.Header().Set("Content-Type", "audio/wav")
	w.WriteHeader(http.StatusOK)

	if voice != DEFAULT_VOICE {
		voices, err := getListOfVoices()
		if err != nil {
			http.Error(w, "Error getting list of voices", http.StatusInternalServerError)
			return
		}

		var voiceFound bool
		for _, v := range voices {
			if v == voice {
				voiceFound = true
				break
			}
		}

		if !voiceFound {
			voice = DEFAULT_VOICE
		}
	}

	stdoutPipe, err := runExecutable(jsonBody.Text, voice)
	if err != nil {
		http.Error(w, "Error running executable", http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(w, stdoutPipe)
	if err != nil {
		http.Error(w, "Error streaming audio data", http.StatusInternalServerError)
		return
	}
}

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "audio/wav")
	w.WriteHeader(http.StatusOK)

	inputText := r.URL.Query().Get("text")
	// trim whitespace
	inputText = strings.TrimSpace(inputText)
	if inputText == "" {
		http.Error(w, "Missing Text Parameter.", http.StatusBadRequest)
	}
	voice := r.URL.Query().Get("voice")
	if voice == "" {
		voice = DEFAULT_VOICE
	}

	if voice != DEFAULT_VOICE {
		voices, err := getListOfVoices()
		if err != nil {
			http.Error(w, "Error getting list of voices", http.StatusInternalServerError)
			return
		}

		var voiceFound bool
		for _, v := range voices {
			if v == voice {
				voiceFound = true
				break
			}
		}

		if !voiceFound {
			voice = DEFAULT_VOICE
		}
	}

	stdoutPipe, err := runExecutable(inputText, voice)
	if err != nil {
		http.Error(w, "Error running executable", http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(w, stdoutPipe)
	if err != nil {
		http.Error(w, "Error streaming audio data", http.StatusInternalServerError)
		return
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/tts", handlePostRequest).Methods("POST")
	r.HandleFunc("/api/tts", handleGetRequest).Methods("GET")

	r.HandleFunc("/api/voices", func(w http.ResponseWriter, r *http.Request) {
		voices, err := getListOfVoices()
		if err != nil {
			http.Error(w, "Error getting list of voices", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(voices)
	}).Methods("GET")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.Handle("/", r)

	fmt.Println("Server listening on port " + DEFAULT_PORT)
	http.ListenAndServe(":"+DEFAULT_PORT, nil)
}
