// what should this file be named

class Body {
    constructor(radius, velocity, position) {
        this.radius = radius;
        this.velocity = velocity;
        this.position = position;
    }
    updateVelocity(otherBodies) {
        let newVelocity = [0, 0, 0];
        otherBodies.forEach(otherBody => {
            if (otherBody != this) { // does this work
                // math and magid
            }
        });
        this.velocity = newVelocity;
    }
    updatePosition() {
        for (let i = 0; i < 3; i++) {
            this.position[i] += this.velocity[i];
        }
    }
}

var collection = [];

function cycle(collection) {
    collection.forEach(body => {
        body.UpdateVelocity(collection);
    });
    collection.forEach(body => {
        body.updatePosition();
    });
}