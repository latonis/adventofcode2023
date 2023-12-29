use std::char;
use std::cmp::Ordering;
use std::collections::HashMap;
use std::fs;
use std::env;
use std::iter::zip;
use std::path::Path;
use std::vec;

fn main() {
    let input_path = Path::new(env::current_dir().expect("Where am I?").parent().unwrap()).join("test-input");
    let contents = fs::read_to_string(input_path).expect("Where did my file go?");

    part_one(&contents);
    // part_two(&contents);
}

fn compare(hand1: (u8, &str), hand2: (u8, &str) ) -> Ordering {
    let cards = vec!["2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"];
    if hand1.0 == hand2.0 {
        for (char1, char2) in zip(hand1.1.chars(), hand2.1.chars()) {
            let index1 = cards.iter().position(|&r| r == char1.to_string()).unwrap();
            let index2 = cards.iter().position(|&r| r == char2.to_string()).unwrap();
            if char1 != char2 {
                if index1 > index2 {
                    return Ordering::Greater
                }
                return Ordering::Less
            }
        }
        return Ordering::Greater
    } else if (hand1.0) > (hand2.0) {
        return Ordering::Greater
    }
    return Ordering::Less
}

enum Hand {
    HighCard,
    OnePair,
    TwoPair,
    ThreeOfAKind,
    FullHouse,
    FourOfAKind,
    FiveOfAKind
}

fn part_one(contents: &String) {

    let mut aa = contents.split("\n").map(|line| {
        let letter_counts: HashMap<char, u8> = line.split(" ").nth(0).unwrap().chars().fold(HashMap::new(), |mut map, c| {
            *map.entry(c).or_insert(0) += 1;
            map
        });

        let max = letter_counts.iter().fold(0, |a,b| a.max(*b.1));
        let hand_type = match max {
            a if a == 5 => {
                Hand::FiveOfAKind
            }
            a if a == 4 => {
                Hand::FourOfAKind
            }
            a if a == 3 => {
                if letter_counts.values().collect::<Vec<_>>().contains(2) {
                    Hand::FullHouse
                } else {
                    Hand::ThreeOfAKind
                }
            }
            a if a == 2 => {
                Hand::TwoPair
            }
            _ => {Hand::HighCard}
        }

        (,line.split(" ").nth(0).unwrap() ,line.split(" ").nth(1).unwrap().parse::<u64>().unwrap())
    
    }).collect::<Vec<(u8, &str, u64)>>();

    aa.sort_by(|a, b| {
        compare((a.0, a.1), (b.0, b.1))
    });

    dbg!(aa.iter().fold(0, |acc, a| acc + (a.2 * aa.iter().position(|b| a.1 == b.1).unwrap() as u64 + 1)));
}

fn part_two(contents: &String) {
    println!("Part Two: {contents}");
}