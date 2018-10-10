pub fn series(digits: &str, len: usize) -> Vec<String> {
    let mut v: Vec<String> = Vec::new();

    if len <= digits.len() {
        for i in 0..=digits.len()-len {
            v.push(digits.chars().skip(i).take(len).collect());
        }
    }

    v
}
