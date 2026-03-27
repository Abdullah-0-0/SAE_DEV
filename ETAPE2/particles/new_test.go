package particles

import (
	"container/list"
	"math/rand"
	"project-particles/config"
	"project-particles/particle"
	"testing"
)

// TestCreeSystemSpawnFixe teste la création d'un système de particules avec spawn à position fixe.
// Elle vérifie que le nombre correct de particules est créé et qu'elles sont toutes
// positionnées à la même location (SpawnX, SpawnY) avec les bons paramètres initiaux.
func TestCreeSystemSpawnFixe(t *testing.T) {
	old := config.General
	defer func() { config.General = old }()

	config.General.PartGener = 3
	config.General.RandomSpawn = false
	config.General.SpawnX = 10
	config.General.SpawnY = 20

	sys := NewSystem()
	// Le système commence vide, on appelle le générateur pour créer les particules
	sys.Generateur()
	if sys.Content.Len() != 3 {
		t.Fatalf("On attendait 3 particules, on en a eu %d", sys.Content.Len())
	}
	if sys.SpawnAccumulator != 0 {
		t.Errorf("SpawnAccumulator devrait être égal à 0, mais on a obtenu %v", sys.SpawnAccumulator)
	}

	for e := sys.Content.Front(); e != nil; e = e.Next() {
		p, ok := e.Value.(*particle.Particle)
		if !ok {
			t.Fatalf("L'élément de liste n'est pas *particle.Particle")
		}
		if p.PositionX != float64(config.General.SpawnX) || p.PositionY != float64(config.General.SpawnY) {
			t.Errorf("Position inattendue : obtenu (%v,%v), attendu (%v,%v)", p.PositionX, p.PositionY, config.General.SpawnX, config.General.SpawnY)
		}
		if p.ScaleX != 1.5 || p.ScaleY != 1.5 {
			t.Errorf("Échelle inattendue : %v,%v", p.ScaleX, p.ScaleY)
		}
		if p.Opacity != 1 {
			t.Errorf("Opacité inattendue : %v", p.Opacity)
		}
	}
}

// TestCreeSystemSpawnAleatoire teste la création d'un système de particules avec spawn aléatoire.
// Elle vérifie que les positions générées sont bien dans les limites de la fenêtre
// et que les valeurs d'âge maximum sont correctes.
func TestCreeSystemSpawnAleatoire(t *testing.T) {
	old := config.General
	defer func() { config.General = old }()

	rand.Seed(42)
	config.General.PartGener = 2
	config.General.RandomSpawn = true
	config.General.WindowSizeX = 100
	config.General.WindowSizeY = 50

	sys := NewSystem()
	// Le système commence vide, on appelle le générateur pour créer les particules
	sys.Generateur()
	if sys.Content.Len() != 2 {
		t.Fatalf("On attendait 2 particules, on en a eu %d", sys.Content.Len())
	}

	for e := sys.Content.Front(); e != nil; e = e.Next() {
		p := e.Value.(*particle.Particle)
		if p.PositionX < 0 || p.PositionX >= float64(config.General.WindowSizeX) {
			t.Errorf("PositionX hors plage : %v", p.PositionX)
		}
		if p.PositionY < 0 || p.PositionY >= float64(config.General.WindowSizeY) {
			t.Errorf("PositionY hors plage : %v", p.PositionY)
		}
		if p.MaxAge < 300 || p.MaxAge >= 600 {
			t.Errorf("Âge maximal hors de la plage attendue : %d", p.MaxAge)
		}
	}
}

// TestOpaciteDiminue teste que l'opacité d'une particule diminue au cours du temps.
// Elle crée une particule avec une opacité de 1.0 et vérifie qu'après l'appel
// à Update(), l'opacité est inférieure à 1.0 (elle vieillit et devient transparente).
func TestOpaciteDiminue(t *testing.T) {
	config.General.SpawnRate = 0
	config.General.InitNumParticles = 0
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
	config.General.OffscreenMargin = 0

	sys := System{Content: list.New(), SpawnAccumulator: 0}

	p := &particle.Particle{
		PositionX: 10, PositionY: 10,
		Rotation: 0,
		ScaleX:   1, ScaleY: 1,
		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity:   1.0,
		VelocityX: 0, VelocityY: 0,
		Age:    0,
		MaxAge: 10,
	}

	sys.Content.PushBack(p)

	sys.Update()

	if p.Opacity >= 1.0 {
		t.Fatalf("On attend une opacité < 1 après Update, obtenu : %v", p.Opacity)
	}
}

// TestGraviteAffecteVitesse teste que la gravité augmente la vélocité verticale des particules.
// Elle crée une particule avec une vélocité Y nulle et vérifie qu'après Update(),
// la vélocité Y augmente (la gravité attire la particule vers le bas).
func TestGraviteAffecteVitesse(t *testing.T) {
	old := config.General
	defer func() { config.General = old }()

	config.General.Gravity = 9.81
	config.General.Friction = 1.0
	config.General.SpawnRate = 0
	config.General.InitNumParticles = 0
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
	config.General.OffscreenMargin = 0

	sys := System{Content: list.New(), SpawnAccumulator: 0}

	p := &particle.Particle{
		PositionX: 100, PositionY: 100,
		Rotation: 0,
		ScaleX:   1, ScaleY: 1,
		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity:   1.0,
		VelocityX: 0, VelocityY: 0,
		Age:    0,
		MaxAge: 100,
	}

	sys.Content.PushBack(p)
	velociteBefore := p.VelocityY

	sys.Update()

	if p.VelocityY <= velociteBefore {
		t.Fatalf("La gravité devrait augmenter VelocityY, avant=%v après=%v", velociteBefore, p.VelocityY)
	}
}

// TestFrictionRalentitParticule teste que la friction ralentit les particules.
// Elle crée une particule avec une vélocité non nulle et vérifie qu'après Update(),
// la vélocité est réduite (multiplicée par le coefficient de friction).
func TestFrictionRalentitParticule(t *testing.T) {
	old := config.General
	defer func() { config.General = old }()

	config.General.Gravity = 0
	config.General.Friction = 0.9
	config.General.SpawnRate = 0
	config.General.InitNumParticles = 0
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
	config.General.OffscreenMargin = 0

	sys := System{Content: list.New(), SpawnAccumulator: 0}

	p := &particle.Particle{
		PositionX: 100, PositionY: 100,
		Rotation: 0,
		ScaleX:   1, ScaleY: 1,
		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity:   1.0,
		VelocityX: 10, VelocityY: 10,
		Age:    0,
		MaxAge: 100,
	}

	sys.Content.PushBack(p)
	vitesseXBefore := p.VelocityX
	vitesseYBefore := p.VelocityY

	sys.Update()

	if p.VelocityX >= vitesseXBefore {
		t.Fatalf("La friction devrait réduire VelocityX, avant=%v après=%v", vitesseXBefore, p.VelocityX)
	}
	if p.VelocityY >= vitesseYBefore {
		t.Fatalf("La friction devrait réduire VelocityY, avant=%v après=%v", vitesseYBefore, p.VelocityY)
	}
}

// TestMiseAJourPosition teste que la position des particules est correctement mise à jour
// en fonction de leur vélocité. Elle crée une particule avec une vélocité et vérifie
// qu'après Update(), la position a changé de manière cohérente.
func TestMiseAJourPosition(t *testing.T) {
	old := config.General
	defer func() { config.General = old }()

	config.General.Gravity = 0
	config.General.Friction = 1.0
	config.General.SpawnRate = 0
	config.General.InitNumParticles = 0
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
	config.General.OffscreenMargin = 0

	sys := System{Content: list.New(), SpawnAccumulator: 0}

	p := &particle.Particle{
		PositionX: 100, PositionY: 100,
		Rotation: 0,
		ScaleX:   1, ScaleY: 1,
		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity:   1.0,
		VelocityX: 5, VelocityY: 3,
		Age:    0,
		MaxAge: 100,
	}

	sys.Content.PushBack(p)
	posXBefore := p.PositionX
	posYBefore := p.PositionY

	sys.Update()

	if p.PositionX != posXBefore+5 {
		t.Fatalf("Position X devrait augmenter de 5, avant=%v après=%v", posXBefore, p.PositionX)
	}
	if p.PositionY != posYBefore+3 {
		t.Fatalf("Position Y devrait augmenter de 3, avant=%v après=%v", posYBefore, p.PositionY)
	}
}

// TestParticuleAgeMaxSupprimee teste que les particules atteignant leur âge maximum
// sont supprimées du système. Elle crée une particule avec Age = MaxAge et
// vérifie qu'elle est supprimée après Update().
func TestParticuleAgeMaxSupprimee(t *testing.T) {
	old := config.General
	defer func() { config.General = old }()

	config.General.Gravity = 0
	config.General.Friction = 1.0
	config.General.SpawnRate = 0
	config.General.InitNumParticles = 0
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
	config.General.OffscreenMargin = 0

	sys := System{Content: list.New(), SpawnAccumulator: 0}

	p := &particle.Particle{
		PositionX: 100, PositionY: 100,
		Rotation: 0,
		ScaleX:   1, ScaleY: 1,
		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity:   0.1,
		VelocityX: 0, VelocityY: 0,
		Age:    10,
		MaxAge: 10,
	}

	sys.Content.PushBack(p)

	if sys.Content.Len() != 1 {
		t.Fatalf("Le système devrait contenir 1 particule avant Update, obtenu %d", sys.Content.Len())
	}

	sys.Update()

	if sys.Content.Len() != 0 {
		t.Fatalf("Le système devrait être vide après Update (particule trop vieille), obtenu %d", sys.Content.Len())
	}
}

// TestRotationAvance teste que la rotation des particules augmente à chaque mise à jour.
// Elle crée une particule avec une rotation initiale et vérifie qu'après Update(),
// la rotation a augmenté selon RotationSpeed.
func TestRotationAvance(t *testing.T) {
	old := config.General
	defer func() { config.General = old }()

	config.General.Gravity = 0
	config.General.Friction = 1.0
	config.General.SpawnRate = 0
	config.General.InitNumParticles = 0
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
	config.General.OffscreenMargin = 0
	config.General.RotationSpeed = 5.0

	sys := System{Content: list.New(), SpawnAccumulator: 0}

	p := &particle.Particle{
		PositionX: 100, PositionY: 100,
		Rotation: 0,
		ScaleX:   1, ScaleY: 1,
		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity:   1.0,
		VelocityX: 0, VelocityY: 0,
		Age:    0,
		MaxAge: 100,
	}

	sys.Content.PushBack(p)
	rotationBefore := p.Rotation

	sys.Update()

	if p.Rotation != rotationBefore+5.0 {
		t.Fatalf("La rotation devrait augmenter de %v, avant=%v après=%v", config.General.RotationSpeed, rotationBefore, p.Rotation)
	}
}

// TestParticulesHorsEcranSupprimees teste que les particules sortant de la fenêtre d'affichage
// sont supprimées du système. Elle crée une particule loin hors écran et
// vérifie qu'elle est supprimée après Update().
func TestParticulesHorsEcranSupprimees(t *testing.T) {
	old := config.General
	defer func() { config.General = old }()

	config.General.Gravity = 0
	config.General.Friction = 1.0
	config.General.SpawnRate = 0
	config.General.InitNumParticles = 0
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
	config.General.OffscreenMargin = 10

	sys := System{Content: list.New(), SpawnAccumulator: 0}

	p := &particle.Particle{
		PositionX: -50, PositionY: 100,
		Rotation: 0,
		ScaleX:   1, ScaleY: 1,
		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity:   1.0,
		VelocityX: 0, VelocityY: 0,
		Age:    0,
		MaxAge: 100,
	}

	sys.Content.PushBack(p)

	if sys.Content.Len() != 1 {
		t.Fatalf("Le système devrait contenir 1 particule avant Update, obtenu %d", sys.Content.Len())
	}

	sys.Update()

	if sys.Content.Len() != 0 {
		t.Fatalf("Le système devrait être vide après Update (particule hors écran), obtenu %d", sys.Content.Len())
	}
}

// TestAccumulateurSpawnFonctionnel teste que l'accumulateur de spawn accumule correctement
// le taux de spawn et génère les bonnes particules. Elle configure un taux de spawn
// et vérifie que les particules sont générées au bon moment.
func TestAccumulateurSpawnFonctionnel(t *testing.T) {
	old := config.General
	defer func() { config.General = old }()

	config.General.Gravity = 0
	config.General.Friction = 1.0
	config.General.SpawnRate = 0.5
	config.General.InitNumParticles = 0
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
	config.General.OffscreenMargin = 0
	config.General.PartGener = 1
	config.General.RandomSpawn = false
	config.General.SpawnX = 400
	config.General.SpawnY = 300

	sys := NewSystem()

	if sys.SpawnAccumulator != 0 {
		t.Fatalf("L'accumulateur devrait être initialisé à 0, obtenu %v", sys.SpawnAccumulator)
	}

	// Première Update : accumule 0.5 (pas de spawn)
	sys.Update()
	if sys.Content.Len() != 0 {
		t.Fatalf("Aucune particule ne devrait être générée, obtenu %d", sys.Content.Len())
	}
	if sys.SpawnAccumulator != 0.5 {
		t.Fatalf("L'accumulateur devrait être 0.5, obtenu %v", sys.SpawnAccumulator)
	}

	// Deuxième Update : accumule 0.5 + 0.5 = 1.0 (spawn 1 particule)
	sys.Update()
	if sys.Content.Len() != 1 {
		t.Fatalf("Une particule devrait être générée, obtenu %d", sys.Content.Len())
	}
	if sys.SpawnAccumulator != 0 {
		t.Fatalf("L'accumulateur devrait être 0 après spawn, obtenu %v", sys.SpawnAccumulator)
	}
}
