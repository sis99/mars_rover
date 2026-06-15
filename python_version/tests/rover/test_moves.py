from python_version.models import Direction, Rover, Plateau
from python_version.exceptions import InvalidMoveError
import pytest

RIGHT_TURN = {
    Direction.N: Direction.E,
    Direction.E: Direction.S,
    Direction.S: Direction.W,
    Direction.W: Direction.N,
}
LEFT_TURN = {
    Direction.N: Direction.W,
    Direction.W: Direction.S,
    Direction.S: Direction.E,
    Direction.E: Direction.N,
}

ROVER_DIRECTIONS = [Direction.N, Direction.E, Direction.S, Direction.W]

ROVERS_AFTER_MOVE = [
    (Direction.N, {"x_coordinate": 1, "y_coordinate": 2}),
    (Direction.E, {"x_coordinate": 2, "y_coordinate": 1}),
    (Direction.S, {"x_coordinate": 1, "y_coordinate": 0}),
    (Direction.W, {"x_coordinate": 0, "y_coordinate": 1}),
]


@pytest.mark.parametrize("direction", ROVER_DIRECTIONS)
def test_turn_right(direction):
    rover = Rover(x_coordinate=1, y_coordinate=1, direction=direction)
    rover.turn_right()
    assert rover.direction == RIGHT_TURN[direction]


@pytest.mark.parametrize("direction", ROVER_DIRECTIONS)
def test_turn_left(direction):
    rover = Rover(x_coordinate=1, y_coordinate=1, direction=direction)
    rover.turn_left()
    assert rover.direction == LEFT_TURN[direction]


@pytest.mark.parametrize("direction, new_coords", ROVERS_AFTER_MOVE)
def test_move(direction, new_coords):
    rover = Rover(x_coordinate=1, y_coordinate=1, direction=direction)
    plateau = Plateau(max_x=2, max_y=2)
    rover.move(plateau=plateau)
    assert rover.x_coordinate == new_coords["x_coordinate"]
    assert rover.y_coordinate == new_coords["y_coordinate"]


def test_out_of_bounds_move():
    plateau = Plateau(max_x=1, max_y=1)
    rover = Rover(x_coordinate=1, y_coordinate=1, direction=Direction.N)

    with pytest.raises(InvalidMoveError):
        rover.move(plateau=plateau)
