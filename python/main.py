bodies = []


class Body:

    def __init__(self, **kwargs):
        self.radius = kwargs['radius']
        self.position = kwargs['position']
        self.velocity = kwargs['velocity']
        self.mass = kwargs['mass']

    def update_velocity(self, bodies):
        new_velocity = self.velocity
        for otherBody in bodies:
            if self is not otherBody:
                radius = self.dist(otherBody) / 2.0
                r_3 = radius ** 3
                # don't bother with self.mass so we don't have to divide it later to calculate acceleration
                force_base = ((G * otherBody.mass) / r_3)
                force_x = force_base * (self.position[0] - otherBody.position[0])
                force_y = force_base * (self.position[1] - otherBody.position[1])
                force_z = force_base * (self.position[2] - otherBody.position[2])
                new_velocity[0] += force_x
                new_velocity[1] += force_y
                new_velocity[2] += force_z
        self.velocity = new_velocity

    def update_position(self):
        for i in range(0, 3):
            self.position[i] += self.velocity[i]

    def check_collisions(self, bodies):
        for otherBody in bodies:
            if self is not otherBody:
                if self.dist(otherBody) < (self.radius + otherBody.radius):
                    pass

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