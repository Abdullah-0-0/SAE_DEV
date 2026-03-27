package main

import (
	"project-particles/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Update se charge d'appeler la fonction Update du système de particules
// g.system. Elle est appelée automatiquement exactement 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction ne devrait pas être modifiée sauf
// pour les deux dernières extensions.
func (g *game) Update() error {

	g.system.Update()

	// Gestion de la sélection et modification des paramètres

	// Naviguer entre les paramètres avec les flèches haut/bas
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		g.paramIndex = (g.paramIndex - 1 + NumParams) % NumParams
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		g.paramIndex = (g.paramIndex + 1) % NumParams
	}

	// Modifier la valeur du paramètre sélectionné avec les flèches gauche/droite
	delta := 0.0
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if g.paramIndex == ParamRotationSpeed {
			delta = 0.001 // Delta plus petit pour la rotation
		} else {
			delta = 0.01
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		delta = -0.01
	}

	if delta != 0 {
		switch g.paramIndex {
		case ParamGravity:
			if config.General.Gravity+delta >= 0 {
				config.General.Gravity += delta
			}
		case ParamFriction:
			// Friction doit être entre 0 et 1
			newFriction := config.General.Friction + delta
			if newFriction > 0 && newFriction < 2 {
				config.General.Friction = newFriction
			}
		case ParamSpawnRate:
			if config.General.SpawnRate+delta >= 0 {
				config.General.SpawnRate += delta
			}
		case ParamRandomSpawn:
			// Toggle RandomSpawn avec les flèches gauche/droite
			config.General.RandomSpawn = !config.General.RandomSpawn
		case ParamPartGener:
			// Modifier PartGener avec les flèches
			if config.General.PartGener+int(delta*100) >= 0 {
				config.General.PartGener += int(delta * 100)
			}
		case ParamRotationSpeed:
			// Modifier RotationSpeed avec les flèches (minimum 0)
			if config.General.RotationSpeed+delta >= 0 {
				config.General.RotationSpeed += delta
			}
		}
	}

	return nil
}
