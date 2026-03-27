package particle

import (
	"math"
	"project-particles/config"
	"reflect"
	"testing"
)

// TestCreeParticuleArgumentsParDefaut teste que NewParticle() sans arguments
// crée une particule avec les valeurs par défaut (centrée,
// échelle 1, opacité 1, vitesse nulle, âge 0, MaxAge 0).
func TestCreeParticuleArgumentsParDefaut(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX) / 2,
		float64(config.General.WindowSizeY) / 2,
		0,
		1,
		1,
		1,
		1,
		1,
		1,
		0,
		0,
		0,
		0,
	}
	result := NewParticle()
	if !reflect.DeepEqual(result, expected) {
		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

// TestCreeParticuleUnArgument teste que le premier argument est interprété
// comme `occurance` (ici valeur 2) et que les autres valeurs
// par défaut restent correctes.
func TestCreeParticuleUnArgument(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX) / 2,
		float64(config.General.WindowSizeY) / 2,
		0,
		1,
		1,
		1,
		1,
		1,
		1,
		0,
		0,
		0,
		0,
	}
	result := NewParticle(1)
	if !reflect.DeepEqual(result, expected) {
		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

// TestCreeParticuleDeuxArguments teste que les deux premiers arguments
// définissent `appearance` et `velocityX` correctement.
func TestCreeParticuleDeuxArguments(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX) / 2,
		float64(config.General.WindowSizeY) / 2,
		0,
		1,
		1,
		1,
		1,
		0.5,
		1,
		0,
		0,
		0,
		0,
	}
	result := NewParticle(0.5, 1)
	if !reflect.DeepEqual(result, expected) {
		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

// TestCreeParticuleTroisArguments teste l'interprétation des trois premiers
// arguments (appearance, velocityX, velocityY) et l'âge.
func TestCreeParticuleTroisArguments(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX) / 2,
		float64(config.General.WindowSizeY) / 2,
		0,
		1,
		1,
		1,
		0.4,
		0.5,
		1,
		0,
		0,
		0,
		0,
	}
	result := NewParticle(0.4, 0.5, 1)
	if !reflect.DeepEqual(result, expected) {
		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

// TestCreeParticuleQuatreArguments teste la propagation des quatre premiers
// arguments et le positionnement par défaut (centre).
func TestCreeParticuleQuatreArguments(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX) / 2,
		float64(config.General.WindowSizeY) / 2,
		0,
		1,
		1,
		0.3,
		0.4,
		0.5,
		1,
		0,
		0,
		0,
		0,
	}
	result := NewParticle(0.3, 0.4, 0.5, 1)
	if !reflect.DeepEqual(result, expected) {
		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

// TestCreeParticuleCinqArguments teste que les cinq premiers arguments
// (incluant ScaleY) sont appliqués correctement.
func TestCreeParticuleCinqArguments(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX) / 2,
		float64(config.General.WindowSizeY) / 2,
		0,
		1,
		1.1,
		0.3,
		0.4,
		0.5,
		1,
		0,
		0,
		0,
		0,
	}
	result := NewParticle(1.1, 0.3, 0.4, 0.5, 1)
	if !reflect.DeepEqual(result, expected) {
		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

// TestCreeParticuleSixArguments teste six arguments (appearance, scaleX, scaleY,
// color components, etc.) et que l'ordre est respecté.
func TestCreeParticuleSixArguments(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX) / 2,
		float64(config.General.WindowSizeY) / 2,
		0,
		1.2,
		1.1,
		0.3,
		0.4,
		0.5,
		1,
		0,
		0,
		0,
		0,
	}
	result := NewParticle(1.2, 1.1, 0.3, 0.4, 0.5, 1)
	if !reflect.DeepEqual(result, expected) {
		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

// TestCreeParticuleSeptArguments teste sept arguments; ici le premier est
// `math.Pi` pour tester la gestion d'un float spécial.
func TestCreeParticuleSeptArguments(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX) / 2,
		float64(config.General.WindowSizeY) / 2,
		math.Pi,
		1.2,
		1.1,
		0.3,
		0.4,
		0.5,
		1,
		0,
		0,
		0,
		0,
	}
	result := NewParticle(math.Pi, 1.2, 1.1, 0.3, 0.4, 0.5, 1)
	if !reflect.DeepEqual(result, expected) {
		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

// TestCreeParticuleHuitArguments teste huit arguments et la bonne affectation
// des coordonnées et paramètres supplémentaires.
func TestCreeParticuleHuitArguments(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX) / 2,
		float64(config.General.WindowSizeY) / 2,
		math.Pi,
		1.2,
		1.1,
		0.3,
		0.4,
		0.5,
		1,
		0,
		0,
		0,
		0,
	}
	result := NewParticle(float64(config.General.WindowSizeY), math.Pi, 1.2, 1.1, 0.3, 0.4, 0.5, 1)
	if !reflect.DeepEqual(result, expected) {
		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

// TestCreeParticuleNeufArguments teste un cas plus complet (neuf arguments)
// pour s'assurer que l'ordre des paramètres est correct.
func TestCreeParticuleNeufArguments(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX) / 3,
		float64(config.General.WindowSizeY) / 2,
		math.Pi,
		1.2,
		1.1,
		0.3,
		0.4,
		0.5,
		1,
		0,
		0,
		0,
		0,
	}
	result := NewParticle(float64(config.General.WindowSizeY)/3, float64(config.General.WindowSizeY), math.Pi, 1.2, 1.1, 0.3, 0.4, 0.5, 1)
	if !reflect.DeepEqual(result, expected) {
		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}
