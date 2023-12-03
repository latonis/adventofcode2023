use std::collections::HashMap;
use std::env;
use std::fs;
use std::path::Path;

fn main() {
    let input_path =
        Path::new(env::current_dir().expect("Where am I?").parent().unwrap()).join("test-input");
    let contents = fs::read_to_string(input_path).expect("Where did my file go?");

    part_one(&contents);
    // part_two(&contents);
}

#[derive(Debug)]
struct MarbleValue {
    color: String,
    amount: u32,
}

fn parse(line: &str) -> MarbleValue {
    let split_line = line.split(" ").collect::<Vec<_>>();
    
    MarbleValue {
        color: split_line[1].parse::<String>().unwrap(),
        amount: split_line[0].parse::<u32>().unwrap()
    }
}

fn part_one(contents: &String) {
    // let mut valid_ids_sum: u32 = 0;
    let color_max = HashMap::from([("red", 12), ("blue", 14), ("green", 13)]);

    contents.split("\n").for_each(|line| {
        let game_contents =  line.split(":").collect::<Vec<_>>();
        let game_id = game_contents[0]
            .split(" ")
            .collect::<Vec<_>>()[1]
            .parse::<u32>()
            .unwrap();

        game_contents[1].split(";").for_each(|game| {
            dbg!(game.trim().split(",").collect::<Vec<_>>().iter().map(|item| parse(item.trim())).collect::<Vec<_>>());
        });
    // println!("Part One: {contents}");
    });
}

fn part_two(contents: &String) {
    println!("Part Two: {contents}");
}
