use std::sync::{Mutex, Once};

struct VoteCount {
    red: u32,
    blue: u32,
}

static mut VOTE_COUNT: Option<Mutex<VoteCount>> = None;
static INIT: Once = Once::new();

fn get_vote_count() -> &'static Mutex<VoteCount> {
    unsafe {
        INIT.call_once(|| {
            VOTE_COUNT = Some(Mutex::new(VoteCount { red: 0, blue: 0 }));
        });
        VOTE_COUNT.as_ref().unwrap()
    }
}

#[no_mangle]
pub extern "C" fn vote_red() {
    let count = get_vote_count();
    let mut vote_count = count.lock().unwrap();
    if vote_count.red + vote_count.blue < 11 {
        vote_count.red += 1;
        println!("Voted for red");
        unsafe{write("red".to_owned())};
    } else {
        println!("Voting limit reached");
    }
}

#[no_mangle]
pub extern "C" fn vote_blue() {
    let count = get_vote_count();
    let mut vote_count = count.lock().unwrap();
    if vote_count.red + vote_count.blue < 11 {
        vote_count.blue += 1;
        println!("Voted for blue");
        unsafe{write("blue".to_owned())};
    } else {
        println!("Voting limit reached");
    }
}

#[no_mangle]
pub extern "C" fn get_vote_counts() -> (u32, u32) {
    let count = get_vote_count();
    let vote_count = count.lock().unwrap();
    (vote_count.red, vote_count.blue)
}


extern "C" {
    fn write(candidate: String);
}
