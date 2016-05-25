bodies = []

class Body:

    def __init__(self, **kwargs):
        self.radius = kwargs['radius']
        self.position = kwargs['position']
        self.velocity = kwargs['velocity']

    def update_velocity(bodies):
        new_velocity = [0, 0, 0]
        for body in bodies:
            if self is not body:
                # do the math
                # do the monster math
                pass
        self.velocity = new_velocity

    def update_position():
        for i in range(0, 3):
            self.position[i] += self.velocity[i]

def cycle():
    for b in bodies:
        b.update_velocity(bodies)
    for b in bodies:
        b.update_position()