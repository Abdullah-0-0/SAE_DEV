package particles

import "container/list"

// System définit un système de particules.
// Le contenu est une liste de pointeurs vers des particules définies dans
// le package `particle`.
type System struct {
	Content          *list.List // Liste chaînée contenant les particules (chaque élément est *particle.Particle)
	SpawnAccumulator float64    // Accumulateur fractionnel pour gérer un taux d'apparition non entier
	// Generateur : fonction appelée pour créer/ajouter une particule au système.
	Generateur func()
}
