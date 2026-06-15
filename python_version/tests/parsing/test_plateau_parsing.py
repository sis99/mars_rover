from parsing import parse_plateau
from exceptions import InvalidPlateauSizeError
from models import Plateau
import pytest

INVALID_PLATEAUS = [
    # negative coord
    "0 -1",
    # same as min coords
    "0 0",
    # invalid separator
    "0, 1",
    # invalid input
    "ekqnfek",
    # extra spaces
    " 1   1",
]


@pytest.mark.parametrize("line", INVALID_PLATEAUS)
def test_invalid_plateau(line):
    with pytest.raises(InvalidPlateauSizeError):
        parse_plateau(line)


VALID_PLATEAUS = [
    # valid positive values
    "1 1",
    # more than one digit
    "100 200",
    # one dimension at minimum (line-strip plateau)
    "0 1",
    "1 0",
]


@pytest.mark.parametrize("line", VALID_PLATEAUS)
def test_valid_plateau(line):
    parse_plateau(line)


INVALID_COORDS = [
    # min same as max
    (1, 1, 1, 1),
    # max smaller than min
    (-1, 0, 1, 1),
]


@pytest.mark.parametrize("max_x,max_y,min_x,min_y", INVALID_COORDS)
def test_min_cord_greater_than_max(max_x, max_y, min_x, min_y):
    with pytest.raises(InvalidPlateauSizeError):
        Plateau(max_x=max_x, max_y=max_y, min_x=min_x, min_y=min_y)
