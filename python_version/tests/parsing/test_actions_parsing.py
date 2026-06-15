from parsing import parse_actions
from exceptions import InvalidActionError
from models import Action

import pytest

INVALID_ACTIONS = [
    # valid lowercase
    "l",
    # invalid input
    "SSSS",
    # bad action with good actions
    "LMRT",
    # extra space
    "LM M",
]


@pytest.mark.parametrize("line", INVALID_ACTIONS)
def test_invalid_actions(line):
    with pytest.raises(InvalidActionError):
        parse_actions(line)


VALID_ACTIONS = [
    # single action
    ("R", [Action.R]),
    # many distinct
    ("RL", [Action.R, Action.L]),
    # repeat action
    ("RRRR", [Action.R, Action.R, Action.R, Action.R]),
]


@pytest.mark.parametrize("line,expected_actions", VALID_ACTIONS)
def test_valid_actions(line, expected_actions):
    actions = parse_actions(line)
    assert actions == expected_actions
