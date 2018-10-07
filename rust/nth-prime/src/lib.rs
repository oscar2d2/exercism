fn is_prime(n: u32) -> bool {
    let upper_limit = (n as f64).sqrt() as u32 + 1;
    for i in 2..upper_limit {
        if n % i == 0 {
            return false;
        }
    }

    true
}

pub fn nth(n: u32) -> u32 {
   if n == 0 {
       return 2;
   }

   let mut count = 0;
   let mut next_prime = 3;

   loop {
       while !is_prime(next_prime) {
           next_prime = next_prime + 2; 
       }

       count += 1;

       if count == n {
           return next_prime;
       }

       next_prime += 2;
    }
}
