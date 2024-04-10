from manimlib.imports import *
import math
import random


def distance(p1, p2):
    x1 = p1[0]
    y1 = p1[1]
    x2 = p2[0]
    y2 = p2[1]

    d = math.sqrt((x1-x2)**2+(y1-y2)**2)
    return d


def is_near(move_dot, points):
    for point in points:
        if distance(move_dot, point) < 2:
            return True

    return False


def over_boundary(move_dot):
    x = move_dot[0]
    y = move_dot[1]

    if x < 0 or x > 100 or y < 0 or y > 100:
        return True
    else:
        return False


class Scene1(Scene):
    def construct(self):
        points = [(50, 50)]
        for n in range(4000):
            r = 100*random.random()
            move_dot = random.choice([(0, r), (r, 0), (100, r), (r, 100)])
            o_x = 2*random.normalvariate(0, 0.1)
            o_y = 2*random.normalvariate(0, 0.1)

            while not over_boundary(move_dot):
                if is_near(move_dot, points):
                    points.append(move_dot)
                    posX = move_dot[0]/100*8-4
                    posY = move_dot[1]/100*8-4
                    square1 = Square(side_length=0.1, fill_color=RED,
                                     fill_opacity=1.0, stroke_opacity=0.0).move_to(posX*RIGHT+posY*UP)
                    self.add(square1)
                    self.wait(0.1)
                    break
                else:
                    x = move_dot[0]+o_x
                    y = move_dot[1]+o_y
                    move_dot = (x, y)

            if over_boundary(move_dot):
                continue

        self.wait(2)