package main

type Body struct {
    Radius int          `json:"radius"`
    Position [3]float64 `json:"position"`
    Velocity [3]float64 `json:"velocity"`
}

type Collection struct {
    Bodies []*Body `json:"bodies"`
}

func (collection *Collection) UpdateVelocity(body *Body) {
    force := [3]float64{0, 0, 0} // should this be like 0.0s?
    for _, otherBody := range collection.Bodies {
        if otherBody != body { // check that this works!
            // do the magic stuff with MATH!  AMAZING MATH!
        }
    }
}

func (body *Body) UpdatePosition() {
    body.Posiion[0] += body.Velocity[0]
    body.Posiion[1] += body.Velocity[1]
    body.Posiion[2] += body.Velocity[2]
}

func (collection *Collection) Cycle() {
    for _, body := collection.Bodies {
        collection.UpdateVelocity(body)
    }

    for _, body := collection.Bodies {
        body.UpdatePosition()
    }
}