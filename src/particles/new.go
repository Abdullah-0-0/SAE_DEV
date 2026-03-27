package particles

import (
	"container/list"
	"math/rand" // pour générer des nombres aléatoires
	"project-particles/config"
)

// NewSystem est une fonction qui initialise un système de particules et le retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.

func NewSystem() System {

	l := list.New() // crée une nouvelle liste qui est vide

	// pour pouvoir crée autant de particules que défini dans InitNumParticles (dans config.json)
	for i := 0; i < config.General.InitNumParticles; i++ {

		// Position initiale
		var x, y float64

		//Situer dans config.json : "RandomSpawn": .., Cette ligne permet met de renvoyer un résultat selon le terme choisi true ou false (SpawnX, SpawnY)

		if config.General.RandomSpawn { // [TRUE] position aléatoire dans la fenêtre ici Chaque particule naît à une position aléatoire

			x = rand.Float64() * float64(config.General.WindowSizeX)
			y = rand.Float64() * float64(config.General.WindowSizeY)

		} else { // [FALSE] position fixe ici toutes les particules naissent au même point

			x = float64(config.General.SpawnX)
			y = float64(config.General.SpawnY)
		}

		// pour définir la vitesse pour chaque particule
		vx := rand.Float64()*4 - 3
		vy := rand.Float64()*4 - 3

		// Création de la particule
		p := &Particle{
			PositionX: x, PositionY: y, // position départ
			VelocityX: vx, VelocityY: vy, // vitesse
			ScaleX: 1, ScaleY: 1, // taille normale

			ColorRed:   rand.Float64(), // rouge aléatoire entre 0 et 1
			ColorGreen: rand.Float64(), // vert aléatoire
			ColorBlue:  rand.Float64(), // bleu aléatoire

			Opacity: 1, // opaque
		}

		l.PushBack(p) // ajoute la particule dans la liste
	}

	return System{Content: l} // retourne le système avec toutes les particules créées
}
