package main

import (
	"fmt"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/particle"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

// Draw se charge d'afficher à l'écran l'état actuel du système de particules
// g.system. Elle est appelée automatiquement environ 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction pourra être légèrement modifiée quand
// c'est précisé dans le sujet.
func (g *game) Draw(screen *ebiten.Image) {

	for e := g.system.Content.Front(); e != nil; e = e.Next() {
		//p, ok := e.Value.(*particles.Particle)
		p, ok := e.Value.(*particle.Particle)
		if ok {
			options := ebiten.DrawImageOptions{}
			options.GeoM.Rotate(p.Rotation)
			options.GeoM.Scale(p.ScaleX, p.ScaleY)
			options.GeoM.Translate(p.PositionX, p.PositionY)
			//options.ColorM.Scale(p.ColorRed, p.ColorGreen, p.ColorBlue, p.Opacity)
			options.ColorScale.Scale(float32(p.ColorRed), float32(p.ColorGreen), float32(p.ColorBlue), float32(p.Opacity))
			screen.DrawImage(assets.ParticleImage, &options)
		}
	}

	if config.General.Debug {
		ebitenutil.DebugPrint(screen, fmt.Sprint(ebiten.CurrentTPS()))
	}

	// Affichage de l'interface de modification des paramètres
	drawParameterInterface(screen, g)
}

// drawParameterInterface affiche l'interface permettant de modifier les paramètres
// du système de particules
func drawParameterInterface(screen *ebiten.Image, g *game) {
	// Titre
	text.Draw(screen, "\n=== Parametres (fleches haut ou bas pour naviguer et droite ou gauche pour modifier) ===\n", basicfont.Face7x13, 10, 20, colorred)

	// Affichage des paramètres
	y := 40
	paramNames := []string{"Gravite", "Friction", "SpawnRate", "RandomSpawn", "PartGener", "RotationSpeed"}
	paramValues := []float64{
		config.General.Gravity,
		config.General.Friction,
		config.General.SpawnRate,
		boolToFloat64(config.General.RandomSpawn),
		float64(config.General.PartGener),
		config.General.RotationSpeed,
	}

	for i := 0; i < len(paramValues); i++ {
		// Sélecteur
		selector := "  "
		if i == g.paramIndex {
			selector = "-> "
		}

		// Texte du paramètre
		paramText := fmt.Sprintf("\n%s%-12s: %.4f", selector, paramNames[i], paramValues[i])
		text.Draw(screen, paramText, basicfont.Face7x13, 10, y, colorred)
		y += 20
	}

	// Instructions
	y += 10
	text.Draw(screen, "\n↑/↓: sélectionner   \n←/→: modifier", basicfont.Face7x13, 10, y, colorGris)
}

// boolToFloat64 convertit un booléen en float64 (false -> 0.0, true -> 1.0)
func boolToFloat64(b bool) float64 {
	if b == true {
		return 1.0
	}
	return 0.0
}
