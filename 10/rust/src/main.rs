use std::fs;
use std::env;
use std::path::Path;

fn main() {
    let input_path = Path::new(env::current_dir().expect("Where am I?").parent().unwrap()).join("input");
    let contents = fs::read_to_string(input_path).expect("Where did my file go?");

    part_one(&contents);
    part_two(&contents);
}

fn part_one(contents: &String) {
    println!("Part One: {contents}");
}

fn part_two(contents: &String) {
    println!("Part Two: {contents}");
}