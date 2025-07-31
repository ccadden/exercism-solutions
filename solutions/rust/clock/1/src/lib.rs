#[derive(PartialEq, Debug)]
pub struct Clock {
    hours: i32,
    minutes: i32,
}

const MINUTES_IN_HOURS: i32 = 60;
const HOURS_IN_DAY: i32 = 24;

impl std::fmt::Display for Clock {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "{:0>2}:{:0>2}", self.hours, self.minutes)
    }
}

impl Clock {
    pub fn new(mut hours: i32, mut minutes: i32) -> Self {
        if minutes < 0 {
            hours = hours - 1 + (minutes / MINUTES_IN_HOURS);
            minutes = MINUTES_IN_HOURS + (minutes % MINUTES_IN_HOURS);
        }

        if hours < 0 {
            hours = HOURS_IN_DAY + (hours % HOURS_IN_DAY);
        }

        if minutes >= MINUTES_IN_HOURS {
            let additional_hours = minutes / MINUTES_IN_HOURS;
            minutes = minutes % MINUTES_IN_HOURS;
            hours = additional_hours + hours;
        }

        if hours >= HOURS_IN_DAY {
            hours = hours % HOURS_IN_DAY;
        }

        return Clock { hours, minutes };
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        Clock::new(self.hours, self.minutes + minutes)
    }
}
