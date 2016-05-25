package main

import (
    "math"
)

type Body struct {
    Radius int          `json:"radius"`
    Position [3]float64 `json:"position"`
    Velocity [3]float64 `json:"velocity"`
    Mass float64        `json:"mass"`
}

type Collection struct {
    Bodies []*Body `json:"bodies"`
}

func (collection *Collection) UpdateVelocity(body *Body) {
    new_velocity := body.Velocity // easier to manipulate
    for _, otherBody := range collection.Bodies {
        if otherBody != body { // check that this works!
            radius := body.Dist(otherBody) / 2.0
            r_3 = math.Pow(radius, 3.0)
            // don't bother with self.mass so we don't have to divide it later to calculate acceleration
            force_base := ((G * otherBody.Mass) / r_3)
            force_x = force_base * (otherBody.Position[0] - body.Position[0])
            force_y = force_base * (otherBody.Position[1] - body.Position[1])
            force_z = force_base * (otherBody.Position[2] - body.Position[2])
            new_velocity[0] += force_x
            new_velocity[1] += force_y
            new_velocity[2] += force_z
        }
    }
    body.Velocity = new_velocity
}

func (collection *Collection) CheckCollisions(body *Body) {
    for _, otherBody := range collection.Bodies {
        if otherBody != body {
            if body.Dist(otherBody) < (body.Radius + otherBody.Radius) {
                // math
            }
        }
    }
}

func (body *Body) Dist(otherBody *Body) float64 {
    dx := body.Position[0] - otherBody.Position[0]
    dy := body.Position[1] - otherBody.Position[1]
    dz := body.Position[2] - otherBody.Position[2]
    return math.Sqrt( dx * dx + dy * dy + dz * dz )
}

func (body *Body) UpdatePosition() {
    body.Position[0] += body.Velocity[0]
    body.Position[1] += body.Velocity[1]
    body.Position[2] += body.Velocity[2]
}

func (collection *Collection) Cycle() {
    for _, body := range collection.Bodies {
        collection.UpdateVelocity(body)
    }

    for _, body := range collection.Bodies {
        body.UpdatePosition()
    }

    for _, body := range collection.Bodies {
        collection.CheckCollisions(body)
    }
}