package particle

func NewParticle(x, y, vx, vy float64) *Particle {

    return &Particle{
        PositionX: x, PositionY: y,
        VelocityX: vx, VelocityY: vy, //VelocityX = vitesse horizontale (vers la droite ou la gauche) et VelocityY = verticale (vers le haut ou le bas).
        Rotation:  0,
        ScaleX:    1, ScaleY: 1,
        ColorRed:  1, ColorGreen: 1, ColorBlue: 1,
        Opacity:   1,
    }
}

