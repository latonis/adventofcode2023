use std::fs;
use std::env;
use std::path::Path;
use std::collections::HashMap;

fn main() {
    let input_path = Path::new(env::current_dir().expect("Where am I?").parent().unwrap()).join("input");
    let contents = fs::read_to_string(input_path).expect("Where did my file go?");

    part_one(&contents);
    part_two(&contents);
}

fn part_one(contents: &String) {
    let mut total_val: u32 = 0;
    contents.split("\n").for_each(|line| {
        let res = line.chars().filter(|c| c.is_numeric()).collect::<Vec<char>>();
        total_val += format!("{}{}", res.first().unwrap(), res.last().unwrap()).to_string().parse::<u32>().unwrap()
    });

    println!("{}", total_val);
}

fn part_two(contents: &String) {
    let mut total_val: u32 = 0;
    let num_string_map = HashMap::from([
        ("one", 1),
        ("two", 2),
        ("three", 3),
        ("four", 4),
        ("five", 5),
        ("six", 6),
        ("seven", 7),
        ("eight", 8),
        ("nine", 9)
    ]);
    contents.split("\n").for_each(|line| {
        let mut nums = Vec::<u32>::new();
        for (idx, c) in line.chars().enumerate() {
            if c.is_numeric()  {
                nums.push(c.to_string().parse::<u32>().unwrap());
            } else {
                let fake_string = line.chars().collect::<Vec<char>>();        
                for i in 2..6 {
                    if idx + i < fake_string.len() {
                        let ltr_check: String = fake_string[idx..idx+i+1].into_iter().collect();
                        if num_string_map.contains_key(&ltr_check.as_str()){
                            nums.push(num_string_map[&ltr_check.as_str()].clone())
                        }
                    }
                }
            }    
        }
        total_val += format!("{}{}", nums.first().unwrap(), nums.last().unwrap()).to_string().parse::<u32>().unwrap();
    });
    println!("{}", total_val);

}