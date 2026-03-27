package particle

// Particle définit une particule.
// Elle possède une position, une rotation, une taille, une couleur, et une
// opacité. Vous ajouterez certainement d'autres caractéristiques aux particules
// durant le projet.

type Particle struct {
	
    PositionX, PositionY float64   // position
    VelocityX, VelocityY float64   // vitesse
    Rotation             float64   // rotation
    ScaleX, ScaleY       float64   // taille
    ColorRed, ColorGreen, ColorBlue float64 // couleur
    Opacity              float64   // transparence
}


