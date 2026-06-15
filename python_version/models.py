from enum import Enum
from dataclasses import dataclass

from exceptions import InvalidMoveError, InvalidPlateauSizeError


class Direction(Enum):
    N = "N"
    E = "E"
    S = "S"
    W = "W"

    @classmethod
    def all_directions(cls) -> list[Direction]:
        return list(cls)


class Action(Enum):
    L = "L"
    R = "R"
    M = "M"


@dataclass
class Plateau:
    max_x: int
    max_y: int
    min_x: int = 0
    min_y: int = 0

    def __post_init__(self):
        if self.max_x <= self.min_x and self.max_y <= self.min_y:
            raise InvalidPlateauSizeError

    def contains(self, x: int, y: int) -> bool:
        return self.min_x <= x <= self.max_x and self.min_y <= y <= self.max_y


@dataclass
class Rover:
    x_coordinate: int
    y_coordinate: int
    direction: Direction

    @property
    def coordinates(self) -> str:
        return f"{self.x_coordinate} {self.y_coordinate} {self.direction.value}"

    def turn_right(self):
        all_directions = Direction.all_directions()
        current_index = all_directions.index(self.direction)
        self.direction = all_directions[(current_index + 1) % len(all_directions)]

    def turn_left(self):
        all_directions = Direction.all_directions()
        current_index = all_directions.index(self.direction)
        self.direction = all_directions[(current_index - 1) % len(all_directions)]

    def get_new_coordinates(self) -> tuple[int, int]:
        match self.direction:
            case Direction.N:
                return (self.x_coordinate, self.y_coordinate + 1)
            case Direction.E:
                return (self.x_coordinate + 1, self.y_coordinate)
            case Direction.S:
                return (self.x_coordinate, self.y_coordinate - 1)
            case Direction.W:
                return (self.x_coordinate - 1, self.y_coordinate)
            case _:
                raise ValueError(f"Unknown direction: {self.direction}")

    def move(self, plateau: Plateau):
        new_x, new_y = self.get_new_coordinates()
        if not plateau.contains(new_x, new_y):
            raise InvalidMoveError
        self.x_coordinate = new_x
        self.y_coordinate = new_y

    def execute_action(self, *, action: Action, plateau: Plateau):
        match action:
            case Action.L:
                self.turn_left()
            case Action.R:
                self.turn_right()
            case Action.M:
                self.move(plateau)
