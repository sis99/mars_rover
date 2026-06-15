from python_version.parsing import extract_rover
from python_version.exceptions import InvalidRoverParamsError
from python_version.models import Plateau, Direction
import pytest

INVALID_ROVERS = [
    # negative coord
    "0 -1 N",
    # invalid direction
    "0 0 P",
    # invalid separator
    "0, 1; N",
    # invalid input
    "ekqnfek",
    # extra spaces
    " 0  -1  N  ",
]


@pytest.mark.parametrize("line", INVALID_ROVERS)
def test_invalid_rover(line):
    with pytest.raises(InvalidRoverParamsError):
        plateau = Plateau(max_x=1, max_y=1)
        extract_rover(plateau=plateau, line=line)


VALID_ROVERS = [
    # negative coord if plateau allows
    "0 -1 N",
    # more than 1 digit
    "11 10 S",
    # single digit
    "1 1 W",
]


@pytest.mark.parametrize("line", VALID_ROVERS)
def test_valid_rover(line):
    plateau = Plateau(max_x=20, max_y=20, min_x=-1, min_y=-1)
    rover = extract_rover(plateau=plateau, line=line)

    x_coordinate, y_coordinate, direction = line.split(" ")
    assert rover.x_coordinate == int(x_coordinate)
    assert rover.y_coordinate == int(y_coordinate)
    assert rover.direction == Direction(direction)
    assert rover.coordinates == f"{x_coordinate} {y_coordinate} {direction}"
