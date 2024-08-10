from numba import jit
from time import monotonic


def read_races_from_file(filename) -> list:
    with open(filename) as f:
        data = f.readlines()

    times = list(map(int, data[0].split(":")[1].split()))
    distances = list(map(int, data[1].split(":")[1].split()))

    print(f"Times: {times}")
    print(f"Distances: {distances}")

    races = list(zip(times, distances))

    print(f"Races: {races}")

    return races


def read_single_race_from_file(filename) -> tuple:
    with open(filename) as f:
        data = f.readlines()

    race = (
        int(
            "".join(
                [el for el in list(map(str.strip, data[0].split(":")[1])) if el != ""]
            )
        ),
        int(
            "".join(
                [el for el in list(map(str.strip, data[1].split(":")[1])) if el != ""]
            )
        ),
    )
    print(f"input: {race}")
    return race


@jit(nopython=True)
def solvePart1(races) -> int:
    result = 1
    for race in races:
        success_count = 0
        for speed in range(0, race[0] + 1):
            time = race[0] - speed
            distance = speed * time
            if distance > race[1]:
                success_count = success_count + 1
        result = result * success_count
    return result


@jit(nopython=True)
def solvePart2(race) -> int:
    success_count = 0
    for speed in range(0, race[0] + 1):
        time = race[0] - speed
        distance = speed * time
        if distance > race[1]:
            success_count = success_count + 1
    return success_count


if __name__ == "__main__":
    races = read_races_from_file("test1.txt")
    print(f"part1: {solvePart1(races)}")

    race = read_single_race_from_file("test2.txt")

    start = monotonic()
    print(f"part2: {solvePart2(race)}")
    total = monotonic() - start
    print(f"part2 time: {total:.3f}s")
