# Tests du Système de Particules

## Vue d'ensemble
Ce document présente tous les tests unitaires créés pour le système de particules en Go. Les noms des fonctions de test ont été renommés en français et chaque test inclut un commentaire explicatif décrivant ce qu'il teste.

## Tests du package `particle`

### 1. **TestCreeParticuleArgumentsParDefaut**
   - **Description**: Teste que `NewParticle()` sans arguments crée une particule avec les valeurs par défaut (centrée, échelle 1, opacité 1, vitesse nulle, âge 0, MaxAge 0).
   - **Fichier**: [particle/new_test.go](particle/new_test.go)

### 2. **TestCreeParticuleUnArgument**
   - **Description**: Teste que le premier argument est interprété comme `appearance` et que les autres valeurs par défaut restent correctes.
   - **Fichier**: [particle/new_test.go](particle/new_test.go)

### 3. **TestCreeParticuleDeuxArguments**
   - **Description**: Teste que les deux premiers arguments définissent `appearance` et `velocityX` correctement.
   - **Fichier**: [particle/new_test.go](particle/new_test.go)

### 4. **TestCreeParticuleTroisArguments**
   - **Description**: Teste l'interprétation des trois premiers arguments (appearance, velocityX, velocityY) et l'âge.
   - **Fichier**: [particle/new_test.go](particle/new_test.go)

### 5. **TestCreeParticuleQuatreArguments**
   - **Description**: Teste la propagation des quatre premiers arguments et le positionnement par défaut (centre).
   - **Fichier**: [particle/new_test.go](particle/new_test.go)

### 6. **TestCreeParticuleCinqArguments**
   - **Description**: Teste que les cinq premiers arguments (incluant ScaleY) sont appliqués correctement.
   - **Fichier**: [particle/new_test.go](particle/new_test.go)

### 7. **TestCreeParticuleSixArguments**
   - **Description**: Teste six arguments (appearance, scaleX, scaleY, color components, etc.) et que l'ordre est respecté.
   - **Fichier**: [particle/new_test.go](particle/new_test.go)

### 8. **TestCreeParticuleSeptArguments**
   - **Description**: Teste sept arguments; ici le premier est `math.Pi` pour tester la gestion d'un float spécial.
   - **Fichier**: [particle/new_test.go](particle/new_test.go)

### 9. **TestCreeParticuleHuitArguments**
   - **Description**: Teste huit arguments et la bonne affectation des coordonnées et paramètres supplémentaires.
   - **Fichier**: [particle/new_test.go](particle/new_test.go)

### 10. **TestCreeParticuleNeufArguments**
   - **Description**: Teste un cas plus complet (neuf arguments) pour s'assurer que l'ordre des paramètres est correct.
   - **Fichier**: [particle/new_test.go](particle/new_test.go)

## Tests du package `particles` - Création du système

### 1. **TestCreeSystemSpawnFixe**
   - **Description**: Teste la création d'un système de particules avec spawn à position fixe. Vérifie que le nombre correct de particules est créé et qu'elles sont toutes positionnées à la même location (SpawnX, SpawnY) avec les bons paramètres initiaux.
   - **Fichier**: [particles/new_test.go](particles/new_test.go)

### 2. **TestCreeSystemSpawnAleatoire**
   - **Description**: Teste la création d'un système de particules avec spawn aléatoire. Vérifie que les positions générées sont bien dans les limites de la fenêtre et que les valeurs d'âge maximum sont correctes.
   - **Fichier**: [particles/new_test.go](particles/new_test.go)

### 3. **TestAccumulateurSpawnFonctionnel**
   - **Description**: Teste que l'accumulateur de spawn accumule correctement le taux de spawn et génère les bonnes particules.
   - **Fichier**: [particles/new_test.go](particles/new_test.go)

## Tests du package `particles` - Mises à jour

### 1. **TestOpaciteDiminue**
   - **Description**: Teste que l'opacité d'une particule diminue au cours du temps. Elle crée une particule avec une opacité de 1.0 et vérifie qu'après l'appel à Update(), l'opacité est inférieure à 1.0.
   - **Fichier**: [particles/new_test.go](particles/new_test.go)

### 2. **TestGraviteAffecteVitesse**
   - **Description**: Teste que la gravité augmente la vélocité verticale des particules.
   - **Fichier**: [particles/new_test.go](particles/new_test.go)

### 3. **TestFrictionRalentitParticule**
   - **Description**: Teste que la friction ralentit les particules.
   - **Fichier**: [particles/new_test.go](particles/new_test.go)

### 4. **TestMiseAJourPosition**
   - **Description**: Teste que la position des particules est correctement mise à jour en fonction de leur vélocité.
   - **Fichier**: [particles/new_test.go](particles/new_test.go)

### 5. **TestParticuleAgeMaxSupprimee**
   - **Description**: Teste que les particules atteignant leur âge maximum sont supprimées du système.
   - **Fichier**: [particles/new_test.go](particles/new_test.go)

### 6. **TestRotationAvance**
   - **Description**: Teste que la rotation des particules augmente à chaque mise à jour.
   - **Fichier**: [particles/new_test.go](particles/new_test.go)

### 7. **TestParticulesHorsEcranSupprimees**
   - **Description**: Teste que les particules sortant de la fenêtre d'affichage sont supprimées du système.
   - **Fichier**: [particles/new_test.go](particles/new_test.go)

## Tests du package `particles` - Avancés (update_test.go)

### 1. **TestMiseAJourGraviteEtFriction**
   - **Description**: Teste que la gravité et la friction sont correctement appliquées à toutes les particules du système.
   - **Fichier**: [particles/update_test.go](particles/update_test.go)

### 2. **TestMiseAJourAgeParticule**
   - **Description**: Teste que l'âge des particules augmente correctement à chaque appel à Update().
   - **Fichier**: [particles/update_test.go](particles/update_test.go)

### 3. **TestMiseAJourOpaciteDecroissante**
   - **Description**: Teste que l'opacité des particules décroît linéairement avec leur âge jusqu'à atteindre 0 quand MaxAge est atteint.
   - **Fichier**: [particles/update_test.go](particles/update_test.go)

### 4. **TestMiseAJourRotationParticule**
   - **Description**: Teste que la rotation des particules augmente correctement selon le RotationSpeed.
   - **Fichier**: [particles/update_test.go](particles/update_test.go)

### 5. **TestMiseAJourSuppressionsMultiples**
   - **Description**: Teste que plusieurs particules expirées sont correctement supprimées du système en une seule Update().
   - **Fichier**: [particles/update_test.go](particles/update_test.go)

### 6. **TestMiseAJourParticulesHorsEcran**
   - **Description**: Teste que plusieurs variantes de positions hors écran sont correctement détectées et supprimées.
   - **Fichier**: [particles/update_test.go](particles/update_test.go)

### 7. **TestMiseAJourAccumulateurSpawn**
   - **Description**: Teste que l'accumulateur de spawn fonctionne correctement et génère les particules au bon moment.
   - **Fichier**: [particles/update_test.go](particles/update_test.go)

## Exécution des tests

Pour exécuter tous les tests:

```bash
# Tests du package particle
go test ./particle -v

# Tests du package particles
go test ./particles -v

# Tous les tests
go test ./... -v
```

## Couverture des tests

- **Création de particules**: 10 tests couvrant différentes combinaisons d'arguments
- **Création du système**: 3 tests couvrant spawn fixe, aléatoire, et accumulateur
- **Mises à jour**: 14 tests couvrant gravité, friction, opacité, rotation, suppression, etc.
- **Total**: 27 tests unitaires

Tous les tests sont en français avec des noms clairs et des commentaires explicatifs pour chaque test.
