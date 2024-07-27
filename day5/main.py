#!/usr/bin/env python


def read_seeds(from_file) -> list:
    """
    Reads a list of seeds from a file.
    """
    first_line = from_file.readline()
    return list(map(int, first_line.strip().split(" ")[1:]))


def print_maps(maps_to_print):
    print(f"maps[{len(maps_to_print)}]:")
    for m in maps_to_print:
        print(m)


def read_data_from_file(file_path) -> tuple[list, list]:
    """
    Reads all the data from input file.
    """
    with open(file_path, "r") as input_file:
        seeds = read_seeds(input_file)
        maps = []
        category_maps = []
        for line in input_file.readlines():
            if line != "\n":
                if line[-2] != ":":
                    category_maps.append(list(map(int, line.split(" "))))
            else:
                if len(category_maps) != 0:
                    maps.append(category_maps)
                    category_maps = []
        maps.append(category_maps)
        return seeds, maps


def part1(file_path) -> int:
    seeds, maps = read_data_from_file(file_path)
    # print(f"seeds[{len(seeds)}] =", seeds)
    # print_maps(maps)
    locations = []
    for seed in seeds:
        next_value = seed
        # print(seed)
        for category_maps in maps:
            for m in category_maps:
                dst, src, length = m
                if (next_value >= src) and (next_value <= src + length - 1):
                    next_value = next_value + (dst - src)
                    break
                # print(dst, src, length, next_value)
        locations.append(next_value)
    return min(locations)


def part2(file_path) -> int:
    seeds, maps = read_data_from_file(file_path)
    seeds = [(seeds[i], seeds[i] + seeds[i + 1] - 1) for i, _ in enumerate(seeds) if i % 2 == 0]
    # print(f"seeds[{len(seeds)}] =", seeds)
    # print_maps(maps)
    locations = []
    for seed in seeds:
        current_values = [seed]
        for category_maps in maps:
            category_result = []
            # print(current_values)
            for value in current_values:
                x1, x2 = value
                done = False
                for m in category_maps:
                    dst, src, length = m
                    y1, y2 = src, src + length - 1
                    delta = dst - src

                    # case 1
                    if (y1 > x2) or (y2 < x1):
                        continue

                    # case 2
                    if (x1 >= y1) and (x2 <= y2):
                        done = True
                        category_result.append((x1 + delta, x2 + delta))
                        break

                    # case 3
                    if (x1 < y1) and (y1 < x2) and (x2 <= y2):
                        done = True
                        current_values.append((x1, y1 - 1))
                        category_result.append((y1 + delta, x2 + delta))
                        continue

                    # case 4
                    if (x1 >= y1) and (x2 > y2) and (y2 > x1):
                        done = True
                        current_values.append((y2 + 1, x2))
                        category_result.append((x1 + delta, y2 + delta))
                        continue

                    # case 5
                    if (x1 < y1) and (x2 > y2):
                        done = True
                        current_values.append((x1, y1 - 1))
                        current_values.append((y2 + 1, x2))
                        category_result.append((y1 + delta, y2 + delta))
                        continue

                if not done:
                    category_result.append(value)

            current_values = list(set(category_result))
            category_result.clear()
        locations.extend(current_values)
    # print("locations=", locations)
    # print(len(locations))
    return min(locations)[0]


if __name__ == "__main__":
    print("\nAOC-2023 Day 5 solution")
    print("\nPart1:")
    print("test1.txt: ", part1("test1.txt"))
    print("test2.txt: ", part1("test2.txt"))
    print("\nPart2:")
    print("test1.txt: ", part2("test1.txt"))
    print("test2.txt: ", part2("test2.txt"))
    print()