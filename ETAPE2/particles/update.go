package particles

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
import (
	"project-particles/config"
	"project-particles/particle"
)

func (s *System) Update() {
	// Gravité et friction
	//5.1 Gravité
	gravity := config.General.Gravity
	friction := config.General.Friction // Friction de l'air (ralentissement)

	// Gestion du taux d'apparition avec accumulation des fractions
	s.SpawnAccumulator += config.General.SpawnRate
	for s.SpawnAccumulator >= 1.0 {
		if s.Generateur != nil {
			s.Generateur()
		}
		s.SpawnAccumulator -= 1.0
	}

	// Parcourir toutes les particules
	for e := s.Content.Front(); e != nil; {
		next := e.Next()
		p, ok := e.Value.(*particle.Particle)

		if ok == false {
			e = next
			continue
		}

		// Appliquer la gravité
		p.VelocityY += gravity

		// Appliquer la friction
		p.VelocityX *= friction
		p.VelocityY *= friction

		// Mettre à jour la position en fonction de la vélocité
		p.PositionX += p.VelocityX
		p.PositionY += p.VelocityY

		//5.2 Extérieur de l’écran
		// Supprimer les particules définitivement hors écran (avec marge)
		if config.General.OffscreenMargin > 0 {
			margin := float64(config.General.OffscreenMargin)
			if p.PositionX < -margin || p.PositionX > float64(config.General.WindowSizeX)+margin || p.PositionY < -margin || p.PositionY > float64(config.General.WindowSizeY)+margin {
				s.Content.Remove(e)
				e = next
				continue
			}
		}

		// Incrémenter l'âge de la particule
		p.Age++

		// Faire tourner la particule
		p.Rotation += config.General.RotationSpeed

		// Réduire l'opacité au fil du temps
		if p.MaxAge > 0 {
			opacityReduction := float64(p.Age) / float64(p.MaxAge)
			if opacityReduction > 1 {
				p.Opacity = 0
			} else {
				p.Opacity = 1.0 - opacityReduction
			}
		}

		// suppression si morte
		if p.Age >= p.MaxAge {
			s.Content.Remove(e)
		}
		e = next
	}
}
