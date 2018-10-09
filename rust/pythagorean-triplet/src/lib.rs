pub fn find() -> Option<u32> {
    for a in 1u32..=1000 {
        for b in 1u32..=1000 {
            for c in 1u32..=1000 {
                if a.pow(2) + b.pow(2) == c.pow(2) && a + b + c == 1000{
                    return Some(a*b*c);
                }
            }
        }
    }

    None
}
