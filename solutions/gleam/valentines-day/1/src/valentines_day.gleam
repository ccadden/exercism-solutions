pub type Approval{
  Yes
  No
  Maybe
}
pub type Cuisine {
  Korean
  Turkish
}

pub type Genre {
  Crime
  Horror
  Romance
  Thriller
}

pub type Activity {
  BoardGame
  Chill
  Movie(Genre)
  Restaurant(Cuisine)
  Walk(Int)
}

pub fn rate_activity(activity: Activity) -> Approval {
  case activity {
    BoardGame -> No
    Chill -> No
    Movie(g) -> handle_genre(g)
    Restaurant(c) -> handle_cuisine(c)
    Walk(d) -> handle_walk(d)
  }
}

fn handle_genre(g: Genre) -> Approval {
  case g {
    Romance -> Yes
    _ -> No
  }
}

fn handle_cuisine(c: Cuisine) -> Approval {
  case c {
    Korean -> Yes
    Turkish -> Maybe
  }
}

fn handle_walk(d: Int) -> Approval {
  case d {
    d if d > 11 -> Yes
    d if d > 6 -> Maybe
    _ -> No
  }
}
