package particle

import (
	"project-particles/config"
)

// NewParticle crée une particule avec des valeurs par défaut ou optionnelles.
// Les arguments optionnels (ordre) : PositionX, PositionY, Rotation, ScaleX,
// ScaleY, ColorRed, ColorGreen, ColorBlue, Opacity, VelocityX, VelocityY,
// Age, MaxAge. Les 9 premiers champs (PositionX..Opacity) sont remplis
// right-aligned par rapport aux arguments fournis. Les arguments supplémentaires
// après ces 9 champs sont mappés dans l'ordre sur VelocityX, VelocityY, Age, MaxAge.
func NewParticle(optionals ...float64) Particle {
	// valeurs par défaut : centre de la fenêtre, rotation 0, échelle 1, couleurs 1, opacité 1,
	// vitesses 0, age 0, MaxAge 0
	defauts := []float64{
		float64(config.General.WindowSizeX) / 2, // PositionX
		float64(config.General.WindowSizeY) / 2, // PositionY
		0,                                       // Rotation
		1,                                       // ScaleX
		1,                                       // ScaleY
		1,                                       // ColorRed
		1,                                       // ColorGreen
		1,                                       // ColorBlue
		1,                                       // Opacity
	}

	// `dyn` contient les paramètres dynamiques : VelocityX, VelocityY, Age, MaxAge
	// Valeurs par défaut = 0 (vitesse nulle, âge 0, MaxAge 0)
	dyn := []float64{0, 0, 0, 0} // VelocityX, VelocityY, Age, MaxAge

	// Si aucun argument, retourner defaults
	if len(optionals) == 0 {
		return Particle{
			PositionX:  defauts[0],
			PositionY:  defauts[1],
			Rotation:   defauts[2],
			ScaleX:     defauts[3],
			ScaleY:     defauts[4],
			ColorRed:   defauts[5],
			ColorGreen: defauts[6],
			ColorBlue:  defauts[7],
			Opacity:    defauts[8],
			VelocityX:  dyn[0],
			VelocityY:  dyn[1],
			Age:        int(dyn[2]),
			MaxAge:     int(dyn[3]),
		}
	}

	// `attri` contient les 9 champs  (position..opacity).
	// On copie les valeurs par défaut puis on remplace right-aligned
	attri := make([]float64, len(defauts))
	copy(attri, defauts)
	lo := len(optionals)
	// Le nombre d'arguments d'apparition fournis est au maximum de len(defaults)
	if lo <= 9 {
		// Si on a moins d'arguments , on les aligne à droite.
		start := 9 - lo
		for i := 0; i < lo; i++ {
			attri[start+i] = optionals[i]
		}
		// dyn reste avec ses valeurs par défaut
	} else {
		// Si on a plus d'arguments, les premiers remplissent l'apparence,
		// les suivants remplissent `dyn` dans l'ordre (VelocityX, VelocityY, Age, MaxAge).
		for i := 0; i < len(defauts); i++ {
			attri[i] = optionals[i]
		}
		rem := optionals[len(defauts):] //garde le reste des argument
		for i := 0; i < len(dyn) && i < len(rem); i++ {
			dyn[i] = rem[i]
		}
	}

	return Particle{
		PositionX:  attri[0],
		PositionY:  attri[1],
		Rotation:   attri[2],
		ScaleX:     attri[3],
		ScaleY:     attri[4],
		ColorRed:   attri[5],
		ColorGreen: attri[6],
		ColorBlue:  attri[7],
		Opacity:    attri[8],
		VelocityX:  dyn[0],
		VelocityY:  dyn[1],
		Age:        int(dyn[2]),
		MaxAge:     int(dyn[3]),
	}
}
