import argparse
from parsing import parse_file


def main(path: str) -> str:
    plateau, rovers_w_actions = parse_file(path)

    for rover, actions in rovers_w_actions:
        for action in actions:
            rover.execute_action(action=action, plateau=plateau)

    return "\n".join([rover.coordinates for rover, _ in rovers_w_actions])


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Mars Rover navigator")
    parser.add_argument("path", help="Path to the rover instructions file")
    args = parser.parse_args()
    print(main(args.path))
