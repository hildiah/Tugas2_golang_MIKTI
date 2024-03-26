package main

import "fmt"

func main() {
	var (
		teamScores = make(map[string][]int)
		teams      = []string{"Tim Lumba-lumba", "Tim Koala"}
	)

	// Input skor untuk setiap tim dan setiap pertandingan
	for _, team := range teams {
		fmt.Println("Masukkan skor untuk", team)
		for i := 1; i <= 3; i++ {
			var score int
			fmt.Print("Pertandingan", i, ": ")
			fmt.Scanln(&score)
			teamScores[team] = append(teamScores[team], score)
		}
	}

	// Hitung skor rata-rata untuk setiap tim
	averageScores := make(map[string]float64)
	for team, scores := range teamScores {
		sum := 0
		for _, score := range scores {
			sum += score
		}
		averageScores[team] = float64(sum) / float64(len(scores))
	}

	// Tentukan pemenang atau hasil seri
	if averageScores[teams[0]] > averageScores[teams[1]] {
		fmt.Println(teams[0], "memenangkan kompetisi dengan skor rata-rata", fmt.Sprintf("%.2f", averageScores[teams[0]]))
	} else if averageScores[teams[0]] < averageScores[teams[1]] {
		fmt.Println(teams[1], "memenangkan kompetisi dengan skor rata-rata", fmt.Sprintf("%.2f", averageScores[teams[1]]))
	} else {
		fmt.Println("Kompetisi berakhir seri dengan skor rata-rata", fmt.Sprintf("%.2f", averageScores[teams[0]]))
	}
}
