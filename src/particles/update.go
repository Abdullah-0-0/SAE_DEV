package particles

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.

func (s *System) Update() {
	
	// parcourt toutes les particules dans la liste
	for e := s.Content.Front(); e != nil; e = e.Next() {

		// On récupère la particule actuelle
		p, ok := e.Value.(*Particle)
		if ok {
			
			p.PositionX += p.VelocityX // déplacement horizontal
            p.PositionY += p.VelocityY // déplacement vertical
		}
	}
}

