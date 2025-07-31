// The code below is a stub. Just enough to satisfy the compiler.
// In order to pass the tests you can add-to or change any of this code.

#[derive(Debug)]
pub struct Duration {
    seconds: u64,
}

const EARTH_SECONDS_PER_YEAR: f64 = 31557600.0;

impl From<u64> for Duration {
    fn from(s: u64) -> Self {
        return Duration { seconds: s };
    }
}

pub trait Planet {
    const EARTH_YEAR_MODIFIER: f64;

    fn years_during(d: &Duration) -> f64 {
        d.seconds as f64 / EARTH_SECONDS_PER_YEAR / Self::EARTH_YEAR_MODIFIER
    }
}

pub struct Mercury;
pub struct Venus;
pub struct Earth;
pub struct Mars;
pub struct Jupiter;
pub struct Saturn;
pub struct Uranus;
pub struct Neptune;

macro_rules! init_earth_year_modifiers {
    ($($t:ident => $e:expr), +) => {
        $(
        impl Planet for $t {
        const EARTH_YEAR_MODIFIER: f64 = $e;
        }
        )*
    }
}

init_earth_year_modifiers! {
    Mercury => 0.2408467,
    Venus => 0.61519726,
    Earth => 1.0,
    Mars => 1.8808158,
    Jupiter => 11.862615,
    Saturn => 29.447498,
    Uranus => 84.016846,
    Neptune =>164.79132
}
