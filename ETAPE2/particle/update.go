package particle

func (p *Particle) Update(float64) {
	// Mettre à jour la position en fonction de la vélocité
	p.PositionX += p.VelocityX
	p.PositionY += p.VelocityY
}
