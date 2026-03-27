package particles

import (
	"math"
	"math/rand"
	"project-particles/config"
)

// spawnPosition renvoie une position (x,y) valide pour la création d'une
// particule. Si `config.General.SpawnRadius > 0` la position est choisie
// aléatoirement à l'intérieur du cercle centré en (SpawnX,SpawnY). Sinon,
// si `RandomSpawn` est vrai la position est choisie dans la fenêtre.
// Sinon la position fixe (SpawnX,SpawnY) est renvoyée.
func spawnPosition() (float64, float64) {
	var px, py float64
	if config.General.RandomSpawn == true {
		if config.General.SpawnRadius > 0 {
			r := math.Sqrt(rand.Float64()) * float64(config.General.SpawnRadius)
			theta := rand.Float64() * 2 * math.Pi
			px = float64(config.General.SpawnX) + r*math.Cos(theta)
			py = float64(config.General.SpawnY) + r*math.Sin(theta)
		} else {
			px = rand.Float64() * float64(config.General.WindowSizeX)
			py = rand.Float64() * float64(config.General.WindowSizeY)
		}
	} else {
		px = float64(config.General.SpawnX)
		py = float64(config.General.SpawnY)
	}
	return px, py
}
