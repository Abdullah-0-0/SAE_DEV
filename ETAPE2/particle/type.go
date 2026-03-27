package particle

import "container/list"

// Package particle : types et constructeurs pour une particule isolée.
// Ce paquet contient la définition de la particule (aspect + état) et
// un type System léger pour stocker une liste de particules.

// System définit un système de particules.
// Content : liste chaînée de pointeurs vers des particules.
// SpawnAccumulator : accumulateur flottant utilisé pour gérer les taux d'apparition.
type System struct {
	Content          *list.List // Liste chaînée contenant les particules (chaque élément est *particle.Particle)
	SpawnAccumulator float64    // Accumulateur utilisé pour gérer le spawn progressif (fraction de particule)
}

// Particle définit une particule.
// Les champs décrivent l'apparence (position, échelle, couleur, opacité ...)
// ainsi que l'état dynamique (vitesse, âge, durée de vie maximale).
type Particle struct {
	// Apparence / position
	PositionX, PositionY            float64 // PositionX/PositionY : coordonnées (pixels)
	Rotation                        float64 // Rotation : orientation (radians)
	ScaleX, ScaleY                  float64 // ScaleX/ScaleY : facteur d'échelle sur chaque axe
	ColorRed, ColorGreen, ColorBlue float64 // Color* : composantes de couleur normalisées (0..1)
	Opacity                         float64 // Opacity : transparence (1 = opaque, 0 = transparent)

	// État dynamique
	VelocityX, VelocityY float64 // VelocityX/VelocityY : vitesse en pixels par frame
	Age                  int     // Age : nombre de frames écoulées depuis la naissance
	MaxAge               int     // MaxAge : durée de vie maximale en frames (0 = illimité)
}
