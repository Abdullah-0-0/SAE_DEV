package particles

import (
	"container/list"
	"math/rand"
	"project-particles/config"
	"project-particles/particle"
)

func NewSystem() System {
	// `l` contient la liste des particules créées
	l := list.New()
	// SpawnAccumulator initialisé à 0 (aucune fraction accumulée)
	sys := System{Content: l, SpawnAccumulator: 0}

	// Assigner le générateur via la fonction Generateur
	sys.Generateur = Generateur(&sys)
	return sys
}

// Generateur retourne une fonction qui génère des particules et les ajoute au système.
func Generateur(sys *System) func() {
	return func() {
		// Génère `part_gener` particules à la fois
		nb_par := config.General.PartGener
		if nb_par <= 0 {
			return
		}
		for i := 0; i < nb_par; i++ {
			var px, py float64
			if config.General.RandomSpawn == true {
				px = rand.Float64() * float64(config.General.WindowSizeX)
				py = rand.Float64() * float64(config.General.WindowSizeY)
				// pour le cercle
				//px, py = spawnPosition()
			} else {
				px = float64(config.General.SpawnX)
				py = float64(config.General.SpawnY)
				// pour le cercle
				//px, py = spawnPosition()
			}
			// vitesses aléatoires et durée de vie
			vx := (rand.Float64() - 0.5) * 10
			vy := (rand.Float64() - 0.5) * 10
			maxAge := 300 + rand.Intn(120)
			p := particle.NewParticle(px, py, 0, 1.5, 1.5, rand.Float64(), rand.Float64(), rand.Float64(), 1, vx, vy, 0, float64(maxAge))
			sys.Content.PushBack(&p)
		}
	}
}
