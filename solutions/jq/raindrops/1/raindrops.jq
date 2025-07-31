.number as $input |
{"Pling": 3, "Plang": 5, "Plong": 7} |
to_entries |
map(select($input % .value == 0) | .key) |
add // $input
