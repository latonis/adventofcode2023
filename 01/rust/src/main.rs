use std::fs;
use std::env;
use std::path::Path;

fn main() {
    let input_path = Path::new(env::current_dir().expect("Where am I?").parent().unwrap()).join("input");
    let contents = fs::read_to_string(input_path).expect("Where did my file go?");

    part_one(&contents);
    // part_two(&contents);
}

fn part_one(contents: &String) {
    let mut total_val: u32 = 0;
    contents.split("\n").for_each(|line| {
        let res: Vec<char> = line.chars().filter(|c| c.is_numeric()).collect();
        total_val += format!("{}{}", res.first().unwrap(), res.last().unwrap()).to_string().parse::<u32>().unwrap()
    });

    println!("{}", total_val);
}

fn part_two(contents: &String) {
    println!("Part Two: {contents}");
}