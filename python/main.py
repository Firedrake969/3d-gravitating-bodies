bodies = []


class Body:

    def __init__(self, **kwargs):
        self.radius = kwargs['radius']
        self.position = kwargs['position']
        self.velocity = kwargs['velocity']

    def update_velocity(self, bodies):
        new_velocity = self.velocity
        for body in bodies:
            if self is not body:
                # do the math
                # do the monster math
                pass
        self.velocity = new_velocity

    def update_position(self):
        for i in range(0, 3):
            self.position[i] += self.velocity[i]

    def check_collisions(self, bodies):
        for body in bodies:
            if self is not body:
                if self.dist(body) < (self.radius + body.radius):
                    # stuff

    def dist(self, other_body):
        dx = self.position[0] - otherBody.position[0];
        dy = self.position[1] - otherBody.position[1];
        dz = self.position[2] - otherBody.position[2];
        return Math.sqrt( dx * dx + dy * dy + dz * dz );


def cycle():
    for b in bodies:
        b.update_velocity(bodies)
    for b in bodies:
        b.update_position()
    for b in bodies:
        b.check_collisions(bodies)