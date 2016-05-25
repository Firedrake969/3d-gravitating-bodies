// what should this file be named

//http://www.kevs3d.co.uk/dev/phoria/

class Body {
    constructor(radius, velocity, position) {
        this.radius = radius;
        this.velocity = velocity;
        this.position = position;
    }
    updateVelocity(otherBodies) {
        let newVelocity = this.velocity;
        otherBodies.forEach(otherBody => {
            if (otherBody != this) { // does this work
                // math and magic
            }
        });
        this.velocity = newVelocity;
    }
    updatePosition() {
        for (let i = 0; i < 3; i++) {
            this.position[i] += this.velocity[i];
        }
    }
    checkCollisions(otherBodies) {
        otherBodies.forEach(otherBody => {
            if (otherBody != this) { // see updateVelocity's comment
                dist = this.dist(otherBody);
                if (dist < this.radius + otherBody.radius) {
                    // stuffs
                }
            }
        });
    }
    dist(otherBody) {
        var dx = this.position[0] - otherBody.position[0];
        var dy = this.position[1] - otherBody.position[1];
        var dz = this.position[2] - otherBody.position[2];
        return Math.sqrt( dx * dx + dy * dy + dz * dz );
    }
}

var collection = [];

function cycle(collection) {
    collection.forEach(body => {
        body.updateVelocity(collection);
    });
    collection.forEach(body => {
        body.updatePosition();
    });
    collection.forEach(body => {
        body.checkCollisions(collection);
    });
}