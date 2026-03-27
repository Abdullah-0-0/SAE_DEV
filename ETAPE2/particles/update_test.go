package particles

import (
	"container/list"
	"project-particles/config"
	"project-particles/particle"
	"testing"
)

// TestMiseAJourGraviteEtFriction teste que la gravité et la friction sont correctement appliquées
// à toutes les particules du système. Elle crée plusieurs particules avec des vélocités différentes
// et vérifie que la gravité augmente VelocityY et que la friction les ralentit.
func TestMiseAJourGraviteEtFriction(t *testing.T) {
	old := config.General
	defer func() { config.General = old }()

	config.General.Gravity = 10.0
	config.General.Friction = 0.8
	config.General.SpawnRate = 0
	config.General.InitNumParticles = 0
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
	config.General.OffscreenMargin = 0

	sys := System{Content: list.New(), SpawnAccumulator: 0}

	// Particule 1
	p1 := &particle.Particle{
		PositionX: 100, PositionY: 100,
		Rotation: 0,
		ScaleX:   1, ScaleY: 1,
		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity:   1.0,
		VelocityX: 20, VelocityY: 5,
		Age:    0,
		MaxAge: 100,
	}

	sys.Content.PushBack(p1)

	v1XBefore := p1.VelocityX
	v1YBefore := p1.VelocityY

	sys.Update()

	// VelocityX devrait être multiplié par friction
	if p1.VelocityX != v1XBefore*0.8 {
		t.Fatalf("VelocityX de P1 devrait être réduit par friction, avant=%v après=%v", v1XBefore, p1.VelocityX)
	}

	// VelocityY devrait avoir la gravité ajoutée puis multiplié par friction
	expectedVY := (v1YBefore + 10.0) * 0.8
	if p1.VelocityY != expectedVY {
		t.Fatalf("VelocityY de P1 incorrect, attendu=%v obtenu=%v", expectedVY, p1.VelocityY)
	}
}

// TestMiseAJourAgeParticule teste que l'âge des particules augmente correctement
// à chaque appel à Update(). Elle crée une particule avec Age = 0 et vérifie
// qu'après plusieurs mises à jour, l'âge augmente de 1 à chaque fois.
func TestMiseAJourAgeParticule(t *testing.T) {
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
		VelocityX: 0, VelocityY: 0,
		Age:    0,
		MaxAge: 50,
	}

	sys.Content.PushBack(p)

	for i := 0; i < 5; i++ {
		if p.Age != i {
			t.Fatalf("Itération %d : Age devrait être %d, obtenu %d", i, i, p.Age)
		}
		sys.Update()
	}

	if p.Age != 5 {
		t.Fatalf("Après 5 mises à jour, Age devrait être 5, obtenu %d", p.Age)
	}
}

// TestMiseAJourOpaciteDecroissante teste que l'opacité des particules décroît linéairement
// avec leur âge jusqu'à atteindre 0 quand MaxAge est atteint.
func TestMiseAJourOpaciteDecroissante(t *testing.T) {
	old := config.General
	defer func() { config.General = old }()

	config.General.Gravity = 0
	config.General.Friction = 1.0
	config.General.SpawnRate = 0
	config.General.InitNumParticles = 0
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
	config.General.OffscreenMargin = 0
	config.General.RotationSpeed = 0

	sys := System{Content: list.New(), SpawnAccumulator: 0}

	p := &particle.Particle{
		PositionX: 100, PositionY: 100,
		Rotation: 0,
		ScaleX:   1, ScaleY: 1,
		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity:   1.0,
		VelocityX: 0, VelocityY: 0,
		Age:    0,
		MaxAge: 10,
	}

	sys.Content.PushBack(p)

	// Vérifier que l'opacité commence à 1.0
	if p.Opacity != 1.0 {
		t.Fatalf("À Age=0 : Opacity devrait être 1.0, obtenu %v", p.Opacity)
	}

	// Chaque mise à jour, l'opacité devrait diminuer
	for i := 1; i <= 10; i++ {
		sys.Update()
		expectedOpacity := 1.0 - float64(i)/10.0
		// Utiliser une tolérance pour la précision flottante
		if p.Opacity < expectedOpacity-0.01 || p.Opacity > expectedOpacity+0.01 {
			t.Fatalf("À Age=%d : Opacity devrait être ~%v, obtenu %v", i, expectedOpacity, p.Opacity)
		}
	}

	// À Age=10 (MaxAge), l'opacité devrait être 0
	if p.Opacity != 0.0 {
		t.Fatalf("À Age=10 (MaxAge) : Opacity devrait être 0.0, obtenu %v", p.Opacity)
	}
}

// TestMiseAJourRotationParticule teste que la rotation des particules augmente
// correctement selon le RotationSpeed. Elle crée une particule et fait plusieurs
// mises à jour pour vérifier que la rotation s'accumule correctement.
func TestMiseAJourRotationParticule(t *testing.T) {
	old := config.General
	defer func() { config.General = old }()

	config.General.Gravity = 0
	config.General.Friction = 1.0
	config.General.SpawnRate = 0
	config.General.InitNumParticles = 0
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
	config.General.OffscreenMargin = 0
	config.General.RotationSpeed = 15.0

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

	for i := 0; i < 5; i++ {
		sys.Update()
		expectedRotation := float64((i + 1) * 15)
		if p.Rotation != expectedRotation {
			t.Fatalf("Itération %d : Rotation devrait être %v, obtenu %v", i, expectedRotation, p.Rotation)
		}
	}
}

// TestMiseAJourSuppressionsMultiples teste que plusieurs particules expirées
// sont correctement supprimées du système en une seule Update().
func TestMiseAJourSuppressionsMultiples(t *testing.T) {
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

	// Créer 3 particules : 2 vont expirer après cette Update, 1 continuera
	p1 := &particle.Particle{
		PositionX: 100, PositionY: 100,
		Rotation: 0, ScaleX: 1, ScaleY: 1,
		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity: 1.0, VelocityX: 0, VelocityY: 0,
		Age: 4, MaxAge: 5, // Après Update: Age=5, donc supprimée (Age >= MaxAge)
	}

	p2 := &particle.Particle{
		PositionX: 200, PositionY: 200,
		Rotation: 0, ScaleX: 1, ScaleY: 1,
		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity: 1.0, VelocityX: 0, VelocityY: 0,
		Age: 8, MaxAge: 20, // Après Update: Age=9, reste vivante
	}

	p3 := &particle.Particle{
		PositionX: 300, PositionY: 300,
		Rotation: 0, ScaleX: 1, ScaleY: 1,
		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity: 1.0, VelocityX: 0, VelocityY: 0,
		Age: 4, MaxAge: 5, // Après Update: Age=5, donc supprimée (Age >= MaxAge)
	}

	sys.Content.PushBack(p1)
	sys.Content.PushBack(p2)
	sys.Content.PushBack(p3)

	if sys.Content.Len() != 3 {
		t.Fatalf("Le système devrait contenir 3 particules, obtenu %d", sys.Content.Len())
	}

	sys.Update()

	if sys.Content.Len() != 1 {
		t.Fatalf("Après Update, le système devrait contenir 1 particule, obtenu %d", sys.Content.Len())
	}

	// Vérifier que c'est bien p2 qui reste
	remaining := sys.Content.Front().Value.(*particle.Particle)
	if remaining.Age != 9 {
		t.Fatalf("La particule restante devrait avoir Age=9, obtenu %d", remaining.Age)
	}
}

// TestMiseAJourParticulesHorsEcran teste que plusieurs variantes de positions
// hors écran sont correctement détectées et supprimées.
func TestMiseAJourParticulesHorsEcran(t *testing.T) {
	old := config.General
	defer func() { config.General = old }()

	config.General.Gravity = 0
	config.General.Friction = 1.0
	config.General.SpawnRate = 0
	config.General.InitNumParticles = 0
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
	config.General.OffscreenMargin = 50

	sys := System{Content: list.New(), SpawnAccumulator: 0}

	// Particule trop à gauche
	p1 := &particle.Particle{
		PositionX: -60, PositionY: 300, Rotation: 0,
		ScaleX: 1, ScaleY: 1, ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity: 1.0, VelocityX: 0, VelocityY: 0,
		Age: 0, MaxAge: 100,
	}

	// Particule trop à droite
	p2 := &particle.Particle{
		PositionX: 860, PositionY: 300, Rotation: 0,
		ScaleX: 1, ScaleY: 1, ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity: 1.0, VelocityX: 0, VelocityY: 0,
		Age: 0, MaxAge: 100,
	}

	// Particule trop haut
	p3 := &particle.Particle{
		PositionX: 400, PositionY: -60, Rotation: 0,
		ScaleX: 1, ScaleY: 1, ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity: 1.0, VelocityX: 0, VelocityY: 0,
		Age: 0, MaxAge: 100,
	}

	// Particule trop bas
	p4 := &particle.Particle{
		PositionX: 400, PositionY: 660, Rotation: 0,
		ScaleX: 1, ScaleY: 1, ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity: 1.0, VelocityX: 0, VelocityY: 0,
		Age: 0, MaxAge: 100,
	}

	// Particule valide à l'intérieur
	p5 := &particle.Particle{
		PositionX: 400, PositionY: 300, Rotation: 0,
		ScaleX: 1, ScaleY: 1, ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity: 1.0, VelocityX: 0, VelocityY: 0,
		Age: 0, MaxAge: 100,
	}

	sys.Content.PushBack(p1)
	sys.Content.PushBack(p2)
	sys.Content.PushBack(p3)
	sys.Content.PushBack(p4)
	sys.Content.PushBack(p5)

	if sys.Content.Len() != 5 {
		t.Fatalf("Le système devrait contenir 5 particules, obtenu %d", sys.Content.Len())
	}

	sys.Update()

	if sys.Content.Len() != 1 {
		t.Fatalf("Après Update, seule 1 particule devrait rester, obtenu %d", sys.Content.Len())
	}
}

// TestMiseAJourAccumulateurSpawn teste que l'accumulateur de spawn fonctionne correctement
// et génère les particules au bon moment.
func TestMiseAJourAccumulateurSpawn(t *testing.T) {
	old := config.General
	defer func() { config.General = old }()

	config.General.Gravity = 0
	config.General.Friction = 1.0
	config.General.SpawnRate = 0.3
	config.General.InitNumParticles = 0
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
	config.General.OffscreenMargin = 0
	config.General.PartGener = 1
	config.General.RandomSpawn = false
	config.General.SpawnX = 400
	config.General.SpawnY = 300

	sys := NewSystem()

	// Itération 1 : accumule 0.3
	sys.Update()
	if sys.Content.Len() != 0 {
		t.Fatalf("Itération 1 : Aucune particule devrait être générée, obtenu %d", sys.Content.Len())
	}
	if sys.SpawnAccumulator < 0.29 || sys.SpawnAccumulator > 0.31 {
		t.Fatalf("Itération 1 : Accumulateur devrait être ~0.3, obtenu %v", sys.SpawnAccumulator)
	}

	// Itération 2 : accumule 0.3 + 0.3 = 0.6 (pas de spawn)
	sys.Update()
	if sys.Content.Len() != 0 {
		t.Fatalf("Itération 2 : Aucune particule devrait être générée, obtenu %d", sys.Content.Len())
	}
	if sys.SpawnAccumulator < 0.59 || sys.SpawnAccumulator > 0.61 {
		t.Fatalf("Itération 2 : Accumulateur devrait être ~0.6, obtenu %v", sys.SpawnAccumulator)
	}

	// Itération 3 : accumule 0.6 + 0.3 = 0.9 (pas de spawn)
	sys.Update()
	if sys.Content.Len() != 0 {
		t.Fatalf("Itération 3 : Aucune particule devrait être générée, obtenu %d", sys.Content.Len())
	}
	if sys.SpawnAccumulator < 0.89 || sys.SpawnAccumulator > 0.91 {
		t.Fatalf("Itération 3 : Accumulateur devrait être ~0.9, obtenu %v", sys.SpawnAccumulator)
	}

	// Itération 4 : accumule 0.9 + 0.3 = 1.2 => spawn 1, reste ~0.2
	sys.Update()
	if sys.Content.Len() != 1 {
		t.Fatalf("Itération 4 : 1 particule devrait être générée, obtenu %d", sys.Content.Len())
	}
	if sys.SpawnAccumulator < 0.19 || sys.SpawnAccumulator > 0.21 {
		t.Fatalf("Itération 4 : Accumulateur devrait être ~0.2, obtenu %v", sys.SpawnAccumulator)
	}
}
