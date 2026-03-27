package config

// Config définit les champs qu'on peut trouver dans un fichier de config.
// Dans le fichier les champs doivent porter le même nom que dans le type si
// dessous, y compris les majuscules. Tous les champs doivent obligatoirement
// commencer par des majuscules, sinon il ne sera pas possible de récupérer
// leurs valeurs depuis le fichier de config.
// Vous pouvez ajouter des champs et ils seront automatiquement lus dans le
// fichier de config. Vous devrez le faire plusieurs fois durant le projet.
type Config struct {
	WindowTitle              string
	WindowSizeX, WindowSizeY int
	ParticleImage            string
	Debug                    bool
	InitNumParticles         int
	RandomSpawn              bool
	SpawnX, SpawnY           int
	SpawnRate                float64
	// PartGener : nombre de particules générées à chaque appel du générateur
	PartGener int
	// SpawnRadius : si >0, limite le spawn aléatoire à l'intérieur d'un
	SpawnRadius int
	// Gravity : accélération gravitationnelle appliquée aux particules (pixels/frame^2)
	Gravity float64
	// OffscreenMargin : marge (en pixels) au-delà de laquelle une particule est
	// considérée comme définitivement hors écran et peut être supprimée.
	OffscreenMargin int
	// Friction : coefficient de friction appliqué aux particules (0 < Friction < 1) par preference
	Friction float64
	// RotationSpeed : vitesse de rotation des particules (radians par frame)
	RotationSpeed float64
}

var General Config
