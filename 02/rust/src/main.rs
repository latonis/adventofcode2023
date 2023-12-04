use std::collections::HashMap;
use std::env;
use std::fs;
use std::path::Path;

fn main() {
    let input_path =
        Path::new(env::current_dir().expect("Where am I?").parent().unwrap()).join("input");
    let contents = fs::read_to_string(input_path).expect("Where did my file go?");

    part_one(&contents);
    // part_two(&contents);
}

#[derive(Copy, Clone, Debug, Eq, PartialEq, Hash)]
enum Color {
    Red,
    Green,
    Blue,
    Unknown,
}

#[derive(Copy, Debug, Clone)]
struct MarbleValue {
    color: Color,
    amount: u32,
}

#[derive(Debug, Clone)]
struct MarbleGame {
    marbles: Vec<MarbleValue>,
}

fn parse(line: &str) -> MarbleValue {
    let split_line = line.split(" ").collect::<Vec<_>>();
    let color = match split_line[1].parse::<String>().unwrap().as_str() {
        "red" => Color::Red,
        "green" => Color::Green,
        "blue" => Color::Blue,
        _ => Color::Unknown,
    };

    MarbleValue {
        color,
        amount: split_line[0].parse::<u32>().unwrap(),
    }
}

fn part_one(contents: &String) {
    let mut total = 0;
    let color_max = HashMap::from([(Color::Red, 12), (Color::Blue, 14), (Color::Green, 13)]);
    let mut game: HashMap<u32, MarbleGame> = HashMap::new();

    contents.split("\n").for_each(|line| {
        let game_contents = line.split(":").collect::<Vec<_>>();
        let game_id = game_contents[0].split(" ").collect::<Vec<_>>()[1]
            .parse::<u32>()
            .unwrap();

        let marble_games: Vec<_> = game_contents[1]
            .split(";")
            .map(|game| {
                game.trim()
                    .split(",")
                    .collect::<Vec<_>>()
                    .iter()
                    .map(|item| parse(item.trim()))
                    .collect::<Vec<_>>()
            })
            .collect::<Vec<_>>();

        game.insert(
            game_id,
            MarbleGame {
                marbles: marble_games.iter().cloned().flatten().collect(),
            },
        );
    });

    total += game
        .into_iter()
        .filter(|(_, marble_game)| {
            let mut valid = true;
            for marble in &marble_game.marbles {
                if marble.amount > *color_max.get(&marble.color).unwrap() {
                    valid = false;
                }
            }
            valid
        })
        .map(|(id, _)| id)
        .fold(0, |sum, item| sum + item);

    println!("{}", total);
}

fn part_two(contents: &String) {
    println!("Part Two: {contents}");
}
