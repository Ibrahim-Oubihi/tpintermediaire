package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ----------------------------------------------------------------

func timeibra(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		day := time.Now()
		fmt.Fprintf(w, "%sh%s\n", (day.Hour()), (day.Minute()))
		return
	default:
		fmt.Fprintf(w, "Method not allowed")
		return
	}
}

// ----------------------------------------------------------------

func diceibra(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:

		rand.Seed(time.Now().UnixNano())
		fmt.Fprintf(w, "%d", rand.Intn(1000)+1)
		return
	default:
		fmt.Fprintf(w, "Method not allowed")
		return
	}
}

// ----------------------------------------------------------------

func dicesibra(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:

		if req.URL.Query().Get("type") != "" {
			if req.URL.Query().Get("type")[0] == 'd' {

				sides, _ := strconv.Atoi(req.URL.Query().Get("type")[1:])

				for i := 0; i < 15; i++ {
					fmt.Fprintf(w, "%d ", rand.Intn(sides)+1)
				}
				return

			} else {
				fmt.Fprintf(w, "Bad request")
				return
			}
		} else {

			cote := []int{2, 4, 6, 8, 10, 12, 20, 100}

			for i := 0; i < 15; i++ {
				sides := cote[rand.Intn(len(cote))]
				fmt.Fprintf(w, "%d ", rand.Intn(sides)+1)
			}
			return
		}

		return
	default:
		fmt.Fprintf(w, "Method not allowed")
		return
	}
}

// ----------------------------------------------------------------

func randomize_wordsibra(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:

		rand.Seed(time.Now().UnixNano())

		d_words := req.FormValue("words")

		words := strings.Split(d_words, " ")

		rand.Shuffle(len(words), func(i, j int) { words[i], words[j] = words[j], words[i] })

		d_words = strings.Join(words, " ")

		fmt.Fprintf(w, "%s", d_words)
		return

	default:
		fmt.Fprintf(w, "Method not allowed")
		return
	}
}

// ----------------------------------------------------------------

func capitalize_d_wordsibra(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:

		return
	default:
		fmt.Fprintf(w, "Method not allowed")
		return
	}
}

// ----------------------------------------------------------------

func main() {
	http.HandleFunc("/", timeibra)
	http.HandleFunc("/ibrahim/dice", diceibra)
	http.HandleFunc("/ibrahim/dices", dicesibra)
	http.HandleFunc("/dices?type=", dicesibra)
	http.HandleFunc("/ibrahim/randomize-words", randomize_wordsibra)
	http.HandleFunc("/ibrahim/semi-capitalize-d_words", capitalize_d_wordsibra)
	http.ListenAndServe(":4567", nil)
}
