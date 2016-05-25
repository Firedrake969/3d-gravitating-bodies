package main

import (
    "math"
)

type Body struct {
    Radius int          `json:"radius"`
    Position [3]float64 `json:"position"`
    Velocity [3]float64 `json:"velocity"`
}

type Collection struct {
    Bodies []*Body `json:"bodies"`
}

func (collection *Collection) UpdateVelocity(body *Body) {
    force := body.Velocity // easier to manipulate
    for _, otherBody := range collection.Bodies {
        if otherBody != body { // check that this works!
            // do the magic stuff with MATH!  AMAZING MATH!
        }
    }
    body.Velocity = force
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