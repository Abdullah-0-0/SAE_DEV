## Modification dynamique du système de particules

### Fonctionnalités implémentées :

#### 1. **Interface de modification des paramètres** ([draw.go](draw.go))
   - Affiche les trois paramètres modifiables : **Gravité**, **Friction**, et **SpawnRate**
   - Affiche un indicateur `→` pour montrer le paramètre actuellement sélectionné
   - Montre les valeurs actuelles avec 4 décimales pour la précision
   - Inclut des instructions d'utilisation

#### 2. **Gestion des entrées clavier** ([update.go](update.go))
   - **Flèches ↑/↓** : Sélectionner le paramètre précédent/suivant (navigation cyclique)
   - **Flèches ←/→** : Diminuer/augmenter la valeur du paramètre sélectionné (delta = 0.01)
   - Validation des valeurs :
     - **Gravité** : Doit être ≥ 0
     - **Friction** : Doit être entre 0 et 1
     - **SpawnRate** : Doit être ≥ 0

#### 3. **Structure de données** ([game.go](game.go))
   - Ajout du champ `paramIndex` à la structure `game` pour tracker le paramètre sélectionné
   - Constantes énumérées pour les paramètres (ParamGravity, ParamFriction, ParamSpawnRate, NumParams)

#### 4. **Fichiers créés/modifiés** :
   - **[colors.go](colors.go)** - Nouvelles : Définit les couleurs pour l'interface (blanc et gris)
   - **[draw.go](draw.go)** - Modifié : Ajout de `drawParameterInterface()` et affichage de l'interface
   - **[update.go](update.go)** - Modifié : Gestion complète des entrées clavier et modifications de paramètres
   - **[game.go](game.go)** - Modifié : Ajout de `paramIndex` et constantes énumérées
   - **[main.go](main.go)** - Modifié : Initialisation de `paramIndex`

### Utilisation :
1. Lancer le programme : `.\project-particles.exe`
2. Utiliser les **flèches du clavier** pour naviguer et modifier les paramètres en temps réel
3. Les modifications s'appliquent immédiatement au système de particules

### Points clés de l'implémentation :
- Utilisation de `inpututil.IsKeyJustPressed()` pour la navigation (évite la répétition rapide)
- Utilisation de `ebiten.IsKeyPressed()` pour la modification (permet l'accélération continue)
- Validation des valeurs pour garantir la cohérence du système
- Format d'affichage avec 4 décimales pour la précision
