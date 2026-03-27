package particle

import (
	"testing"
)


// TestNewParticle vérifie que la fonction NewParticle crée une particule correcte.
func TestNewParticle(t *testing.T) {

    // On crée une particule avec position (10,20) et vitesse (1.5,-0.5)
    p := NewParticle(10, 20, 1.5, -0.5)

    // ici on Vérifie la position
    if p.PositionX != 10 || p.PositionY != 20 {
        t.Errorf("Position incorrecte : attendu (10,20), obtenu (%v,%v)", p.PositionX, p.PositionY)
    }

    // ici on vérifie la vitesse
    if p.VelocityX != 1.5 || p.VelocityY != -0.5 {
        t.Errorf("Vitesse incorrecte : attendu (1.5,-0.5), obtenu (%v,%v)", p.VelocityX, p.VelocityY)
    }

    // ici on Vérifie les valeurs par défaut (taille, couleur, opacité)
    if p.ScaleX != 1 || p.ScaleY != 1 {
        t.Errorf("Taille incorrecte : attendu (1,1), obtenu (%v,%v)", p.ScaleX, p.ScaleY)
    }
    if p.ColorRed < 0 || p.ColorRed > 1 ||
       p.ColorGreen < 0 || p.ColorGreen > 1 ||
       p.ColorBlue < 0 || p.ColorBlue > 1 {
        t.Errorf("Couleur incorrecte : attendu valeurs entre 0 et 1, obtenu (%v,%v,%v)",
            p.ColorRed, p.ColorGreen, p.ColorBlue)
    }
    if p.Opacity != 1 {
        t.Errorf("Opacité incorrecte : attendu 1, obtenu %v", p.Opacity)
    }
}
