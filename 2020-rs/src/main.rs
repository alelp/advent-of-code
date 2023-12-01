use std::env;
use std::fs::File;
use std::io::{self, BufRead, BufReader};

fn main() {
    let args: Vec<String> = env::args().collect();
    let args_count = args.len();

    let days = &args[1..args_count];

    for x in days {
        match x.as_str() {
            "1" => day1(x),
            "2" => day2(x),
            "3" => day3(x),
            "4" => day4(x),
            _ => println!("Invalid day {}", x),
        }
    }
}

fn print_result_check(day: &str, part1_detail:&str, part1_value:i32, part1_excpected_value:i32, part2_detail:&str, part2_value:i32, part2_excpected_value:i32) {
    println!("---------------------------------------------------------");
    println!("* Day {}", day);
    println!("\tPart 1 | {}: {} - Pass: {} (expected {})", part1_detail, part1_value, part1_value == part1_excpected_value, part1_excpected_value);
    println!("\tPart 2 | {}: {} - Pass: {} (expected {})", part2_detail, part2_value, part2_value == part2_excpected_value, part2_excpected_value);
}

fn print_result(day: &str, part1_detail:&str, part1_value:i32, part2_detail:&str, part2_value:i32) {
    println!("---------------------------------------------------------");
    println!("* Day {}", day);
    println!("\tPart 1 | {}: {}", part1_detail, part1_value);
    println!("\tPart 2 | {}: {}", part2_detail, part2_value);
}

fn read_lines(day: &str) -> io::Lines<BufReader<File>> {
    let filename = format!("resources/input-day{}.txt", day);
    let file = File::open(filename).unwrap(); 
    return io::BufReader::new(file).lines(); 
}

fn day1(day: &str) {
    let numbers: Vec<i32> = read_lines(day)
        .map(|line| line.unwrap().parse().unwrap())
        .collect();

    let mut value1 = -1;
    let mut value2 = -1;

    for x in &numbers {
        for y in &numbers {
            if value1 < 0 && x + y == 2020 {
                value1 = x * y;
            }

            if value2 < 0 {
                for z in &numbers {
                    if x + y + z == 2020 {
                        value2 = x * y * z;
                        break;
                    }
                }
            }

            if value1 > 0 && value2 > 0 {
                break;
            }
        }
    }

    print_result(day, "Check 1", value1, "Check 2", value2);
}

fn day2(day: &str) {
    let lines = read_lines(day);

    let mut check1_count = 0;
    let mut check2_count = 0;
    for line in lines {
        let line = line.unwrap();
        let password_data: Vec<&str> = line.split(' ').collect();
        let range: Vec<&str> = password_data[0].split('-').collect();
        let pos1 = range[0].parse::<usize>().unwrap();
        let pos2 = range[1].parse::<usize>().unwrap();
        let key_char = password_data[1].chars().nth(0).unwrap();
        let password = password_data[2..].join(" ");

        let mut key_found_count = 0;
        for char in password.chars() {
            if char == key_char {
                key_found_count += 1;
            }
        }

        if pos1 <= key_found_count && key_found_count <= pos2 {
            check1_count += 1;
        }

        if password.chars().nth(pos1-1) == Some(key_char) && password.chars().nth(pos2-1) != Some(key_char) {
            check2_count += 1;
        }
        else if password.chars().nth(pos1-1) != Some(key_char) && password.chars().nth(pos2-1) == Some(key_char) {
            check2_count += 1;
        }
    }

    print_result(day, "Valid Passwords for check 1", check1_count, "Valid Passwords for check 2", check2_count);
}

fn day3(day:&str) {
    #[derive(Copy, Clone)]
    struct SlopeDefinition {
        right: usize,
        down: usize,
    }

    #[derive(Copy, Clone)]
    struct SlopeTravel {
        current_column:usize,
        trees_found:usize,
    }

    let lines = read_lines(day);

    const PART1_SLOPE_INDEX:usize = 1;
    const SLOPE_COUNT:usize = 5;
    const SLOPES_DEFINITIONS: [SlopeDefinition; SLOPE_COUNT] = [
        SlopeDefinition { right: 1, down: 1 },
        SlopeDefinition { right: 3, down: 1 }, // (This is the slope you already checked.)
        SlopeDefinition { right: 5, down: 1 },
        SlopeDefinition { right: 7, down: 1 },
        SlopeDefinition { right: 1, down: 2 },
    ];


    let mut slope_travels: [SlopeTravel; SLOPE_COUNT] = [SlopeTravel{ current_column:0, trees_found:0 }; SLOPE_COUNT];

    for (row_index, row_data) in lines.enumerate() {
        let row_data = row_data.unwrap();

        for slope_index in 0..SLOPES_DEFINITIONS.len() {
            let slope:&SlopeDefinition = &SLOPES_DEFINITIONS[slope_index];
            let travel:&mut SlopeTravel = &mut slope_travels[slope_index];

            if row_index % (slope.down as usize) != 0 {
                continue;
            }

            if row_data.chars().nth(travel.current_column) == Some('#') {
                travel.trees_found = travel.trees_found +1;
            }

            travel.current_column += slope.right;
            travel.current_column %= row_data.len();
        }
    }

    let mut accum:usize = 1;
    for travel in slope_travels {
        accum *= travel.trees_found;
    }

    print_result(day,
                 "Trees in slope 3-right, 1-down", slope_travels[PART1_SLOPE_INDEX].trees_found as i32,
                 "multiplied found trees", accum as i32)
}

fn day4(day:&str) {
    let part1_detail:&str = "valid documents";
    let mut part1_value:i32 = 0;
    let part1_excpected_value:i32 = 204;

    let part2_detail:&str = "moco";
    let part2_value:i32 = 0;
    let part2_excpected_value:i32 = 179;

    fn extract_value(f:&str, s:&str) -> &str {
        let result = "";
        return result;
    }

    let required_fields = [
         // (Birth Year)
        ("byr:", |f:&str, s:&str| {
            let y:i32 = extract_value(f, s).parse().unwrap();
            return 1920 <= y && y <= 2002;
        }),
        // (Issue Year)
        ("iyr:", |f:&str, s:&str| {
            return false;
        }),
        /*
        // (Expiration Year)
        ("eyr:", |f:&str, s:&str| {
            return false;
        }),
        // (Height)
        ("hgt:", |f:&str, s:&str| {
            return false;
        }),
        // (Hair Color)
        ("hcl:", |f:&str, s:&str| {
            return false;
        }),
        // (Eye Color)
        ("ecl:", |f:&str, s:&str| {
            return false;
        }),
        // (Passport ID)
        ("pid:", |f:&str, s:&str| {
            return false;
        }),
        */
        //"cid:", // (Country ID) // we're making it optional for the 1st pass
    ];

    let lines = read_lines(day).map(|l| l.unwrap()).collect::<Vec<String>>();
    let line_count = lines.len();
    let mut passport_data = "".to_owned();

    for (index, line) in lines.into_iter().enumerate() {

        passport_data = format!("{} {}", passport_data, line);

        if line.len() == 0 || index == line_count - 1 {
            if required_fields.iter().all(|field_data| passport_data.contains(field_data.0))
            {
                part1_value += 1;

                if required_fields.iter().all(|field_data| passport_data.contains(field_data.0))
                {
                    part2_value += 1;
                }
            }
            passport_data = "".to_owned();
        }
    }

    print_result_check(day, part1_detail, part1_value, part1_excpected_value, part2_detail, part2_value, part2_excpected_value);
}

