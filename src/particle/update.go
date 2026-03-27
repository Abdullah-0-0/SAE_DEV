package particle


func (p *Particle) Update() {
    
    p.PositionX += p.VelocityX
    p.PositionY += p.VelocityY
}


   