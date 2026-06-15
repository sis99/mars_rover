from python_version.models import Rover, Plateau, Direction
import pytest

INVALID_BOUNDARIES = [
    # negative coordinate
    (
        Rover(x_coordinate=-1, y_coordinate=0, direction=Direction.N),
        Plateau(max_x=1, max_y=1),
    ),
    # positive but out of bounds
    (
        Rover(x_coordinate=2, y_coordinate=0, direction=Direction.S),
        Plateau(max_x=1, max_y=1),
    ),
]


@pytest.mark.parametrize("rover,plateau", INVALID_BOUNDARIES)
def test_invalid_boundaries(rover, plateau):
    assert plateau.contains(rover.x_coordinate, rover.y_coordinate) is False


VALID_BOUNDARIES = [
    # can be negative if plateau allows
    (
        Rover(x_coordinate=-1, y_coordinate=-1, direction=Direction.N),
        Plateau(max_x=2, max_y=2, min_x=-1, min_y=-1),
    ),
    # same as plateau's boundary
    (
        Rover(x_coordinate=2, y_coordinate=2, direction=Direction.S),
        Plateau(max_x=2, max_y=2),
    ),
    # inside of plateau
    (
        Rover(x_coordinate=1, y_coordinate=1, direction=Direction.S),
        Plateau(max_x=2, max_y=2),
    ),
]


@pytest.mark.parametrize("rover,plateau", VALID_BOUNDARIES)
def test_valid_boundaries(rover, plateau):
    assert plateau.contains(rover.x_coordinate, rover.y_coordinate) is True
