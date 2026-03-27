## Un système de particules

### Étapes de la création d'un système de particules :

#### 1. **Les fichiers [particle.go] et [particles.go]** 

   D'abord, nous avons décidé de séparer le fichier qui créé la particule du fichier qui contrôle
   l'ensemble des particules.
    
   - [particle.go] `→` Crée une particule avec des caractéristiques initialisées. Ces caractéristiques sont aussi
     définies par défaut. Si toutes les valeurs ne sont pas définies, alors le programme attribue à la particule une
     valeur aléatoire.

   - [particles.go] `→` Génère un nombre de particules à des positions aléatoires. Ce nombre est
     défini dans le fichier [config.json].
   
   Chaque fichier a un fichier de test qui lui est attribué.

#### 2. **Les extensions réalisées** 

   - La gravité  
   - La friction 
   - La vélocité
   - Les particules sont supprimées lorsqu'elles sortent de l'écran
   - Il est possible d'interagir avec le système de particules. C'est-à-dire qu'il est possible de modifier la
     friction, la gravité, la rotation, etc. 

#### 3. **Structure de données** ([game.go](game.go))
  
   - Ajout du champ `paramIndex` à la structure `game` pour tracker le paramètre sélectionné
   - Constantes énumérées pour les paramètres (ParamGravity, ParamFriction, ParamSpawnRate, NumParams)

#### 4. **Fichiers créés/modifiés** :
   - **[colors.go](colors.go)** -
   - **[draw.go](draw.go)** - 
   - **[update.go](update.go)** -
   - **[game.go](game.go)** - 
   - **[main.go](main.go)** - 

### Utilisation :
1. Lancer le programme : `.\project-particles.exe`
2. Utiliser les **flèches du clavier** pour naviguer et modifier les paramètres en temps réel
3. Les modifications s'appliquent immédiatement au système de particules

### Points clés de l'implémentation :
- Utilisation de `inpututil.IsKeyJustPressed()` pour la navigation (évite la répétition rapide)
- Utilisation de `ebiten.IsKeyPressed()` pour la modification (permet l'accélération continue)
- Validation des valeurs pour garantir la cohérence du système
- Format d'affichage avec 4 décimales pour la précision
