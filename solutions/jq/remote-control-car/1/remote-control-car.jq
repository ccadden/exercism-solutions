def new_remote_control_car:
  {
     "battery_percentage": 100,
     "distance_driven_in_meters": 0,
     "nickname": null
   }
;

def new_remote_control_car($nickname):
  {
     "battery_percentage": 100,
     "distance_driven_in_meters": 0,
     "nickname": $nickname
   }
;

def display_distance:
  "\(.distance_driven_in_meters // 0) meters"
;

def display_battery:
  # (if .battery_percentage // 0 == 0 then "empty" else .battery_percentage)
  "Battery \(if .battery_percentage == 0 then "empty" else "at \(.battery_percentage)%" end)"
;

def drive:
  if .battery_percentage > 0 then
    (.battery_percentage = .battery_percentage - 1 |
    .distance_driven_in_meters = .distance_driven_in_meters + 20)
  else . end
;
