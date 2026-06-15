class InvalidPlateauSizeError(Exception):
    pass


class InvalidMoveError(Exception):
    pass


class InvalidActionError(Exception):
    pass


class InvalidRoverParamsError(Exception):
    pass


class EmptyFileError(Exception):
    pass


class MissingActionsError(Exception):
    def __init__(self, rover: str):
        super().__init__(f"Missing actions for rover: {rover}")
