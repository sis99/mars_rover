from re import compile
from pathlib import Path
from typing import TextIO

from exceptions import (
    EmptyFileError,
    InvalidActionError,
    InvalidPlateauSizeError,
    InvalidRoverParamsError,
    MissingActionsError,
)
from models import Direction, Action, Plateau, Rover

# matches two non-negative integers separated by a space e.g. "5 5"
PLATEAU_RE = compile(r"^\d+ \d+$")
# matches two integers and a cardinal direction separated by spaces e.g. "1 2 N"
ROVER_RE = compile(rf"^-?\d+ -?\d+ ({'|'.join(d.value for d in Direction)})$")


def parse_plateau(line: str) -> Plateau:
    if not line:
        raise EmptyFileError

    if not PLATEAU_RE.fullmatch(line):
        raise InvalidPlateauSizeError

    x, y = line.strip().split(" ")
    return Plateau(max_x=int(x), max_y=int(y))


def parse_rover(line: str) -> Rover:
    if not ROVER_RE.fullmatch(line):
        raise InvalidRoverParamsError

    x, y, direction = line.strip().split(" ")
    return Rover(
        x_coordinate=int(x),
        y_coordinate=int(y),
        direction=Direction(direction),
    )


def parse_actions(line: str) -> list[Action]:
    actions = []
    for character in line.strip():
        try:
            actions.append(Action(character))
        except ValueError as e:
            raise InvalidActionError from e
    return actions


def extract_rover(*, plateau: Plateau, line: str) -> Rover:
    rover = parse_rover(line)

    if not plateau.contains(rover.x_coordinate, rover.y_coordinate):
        raise InvalidRoverParamsError

    return rover


def iter_non_empty_lines(f: TextIO):
    for line in f:
        if stripped := line.strip():
            yield stripped


def parse_stream(f: TextIO) -> tuple[Plateau, list[tuple[Rover, list[Action]]]]:
    lines = iter_non_empty_lines(f)
    plateau = parse_plateau(next(lines, ""))

    rovers_w_actions = []

    for line in lines:
        rover = extract_rover(plateau=plateau, line=line)

        actions_definition = next(lines, "")
        if not actions_definition:
            raise MissingActionsError(str(rover))

        actions = parse_actions(actions_definition)
        rovers_w_actions.append((rover, actions))

    return plateau, rovers_w_actions


def parse_file(path: str | Path) -> tuple[Plateau, list[tuple[Rover, list[Action]]]]:
    with open(Path(path)) as f:
        return parse_stream(f)
