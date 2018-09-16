package main

import (
    "net/http"
    "fmt"
    "encoding/json"
    "log"
)

var jsonGrid = `[
   [0, 4, 8, 0, 1, 2, 6, 0, 5],
   [6, 3, 0, 0, 4, 0, 0, 0, 0],
   [7, 0, 0, 8, 0, 6, 3, 0, 0],
   [4, 0, 3, 5, 0, 0, 0, 0, 0],
   [0, 0, 6, 4, 0, 1, 5, 0, 0],
   [0, 0, 0, 0, 0, 9, 1, 0, 4],
   [0, 0, 9, 7, 0, 3, 0, 0, 1],
   [0, 0, 0, 0, 9, 0, 0, 5, 2],
   [1, 0, 4, 2, 6, 0, 7, 9, 0]
]`

func main() {
    routes()

    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRandomGrid() http.HandlerFunc {
    type response struct {
        OriginalGrid [][]int `json:"original_grid"`
        SolvedGrid [][]int `json:"solved_grid"`
    }
    arrayGrid = readGrid(jsonGrid)
    return func(w http.ResponseWriter, r *http.Request) {
        origGrid, SolvGrid := startSolving(jsonGrid)
        resp := &response{OriginalGrid: origGrid, SolvedGrid: SolvGrid}
        resp_json, _ := json.Marshal(resp)
        fmt.Fprintf(w, string(resp_json))
    }
}

var jsonGrid2 = `[
   [0, 0, 6, 8, 2, 9, 7, 1, 5],
   [5, 0, 7, 3, 4, 6, 9, 8, 2],
   [2, 8, 9, 1, 5, 7, 6, 3, 4],
   [3, 6, 4, 5, 9, 2, 8, 7, 1],
   [9, 7, 2, 6, 8, 1, 5, 4, 3],
   [1, 5, 8, 4, 7, 3, 2, 9, 6],
   [8, 2, 3, 7, 1, 5, 4, 6, 9],
   [7, 9, 1, 2, 6, 0, 3, 5, 8],
   [6, 4, 5, 9, 3, 8, 0, 2, 7]
]`

var jsonGrid3 = `[
   [0, 0, 6, 8, 2, 9, 7, 1, 5],
   [0, 0, 7, 3, 4, 6, 9, 8, 2],
   [0, 8, 9, 1, 5, 7, 6, 3, 4],
   [0, 6, 4, 5, 9, 2, 8, 7, 1],
   [0, 7, 2, 6, 8, 1, 5, 4, 3],
   [0, 5, 8, 4, 7, 3, 2, 9, 6],
   [0, 2, 3, 7, 1, 5, 4, 6, 9],
   [0, 9, 1, 2, 6, 0, 3, 5, 8],
   [0, 4, 5, 9, 3, 0, 0, 2, 7]
]`

/*
switch r.Method {
    case "GET":
        fmt.Fprintf(w, "Hello GET, %q", html.EscapeString(r.URL.Path))
    case "POST":
        fmt.Fprintf(w, "Hello POST, %q", html.EscapeString(r.URL.Path))
    default:
        fmt.Fprintf(w, "Not implemented.")
}*/