package piscine

func PodiumPosition(podium [][]string) [][]string {
	for a := 0; a < len(podium); a++ {
		for b := 0; b < len(podium)-1; b++ {
			if podium[b][0] > podium[b+1][0] {
				podium[b], podium[b+1] = podium[b+1], podium[b]
			}
		}
	}
	return podium
}
