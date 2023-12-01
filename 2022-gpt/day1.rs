use std::fs;

fn main() {
    let input = fs::read_to_string("input-day1.txt").expect("Failed to read input file");
    let elves: Vec<Vec<u32>> = input
        .split("\n\n")
        .map(|s| {
            s.lines()
                .map(|line| line.trim().parse::<u32>().unwrap())
                .collect()
        })
        .collect();

    // Part 1: Find the Elf carrying the most Calories
    let max_calories = elves
        .iter()
        .map(|elf| elf.iter().sum::<u32>())
        .max()
        .unwrap();

    println!("Part 1: {}", max_calories);

    // Part 2: Find the total Calories carried by the top three Elves
    let mut calories_per_elf: Vec<u32> = elves.iter().map(|elf| elf.iter().sum()).collect();
    calories_per_elf.sort_unstable_by(|a, b| b.cmp(&a));

    let total_calories_top_three: u32 = calories_per_elf.iter().take(3).sum();

    println!("Part 2: {}", total_calories_top_three);
}

