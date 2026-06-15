from python_version.parsing import parse_file
from python_version.exceptions import (
    EmptyFileError,
    InvalidPlateauSizeError,
    InvalidRoverParamsError,
    MissingActionsError,
)
from python_version.models import Action

import pytest


def test_empty_file():
    with pytest.raises(EmptyFileError):
        parse_file("tests/parsing/files/empty_file")


def test_extra_spaces_handled():
    plateau, rovers_w_actions = parse_file("tests/parsing/files/extra_spaces")
    expected_max_coordinate = 5
    expected_min_coordinate = 0

    assert plateau.max_x == expected_max_coordinate
    assert plateau.max_y == expected_max_coordinate
    assert plateau.min_x == expected_min_coordinate
    assert plateau.min_y == expected_min_coordinate

    expected_len = 1
    assert len(rovers_w_actions) == expected_len

    expected_coord_str = "1 2 N"
    expected_actions = [Action.L, Action.M]

    rover, actions = rovers_w_actions[0][0], rovers_w_actions[0][1]
    assert rover.coordinates == expected_coord_str
    assert actions == expected_actions


def test_missing_plateau():
    with pytest.raises(InvalidPlateauSizeError):
        parse_file("tests/parsing/files/missing_plateau")


def test_missing_rover():
    with pytest.raises(InvalidRoverParamsError):
        parse_file("tests/parsing/files/missing_rover")


def test_missing_actions():
    with pytest.raises(MissingActionsError):
        parse_file("tests/parsing/files/missing_actions")
