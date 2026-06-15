from main import main


def test_full_flow():
    assert main("tests/rover/files/good_basic_case") == "1 3 N\n5 1 E"
