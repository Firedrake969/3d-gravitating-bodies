// what should this file be named

//http://www.kevs3d.co.uk/dev/phoria/

class Body {
    constructor(radius, velocity, position, mass) {
        this.radius = radius;
        this.velocity = velocity;
        this.position = position;
        this.mass = mass;
    }
    updateVelocity(otherBodies) {
        let newVelocity = this.velocity;
        otherBodies.forEach(otherBody => {
            if (otherBody != this) { // does this work
                let radius = this.dist(otherBody) / 2
                let r_3 = Math.pow(radius, 3)
                // don't bother with self.mass so we don't have to divide it later to calculate acceleration
                let force_base = ((G * otherBody.mass) / r_3)
                let force_x = force_base * (this.position[0] - otherBody.position[0])
                let force_y = force_base * (this.position[1] - otherBody.position[1])
                let force_z = force_base * (this.position[2] - otherBody.position[2])
                new_velocity[0] += force_x
                new_velocity[1] += force_y
                new_velocity[2] += force_z
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
        let dx = this.position[0] - otherBody.position[0];
        let dy = this.position[1] - otherBody.position[1];
        let dz = this.position[2] - otherBody.position[2];
        return Math.sqrt( dx * dx + dy * dy + dz * dz );
    }
}

let collection = [];

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